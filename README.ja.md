# go-seqpipeline

[English](README.md) | 日本語  

`iter.Seq` / `iter.Seq2` を使い、型安全なパイプラインをメソッドチェーンで構築するライブラリです。  

## 環境要件とインストール

Go 1.27.1 以降が必要です。  

```sh
go get github.com/meian/go-seqpipeline
```

## 使い方

`From` で始まる関数で入力を取り込み、`Pipe` / `Pipe2` でオペレーターを接続します。  
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
		Pipe2(seq.Map2(func(k int, v string) (string, int) {
			return strconv.Itoa(k), len(v)
		})).
		Pipe2(seq.Filter2(func(k string, v int) bool { return v > 1 })).
		Pipe2(seq.Take(1).Operator2)
	for k, v := range pairs {
		fmt.Println(k, v) // 1 3
	}
}
```

## Seq / Seq2

`Seq[T]` / `Seq2[K, V]` は、Go 標準の [`iter`](https://pkg.go.dev/iter) の関数型に、加工処理を接続する `Pipe` / `Pipe2` メソッドを追加した型です。  
結果は標準イテレーターと同様に、`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
`Seq2[K, V]` の `K` と `V` に型の制限はありませんが、map との相互変換を伴う `FromMap` / `ToMap` では、`K` が `comparable` に制限されます。  
`Seq2` から単一値の `Seq` へ変換し、そのままパイプラインを続けることもできます。  

## 評価と利用上の仕様

### 評価と停止

キャッシュ以外の組み込みの入力変換とオペレーターは遅延評価し、収集関数の呼び出しや反復で処理を開始します。  
`Cache` / `Cache2` は、`Pipe` / `Pipe2` での接続時を含め、呼び出し時に入力を全件評価します。  
channel への出力は、呼び出し時に goroutine で評価を開始します。  
`Seq` / `Seq2` の反復中に下流の `yield` が `false` を返すか、`range` を `break` すると、上流へ停止を伝えます。  
入力やコールバック内部の待機は強制中断できないため、呼び出し側で停止手段を用意してください。  

### 入力・順序・再評価

入出力の変換は入力の反復順序と重複を保ちますが、map の順序は不定で、`ToMap` は同じキーの最後の値を採用します。  
`FromSlice` / `FromMap` は入力をコピーせず、評価時点の要素を参照します。  
キャッシュ以外の組み込みオペレーターの内部状態は評価ごとに初期化し、再評価や並行評価の可否は入力とコールバックに従います。  
channel との相互変換以外の組み込み処理は、goroutine や channel を使いません。  

### 評価結果の再利用

`Pipe(seq.Cache)` / `Pipe2(seq.Cache2)` で途中までの結果を保存し、同じキャッシュから複数の後続処理を組み立てられます。  
キャッシュは接続時に上流を一度だけ評価し、順序と重複を保持し、nil・空入力にも非 nil のシーケンスを返します。  
以降の反復では途中終了後も保存済みの先頭から読み出し、キャッシュより後の処理は反復ごとに実行します。  

```go
cached := seq.FromSlice([]string{"a", "bbb", "cc"}).
    Pipe(seq.Map(func(v string) int { return len(v) })).
    Pipe(seq.Cache)
first := cached.Pipe(seq.Take(1).Operator)
longer := cached.Pipe(seq.Filter(func(v int) bool { return v > 1 }))
fmt.Println(seq.ToSlice(first))  // [1]
fmt.Println(seq.ToSlice(longer)) // [3 2]
```

入力は有限である必要があり、保存する要素数・組数に比例するメモリを使います。  
無限入力はキャッシュより前で件数を制限してください。  
後続処理の件数制限では初回の全件評価を止められません。  

たとえば、次のように無限に値を生成する入力を考えます。  

```go
naturals := seq.Seq[int](func(yield func(int) bool) {
    for n := 1; ; n++ {
        if !yield(n) {
            return
        }
    }
})
```

次の接続は、`Cache` が無限入力を全件評価しようとするため戻りません。  
後ろの `Take(3)` には到達しません。  

```go
naturals.Pipe(seq.Cache).Pipe(seq.Take(3).Operator)
```

`Take(3)` を前に置くと、上流を3件で停止してキャッシュを作成でき、繰り返し利用できます。  
`Seq2` でも `Take(3).Operator2` を `Cache2` より前に接続します。  

```go
cached := naturals.Pipe(seq.Take(3).Operator).Pipe(seq.Cache)
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
```

要素は値として保存し、深いコピーは行わないため、元の要素の置き換えは反映せず、ポインター・スライス・map などで共有する参照先の変更は反映します。  
入力の反復はキャッシュ呼び出しが戻る前に終了し、キャッシュが使われなくなると保存領域は回収可能になります。  
キャッシュは交互・並行に反復できますが、共有する参照先や下流のコールバックは必要に応じて呼び出し側で同期してください。  

### nil と空の入力

`Seq` / `Seq2` のゼロ値は nil で、標準イテレーターとの相互変換でも nil を保持します。  
組み込みオペレーターは nil を空として扱いますが、直接の呼び出し・反復や `ToSlice` / `ToMap` への入力は panic します。  
空の非 nil イテレーターでは、`ToSlice` は nil、`ToMap` は空の非 nil の map を返します。  
独自オペレーターの評価・停止・nil の扱いは、その実装に従います。  
各 API の境界条件は [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) と [テスト](seq) を参照してください。  

## 入力と出力

`From` で始まる関数は、コレクション、標準イテレーター、channel をパイプラインの入力へ変換します。  
`To` で始まる関数は、結果をコレクションへ収集するか、標準イテレーターや channel へ変換します。  

| 操作 | Seq | Seq2 |
| --- | --- | --- |
| コレクションから入力 | `FromSlice` | `FromMap` |
| 標準イテレーターから入力 | `FromSeq` | `FromSeq2` |
| コレクションへ収集 | `ToSlice` | `ToMap`（キーは `comparable`） |
| 標準イテレーターへ変換 | `ToSeq` | `ToSeq2` |
| channel から入力 / channel へ出力 | `FromChannel` / `ToChannel` | — |

標準の `iter.Seq` / `iter.Seq2` を受け取る関数と組み合わせる場合は、`ToSeq` / `ToSeq2` で変換して渡します。  

### channel との相互変換

channel の通信待ちを含めて途中の中断をサポートするため、channel と相互変換する関数にはコンテキストを渡します。  
途中で処理をやめる場合は、変換時に渡したコンテキストをキャンセルしてください。  

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

src := seq.FromSlice([]int{1, 2, 3})
for v := range seq.ToChannel(ctx, src) {
    fmt.Println(v)
    cancel()
}
```

`cancel()` を呼ぶと残りの送信を中断し、channel を閉じてループを終了します。  
ただし、送信とキャンセルが競合した場合は送信が選ばれることがあり、タイミングによっては次の値を受信します。  

`FromChannel(ctx, ch)` は反復時に受信し、入力を close しません。  
`ToChannel(ctx, src)` はバッファなしの channel を返し、評価終了時に close します。  
上流の通信待ちも中断する場合は、上流にも同じコンテキストを渡してください。  

## オペレーター

go-seqpipeline では、`Seq` / `Seq2` を変換して新しいシーケンスを生成する処理をオペレーターと定義しています。  

### ライブラリが提供するオペレーター

要素の変換や選別、取得件数の制限などの機能を提供しています。  
単一値とキー・値の組に別々の関数で同じ操作を提供する場合は、`Filter` と `Filter2` のように、キー・値の組を扱う名前に `2` を付けて提供しています。  

### 型推論

生成関数の引数から型を推論できる場合は、戻り値の出力型に応じて、そのまま `Pipe` / `Pipe2` に渡すことができます。  
コールバックなどの引数が入出力の型情報を持つため、型引数を指定する必要はありません。  

生成時の引数に型情報を持たない処理は、設定値を返し、そのメソッドで入力から型を推論します。  
単一値を扱う場合は `.Operator` を `Pipe` に、キーと値の組を扱う場合は `.Operator2` を `Pipe2` にメソッド値として渡します。  
接続時に入力の型が決まるため、型引数を明示する必要はありません。  

### 接続

`Seq` と `Seq2` には `Pipe` / `Pipe2` が定義されており、それぞれ `Seq` / `Seq2` に変換するオペレーター関数を受け取ります。  
関数型 `Operator`、`ToSeq2Operator`、`ToSeqOperator`、`Operator2` で、入出力の型の 4 通りの組み合わせを表します。  

| 入力 | メソッド | 変換関数 | 出力 |
| --- | --- | --- | --- |
| `Seq[I]` | `Pipe` | `Operator[I, O]`（`func(Seq[I]) Seq[O]`） | `Seq[O]` |
| `Seq[I]` | `Pipe2` | `ToSeq2Operator[I, K, V]`（`func(Seq[I]) Seq2[K, V]`） | `Seq2[K, V]` |
| `Seq2[K, V]` | `Pipe` | `ToSeqOperator[K, V, O]`（`func(Seq2[K, V]) Seq[O]`） | `Seq[O]` |
| `Seq2[K, V]` | `Pipe2` | `Operator2[K, V, K2, V2]`（`func(Seq2[K, V]) Seq2[K2, V2]`） | `Seq2[K2, V2]` |

`Pipe` / `Pipe2` は変換関数をその場で呼び出し、nil を含めて戻り値をそのまま返します。  
入力の評価と停止伝播は、変換関数の実装に従います。  
4 種類のオペレーターは関数型なので、`op(src)` の形でも直接呼び出せます。  

### 独自オペレーターの実装

入出力にこのライブラリの `Seq` / `Seq2` を使う変換関数を実装し、出力型に応じて `Pipe` / `Pipe2` に直接渡します。  
変換関数は、呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
関数内で `src.Pipe(...)` / `src.Pipe2(...)` を使い、既存のオペレーターを合成することもできます。  
生成時の引数に型情報を持たない処理には、設定値と型パラメーターを持つ変換メソッドを用意し、そのメソッド値を出力型に応じて `Pipe` / `Pipe2` に渡します。  

### 実行可能な利用例

各オペレーターの使い方は、期待出力を持つ [Go Example](seq/example_test.go) にまとめています。  
通常テストでは、境界値、nil、遅延評価、途中終了、再評価などを検証します。  

| 利用例 | Seq | Seq2 |
| --- | --- | --- |
| 要素・キーと値の変換 | [ExampleMap](seq/example_test.go#L11) | [ExampleMap2](seq/example_test.go#L17) |
| 条件による選別 | [ExampleFilter](seq/example_test.go#L27) | [ExampleFilter2](seq/example_test.go#L33) |
| 取得件数の制限と型推論（.Operator / .Operator2） | [ExampleTake](seq/example_test.go#L41) | [ExampleTake_pairs](seq/example_test.go#L47) |
| 先頭の要素・組の読み飛ばし | [ExampleSkip](seq/example_test.go#L255) | [ExampleSkip](seq/example_test.go#L255) |
| 型を変換する処理の接続 | [ExampleSeq_Pipe](seq/example_test.go#L57) | [ExampleSeq2_Pipe2](seq/example_test.go#L199) |
| Seq と Seq2 をまたぐ接続 | [ExampleSeq_Pipe2](seq/example_test.go#L179) | [ExampleSeq2_Pipe](seq/example_test.go#L191) |
| Seq2 からキー・値を取り出す | — | [ExampleKeys](seq/example_test.go#L144)、[ExampleValues](seq/example_test.go#L153) |
| Seq と Seq2 をまたぐオペレーターの直接呼び出し | [ExampleToSeq2Operator](seq/example_test.go#L211) | [ExampleToSeqOperator](seq/example_test.go#L222) |
| 関数オペレーターの接続と直接呼び出し | [ExampleOperator](seq/example_test.go#L81) | [ExampleOperator2](seq/example_test.go#L128) |
| 独自オペレーターによる拡張 | [ExampleOperator_custom](seq/example_test.go#L111) | [ExampleOperator2](seq/example_test.go#L128) |

Example のみを実行するには、次のコマンドを使います。  

```sh
go test ./... -run '^Example' -count=1
```

## ベンチマーク

[ベンチマークレポート](benchmarks/REPORT.md) に、パイプライン・ループ・要素ごとの関数呼び出し・channel・キャッシュの再現可能な比較をまとめています。  
測定対象のコミット、測定条件、生データ、集計スクリプトを含みます。  
ベンチマークとレポートは定期更新用の専用 Issue で更新し、通常のコード変更では再測定を求めません。  

## ライセンス

[MIT License](LICENSE)  
