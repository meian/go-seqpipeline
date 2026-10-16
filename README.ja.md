# go-seqpipeline

[English](README.md) | 日本語  

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
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Take(1).Operator)

	// パイプラインの結果をスライスとして受け取る
	fmt.Println(seq.ToSlice(pipeline)) // [333]

	// パイプラインを直接反復する
	for s := range pipeline {
		fmt.Println(s) // 333
	}
	// Seq2 ではキーと値を加工する
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe(seq.Map2(func(k int, v string) (string, int) {
			return strconv.Itoa(k), len(v)
		})).
		Pipe(seq.Filter2(func(k string, v int) bool { return v > 1 })).
		Pipe(seq.Take2(1).Operator)
	for k, v := range pairs {
		fmt.Println(k, v) // 1 3
	}
}
```

## Seq / Seq2

このライブラリの `Seq[T]` / `Seq2[K, V]` は、Go 標準の [`iter`](https://pkg.go.dev/iter) が提供する `iter.Seq[T]` / `iter.Seq2[K, V]` と互換性があり、`Pipe` を通じて要素を加工する機能を持つ関数型です。  
`Seq` / `Seq2` を `Pipe` を通じて加工したパイプラインは、`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
`Seq2[K, V]` の `K` と `V` に型の制限はありませんが、map との相互変換を伴う `FromMap` / `ToMap` では、`K` が `comparable` に制限されます。  
`Seq2.Keys()` / `Seq2.Values()` は、入力の反復順序と重複を保ち、キーまたは値を `Seq` として取り出します。  
戻り値の `Seq` には、`Pipe` で `Map`、`Filter`、`Take` などを接続できます。  
入力は遅延評価し、下流の途中終了を上流へ伝播します。  
再評価できるかどうかは元のイテレーターに従います。  

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
`Seq` 用の関数型 `Operator[I, O]` は、`func(Seq[I]) Seq[O]` と定義されています。  
`Seq2` では、`Seq2` を別の `Seq2` に変換します。  
`Seq2` 用の関数型 `Operator2[K, V, K2, V2]` は、`func(Seq2[K, V]) Seq2[K2, V2]` と定義されています。  
`Seq` / `Seq2` は `Pipe` でオペレーターを受け取り、その処理を適用した新しい `Seq` / `Seq2` を返します。  

### ライブラリが提供するオペレーター

要素の変換や選別、取得件数の制限などの機能を提供しています。  
`Seq` 用の `Filter` と `Seq2` 用の `Filter2` など同じ機能のオペレーターを用意し、`Seq2` 用のオペレーター名には `2` を付けて提供しています。  

### 型推論

生成関数の引数から型を推論できる場合は、戻り値をそのまま `Pipe` に渡すことができます。  
コールバックなどの引数が入出力の型情報を持つため、型引数を指定する必要はありません。  

生成関数の引数から型を推論できない場合は、戻り値の `.Operator` をメソッド値として `Pipe` に渡す必要があります。  
要素型を含まない設定値だけを受け取るオペレーターは、この方法で入力の要素型、またはキーと値の型を推論します。  

### 接続

`Pipe` は、次の型の変換関数を直接受け取ります。  

- `Seq`: `seq.Operator[I, O]`（`func(seq.Seq[I]) seq.Seq[O]`）
- `Seq2`: `seq.Operator2[K, V, K2, V2]`（`func(seq.Seq2[K, V]) seq.Seq2[K2, V2]`）

`Pipe` は渡された変換関数をその場で呼び出します。  
`Operator` / `Operator2` は関数型なので、`op(src)` の形でも直接呼び出せます。  
独自の変換関数では、この呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
関数内で `src.Pipe(...)` を使い、既存のオペレーターを合成することもできます。  

独自の変換関数も、そのまま `Pipe` に渡せます。  

### 独自オペレーターの実装

`Seq` / `Seq2` を加工する独自のオペレーターを作成できます。  
変換関数を実装し、`Pipe` に直接渡します。  

入出力にはこのライブラリの `Seq` / `Seq2` を使います。  
入力型が未確定の設定値には、`Take` と同様に型パラメーターを持つ `Operator` メソッドを用意し、`.Operator` を `Pipe` に渡します。  

### 実行可能な利用例

各オペレーターの使い方は、期待出力を持つ [Go Example](seq/example_test.go) にまとめています。  
通常テストでは、境界値、nil、遅延評価、途中終了、再評価などを検証します。  

| 利用例 | Seq | Seq2 |
| --- | --- | --- |
| 要素・キーと値の変換 | [ExampleMap](seq/example_test.go#L11) | [ExampleMap2](seq/example_test.go#L17) |
| 条件による選別 | [ExampleFilter](seq/example_test.go#L27) | [ExampleFilter2](seq/example_test.go#L33) |
| 取得件数の制限と型推論（.Operator） | [ExampleTake](seq/example_test.go#L41) | [ExampleTake2](seq/example_test.go#L47) |
| 型を変換する処理の接続 | [ExampleSeq_Pipe](seq/example_test.go#L57) | [ExampleMap2_directCall](seq/example_test.go#L66) |
| 関数オペレーターの接続と直接呼び出し | [ExampleOperator](seq/example_test.go#L81) | [ExampleOperator2](seq/example_test.go#L128) |
| 独自オペレーターによる拡張 | [ExampleOperator_custom](seq/example_test.go#L111) | [ExampleOperator2](seq/example_test.go#L128) |
| キーまたは値を Seq として取り出す | — | [ExampleSeq2_Keys](seq/example_test.go#L144), [ExampleSeq2_Values](seq/example_test.go#L152) |

Example のみを実行するには、次のコマンドを使います。  

```sh
go test ./... -run '^Example' -count=1
```

## 評価と利用上の仕様

### 遅延評価と途中終了

組み込みオペレーターによるパイプライン構築と `ToSeq` / `ToSeq2` では、要素を評価しません。  
`ToSlice` / `ToMap` の呼び出し、またはイテレーターの実行時に、入力の反復順序で処理します。  

下流の `yield` が `false` を返すか、`range` を `break` すると、上流の反復も停止します。  
例えば `Take` / `Take2` は、`n <= 0` なら上流を実行しません。  
上限に達すると `return` で `range` を抜けて上流へ停止を伝えるため、それ以降の上流の処理も停止します。  

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
組み込みの `Map` / `Filter` / `Take` とその `Seq2` 用 API、および `Keys` / `Values` は、nil の入力イテレーターを空として扱います。  
独自オペレーターの nil の扱いは、渡した変換関数の実装に従います。  

`ToSeq` / `ToSeq2` は nil をそのまま返します。  
nil を直接呼び出す、`range` で反復する、`ToSlice` / `ToMap` に渡すと panic します。  
空の非 nil イテレーターを収集すると、`ToSlice` は nil、`ToMap` は空の非 nil の map を返します。  

## ライセンス

[MIT License](LICENSE)  
