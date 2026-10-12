# go-seqpipeline

`iter.Seq` / `iter.Seq2` を使い、型安全なパイプラインをメソッドチェーンで構築するライブラリです。  

## 環境要件とインストール

Go 1.27.1 以降が必要です。  

```sh
go get github.com/meian/go-seqpipeline
```

## 使い方

`From` で始まる関数で入力を取り込み、`Pipe` でオペレーターを接続します。  
接続する処理の入出力から型が推論され、型が一致しない処理は指定できません。  
結果は `for range` で直接反復するか、`ToSlice` / `ToMap` でコレクションへ収集します。  

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

	// パイプラインの結果をスライスとして受け取る
	fmt.Println(seq.ToSlice(pipeline)) // [333]

	// パイプラインを直接反復する
	for s := range pipeline {
		fmt.Println(s) // 333
	}
}
```

## Seq / Seq2

Go 標準の [`iter`](https://pkg.go.dev/iter) は、要素を扱う `iter.Seq[T]` と、キーと値などの組を扱う `iter.Seq2[K, V]` を提供します。  
このライブラリの `Seq[T]` / `Seq2[K, V]` は、`iter.Seq[T]` / `iter.Seq2[K, V]` と互換性があり、`Pipe` を通じて要素を加工する機能を持つ関数型です。  
`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
`Seq2[K, V]` の `K` と `V` に型の制限はありません。  
ただし、map との相互変換を伴う `FromMap` / `ToMap` では、`K` が `comparable` に制限されます。  

このライブラリでは、`Seq2` 用の処理は名前に `2` を付けて提供します。  
`Seq` と `Seq2` をまたぐ変換は提供していません。  

## 入力と出力

`From` で始まる関数は、コレクションや標準イテレーターをパイプラインの入力へ変換します。  
`To` で始まる関数は、結果をコレクションへ収集するか、標準イテレーターへ変換します。  

| 操作 | Seq | Seq2 |
| --- | --- | --- |
| コレクションから入力 | `FromSlice` | `FromMap` |
| 標準イテレーターから入力 | `FromSeq` | `FromSeq2` |
| コレクションへ収集 | `ToSlice` | `ToMap`（キーは `comparable`） |
| 標準イテレーターへ変換 | `ToSeq` | `ToSeq2` |

標準の `iter.Seq` / `iter.Seq2` を受け取る関数と組み合わせる場合は、`ToSeq` / `ToSeq2` で変換して渡します。  

## オペレーター

このライブラリでは、`Seq` を別の `Seq` に変換する処理をオペレーターと呼びます。  
`Seq2` では、`Seq2` を別の `Seq2` に変換します。  
`Seq` / `Seq2` は `Pipe` でオペレーターを受け取り、その処理を適用した新しい `Seq` / `Seq2` を返します。  

### 型推論

生成関数の引数から型を推論できる場合は、戻り値をそのまま `Pipe` に渡します。  
コールバックなどの引数が入出力の型情報を持つため、型引数を指定する必要はありません。  

生成関数の引数から型を推論できない場合は、戻り値の `.Type` をメソッド値として `Pipe` に渡します。  
要素型を含まない設定値だけを受け取るオペレーターは、この方法で入力の要素型、またはキーと値の型を推論します。  

### 接続

`Pipe` は、次の変換関数を直接受け取ります。  

- `Seq`: `func(seq.Seq[I]) seq.Seq[O]`
- `Seq2`: `func(seq.Seq2[K, V]) seq.Seq2[K2, V2]`

`Pipe` は渡された変換関数をその場で呼び出します。  
独自の変換関数では、この呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
関数内で `src.Pipe(...)` を使い、既存のオペレーターを合成することもできます。  

`OperatorFunc` / `Operator2Func` は、変換関数を `Apply` メソッドを持つ専用の関数型へ包みます。  
戻り値は `Pipe` に直接渡せるほか、`Operator` / `Operator2` として扱えます。  

```go
text := seq.OperatorFunc(func(src seq.Seq[int]) seq.Seq[string] {
	return src.Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) }))
})
pipeline := seq.FromSlice([]int{1, 2}).Pipe(text)
fmt.Println(seq.ToSlice(pipeline)) // [1 2]
```

### 独自オペレーターの実装

`Seq` / `Seq2` を加工する独自のオペレーターを作成できます。  
`Operator` / `Operator2` の契約に沿って `Apply` メソッドを実装し、`Pipe(op.Apply)` に渡すことで独自の変換処理を提供します。  
変換関数型で実装すれば、オペレーターを `Pipe` に直接渡すこともできます。  

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
