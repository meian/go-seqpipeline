# go-seqpipeline

`iter.Seq` / `iter.Seq2` を使い、型安全なパイプラインをメソッドチェーンで構築するライブラリです。  

## 環境要件とインストール

Go 1.27.1 以降が必要です。  

```sh
go get github.com/meian/go-seqpipeline
```

## 基本的な使い方

`From` で始まる関数で入力を取り込み、`Pipe` で変換・選別・件数制限を接続します。  
入出力の型が一致しない接続はコンパイルエラーになります。  
`To` で始まる関数で結果を収集するか、標準のイテレーターへ変換します。  

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Take(1).Type)

	fmt.Println(seq.ToSlice(pipeline)) // [333]
}
```

## Seq / Seq2 と命名規則

`Seq[T]` は要素を、`Seq2[K, V]` はキーと値の組を扱う関数型です。  
`Seq2` のキー型は、`ToMap` などで map を扱う場合を除き、`comparable` に制限されません。  
対応する API の名前に `2` が付くものは `Seq2` 用です（`Map2`、`Filter2`、`Take2`、`Operator2` など）。  
スライスや map を扱う API は、対象を表す名前を使います。  

| 操作 | Seq | Seq2 |
| --- | --- | --- |
| コレクションから入力 | `FromSlice` | `FromMap` |
| 標準イテレーターから入力 | `FromSeq` | `FromSeq2` |
| コレクションへ収集 | `ToSlice` | `ToMap`（キーは `comparable`） |
| 標準イテレーターへ変換 | `ToSeq` | `ToSeq2` |

`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
標準の `iter.Seq` / `iter.Seq2` を受け取る関数には、`ToSeq` / `ToSeq2` で変換して渡します。  
`Seq` と `Seq2` をまたぐ変換は提供していません。  

## オペレーターの型推論

生成関数の引数から型を推論できる場合は、戻り値をそのまま `Pipe` に渡します。  
`Map` / `Filter` とその `Seq2` 用 API は、コールバックの引数・戻り値から型が決まります。  

生成関数の引数から型を推論できない場合は、戻り値の `.Type` をメソッド値として `Pipe` に渡します。  
`Take(n)` / `Take2(n)` は件数だけを受け取るため、`Pipe(seq.Take(n).Type)` / `Pipe(seq.Take2(n).Type)` として入力型から推論します。  

## 関数オペレーターの接続

`Pipe` は、次の変換関数を直接受け取ります。  

- `Seq`: `func(seq.Seq[I]) seq.Seq[O]`
- `Seq2`: `func(seq.Seq2[K, V]) seq.Seq2[K2, V2]`

`Pipe` は渡された変換関数をその場で呼び出します。  
独自の変換関数では、この呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
関数内で `src.Pipe(...)` を使い、既存のオペレーターを合成することもできます。  

`OperatorFunc` / `Operator2Func` は、変換関数を `Apply` メソッドを持つ専用の関数型へ包みます。  
戻り値は `Pipe` に直接渡せるほか、`Operator` / `Operator2` として扱えます。  
以下は `main` 関数内の抜粋で、`package` 宣言と `fmt`・`strconv`・`seq` の import を省略しています。  

```go
text := seq.OperatorFunc(func(src seq.Seq[int]) seq.Seq[string] {
	return src.Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) }))
})
pipeline := seq.FromSlice([]int{1, 2}).Pipe(text)
fmt.Println(seq.ToSlice(pipeline)) // [1 2]
```

## 独自オペレーターの拡張

`Seq` / `Seq2` を変更せずにオペレーターを追加できます。  
生成関数が変換関数型を返せば、その戻り値を `Pipe` に直接渡せます。  
関数型には設定メソッドや `Apply` メソッドを追加できます。  
型自体が非公開でも、公開メソッドは外部パッケージから呼び出せます。  

次の例では、既存の `Map` を使って乗算オペレーターを作ります。  
`Pipe(Multiply(10))` で接続できます。  
コードは宣言部分の抜粋で、`package` 宣言と `seq` の import を省略しています。  

```go
func Multiply(factor int) multiplyOperator {
	return func(src seq.Seq[int]) seq.Seq[int] {
		return src.Pipe(seq.Map(func(v int) int { return v * factor }))
	}
}

type multiplyOperator func(seq.Seq[int]) seq.Seq[int]

func (m multiplyOperator) Apply(src seq.Seq[int]) seq.Seq[int] {
	return m(src)
}

var _ seq.Operator[int, int] = multiplyOperator(nil)
```

`Operator` / `Operator2` インターフェイスは、`Apply` による変換の契約を表します。  
構造体などで実装したオペレーターは、`Pipe(op.Apply)` で接続します。  
入出力にはこのライブラリの `Seq` / `Seq2` を使います。  
入力型が未確定の設定値には、`Take` と同様に型パラメーターを持つ `Type` メソッドを用意し、`.Type` を `Pipe` に渡します。  

## 評価と利用上の仕様

### 遅延評価と途中終了

組み込みオペレーターによるパイプライン構築と `ToSeq` / `ToSeq2` では、要素を評価しません。  
`ToSlice` / `ToMap` の呼び出し、またはイテレーターの実行時に、入力の反復順序で処理します。  
下流の `yield` が `false` を返すか、`range` を `break` すると、上流の反復も停止します。  
`Take` / `Take2` は上限に達すると停止し、`n <= 0` なら上流を実行しません。  
独自オペレーターも、返すイテレーターの実行時に処理し、下流からの停止を上流へ伝える必要があります。  

### 再評価と入力の参照

再評価できるかどうかは元のイテレーターに従います。  
`Take` / `Take2` の残り件数は評価ごとに初期化します。  
`FromSlice` / `FromMap` は入力をコピーせず、評価時点の要素を参照します。  
スライスは元の順序で反復し、map の反復順序は保証しません。  
`ToMap` は同じキーでは反復順で後の値を採用します。  

ライブラリは goroutine や channel を使わず、非同期イベント処理を提供しません。  
状態を持つコールバックの再利用や並行アクセスは、呼び出し側で管理してください。  

### nil と空の入力

`Seq` / `Seq2` のゼロ値と、nil の標準イテレーターから変換した値は nil です。  
組み込みの `Map` / `Filter` / `Take` とその `Seq2` 用 API は、nil の入力イテレーターを空として扱います。  
`OperatorFunc` / `Operator2Func` や独自オペレーターの nil の扱いは、渡した変換関数の実装に従います。  

`ToSeq` / `ToSeq2` は nil をそのまま返します。  
nil を直接呼び出す、`range` で反復する、`ToSlice` / `ToMap` に渡すと panic します。  
空の非 nil イテレーターを収集すると、`ToSlice` は nil、`ToMap` は空の非 nil の map を返します。  

## ライセンス

[MIT License](LICENSE)  
