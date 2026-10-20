# 利用ガイド

[English](guide.md) | 日本語  

[README](../README.ja.md) の基本例から先へ進むための共通仕様と接続方法を説明します。  
型推論の基本は [README](../README.ja.md#型推論)、個別 API の契約は [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) の GoDoc を参照してください。  

## seq.Seq / seq.Seq2

`seq.Seq[T]` / `seq.Seq2[K, V]` は、[`iter`](https://pkg.go.dev/iter) の関数型に、加工処理を接続する `Pipe` / `Pipe2` メソッドを追加した型です。  
結果は標準イテレーターと同様に、`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
`seq.Seq2[K, V]` の `K` と `V` に型の制限はありませんが、map との相互変換を伴う `FromMap` / `ToMap` では、`K` が `comparable` に制限されます。  
`seq.Seq2` から単一値の `seq.Seq` へ変換し、そのままパイプラインを続けることもできます。  

## 評価と利用上の仕様

### 評価と停止

キャッシュ以外の組み込みの入力変換とオペレーターは遅延評価し、収集関数の呼び出しや反復で処理を開始します。  
`Cache` / `Cache2` は、`Pipe` / `Pipe2` での接続時を含め、呼び出し時に入力を全件評価します。  
channel への出力は、呼び出し時に goroutine で評価を開始します。  
`seq.Seq` / `seq.Seq2` の反復中に下流の `yield` が `false` を返すか、`range` を `break` すると、上流へ停止を伝えます。  
入力やコールバック内部の待機は強制中断できないため、呼び出し側で停止手段を用意してください。  

### 入力・順序・再評価

入出力の変換は入力の反復順序と重複を保ちますが、map の順序は不定で、`ToMap` は同じキーの最後の値を採用します。  
`FromSlice` / `FromMap` は入力をコピーせず、評価時点の要素を参照します。  
キャッシュ以外の組み込みオペレーターの内部状態は評価ごとに初期化し、再評価や並行評価の可否は入力とコールバックに従います。  
channel との相互変換以外の組み込み処理は、goroutine や channel を使いません。  

### 評価結果の再利用

`Pipe(op.Cache)` / `Pipe2(op.Cache2)` で途中までの結果を保存し、同じキャッシュから複数の後続処理を組み立てられます。  
キャッシュは接続時に上流を一度だけ評価し、順序と重複を保持し、nil・空入力にも非 nil のシーケンスを返します。  
以降の反復では途中終了後も保存済みの先頭から読み出し、キャッシュより後の処理は反復ごとに実行します。  

利用例は [ExampleCache](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache) を参照してください。  

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
naturals.Pipe(op.Cache).Pipe(op.Take(3).Operator)
```

`Take(3)` を前に置くと、上流を3件で停止してキャッシュを作成でき、繰り返し利用できます。  
`seq.Seq2` でも `Take(3).Operator2` を `Cache2` より前に接続します。  

```go
cached := naturals.Pipe(op.Take(3).Operator).Pipe(op.Cache)
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
```

要素は値として保存し、深いコピーは行わないため、元の要素の置き換えは反映せず、ポインター・スライス・map などで共有する参照先の変更は反映します。  
入力の反復はキャッシュ呼び出しが戻る前に終了し、キャッシュが使われなくなると保存領域は回収可能になります。  
キャッシュは交互・並行に反復できますが、共有する参照先や下流のコールバックは必要に応じて呼び出し側で同期してください。  

### nil と空の入力

`seq.Seq` / `seq.Seq2` のゼロ値は nil で、標準イテレーターとの相互変換でも nil を保持します。  
組み込みオペレーターは nil を空として扱いますが、直接の呼び出し・反復や `ToSlice` / `ToMap` への入力は panic します。  
空の非 nil イテレーターでは、`ToSlice` は nil、`ToMap` は空の非 nil の map を返します。  
独自オペレーターの評価・停止・nil の扱いは、その実装に従います。  
各 API の境界条件は、GoDoc と [seq](../seq) / [op](../op) のテストを参照してください。  

## 入力と出力

`From` で始まる関数は、コレクション、標準イテレーター、channel をパイプラインの入力へ変換します。  
`To` で始まる関数は、結果をコレクションへ収集するか、標準イテレーターや channel へ変換します。  

| 操作 | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| コレクションから入力 | `FromSlice` | `FromMap` |
| 標準イテレーターから入力 | `FromSeq` | `FromSeq2` |
| コレクションへ収集 | `ToSlice` | `ToMap`（キーは `comparable`） |
| 標準イテレーターへ変換 | `ToSeq` | `ToSeq2` |
| channel から入力 / channel へ出力 | `FromChannel` / `ToChannel` | — |

`iter.Seq` / `iter.Seq2` を受け取る関数と組み合わせる場合は、`ToSeq` / `ToSeq2` で変換して渡します。  

### channel との相互変換

channel の通信待ちを含めて途中の中断をサポートするため、channel と相互変換する関数にはコンテキストを渡します。  
途中で処理をやめる場合は、変換時に渡したコンテキストをキャンセルしてください。  

キャンセルを含む利用例は [ExampleToChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-ToChannel) を参照してください。  

キャンセルは channel の通信待ちを解除しますが、入力やコールバック内部の待機は強制中断できません。  
ただし、送信とキャンセルが競合した場合は送信が選ばれることがあり、タイミングによっては次の値を受信します。  

`FromChannel(ctx, ch)` は反復時に受信し、入力を close しません。  
`ToChannel(ctx, src)` はバッファなしの channel を返し、評価終了時に close します。  
上流の通信待ちも中断する場合は、上流にも同じコンテキストを渡してください。  

## オペレーター

go-seqpipeline では、`iter.Seq` / `iter.Seq2` を変換して新しいシーケンスを生成する処理をオペレーターと定義しています。  

### ライブラリが提供するオペレーター

要素の変換や選別、取得件数の制限などの機能を提供しています。  
単一値とキー・値の組に別々の関数で同じ操作を提供する場合は、`Filter` と `Filter2` のように、キー・値の組を扱う名前に `2` を付けて提供しています。  

### 接続

`seq.Seq` と `seq.Seq2` の `Pipe` / `Pipe2` は、標準イテレーターを扱うオペレーターを接続し、それぞれ `seq.Seq` / `seq.Seq2` を返します。  
関数型 `op.Operator`、`op.ToSeq2Operator`、`op.ToSeqOperator`、`op.Operator2` で、入出力の型の 4 通りの組み合わせを表します。  

| 入力 | メソッド | 変換関数 | 出力 |
| --- | --- | --- | --- |
| `seq.Seq[I]` | `Pipe` | `op.Operator[I, O]` | `seq.Seq[O]` |
| `seq.Seq[I]` | `Pipe2` | `op.ToSeq2Operator[I, K, V]` | `seq.Seq2[K, V]` |
| `seq.Seq2[K, V]` | `Pipe` | `op.ToSeqOperator[K, V, O]` | `seq.Seq[O]` |
| `seq.Seq2[K, V]` | `Pipe2` | `op.Operator2[K, V, K2, V2]` | `seq.Seq2[K2, V2]` |

`Pipe` / `Pipe2` は入力を標準型へ変換して変換関数をその場で呼び出し、戻り値をライブラリ型へ戻します。  
この型変換では要素を評価・コピーせず、nil を保持します。  
入力の評価と停止伝播は、変換関数の実装に従います。  
`Zip` などの追加入力にライブラリ型を渡す場合は、`ToSeq` / `ToSeq2` で変換します。  

### オペレーターの合成

入力を決める前に変換処理を組み立てるには、単一値入力なら `op.Compose()`、キーと値の組の入力なら `op.Compose2()` を使います。  
最初の `Pipe` / `Pipe2` に渡す処理から、入出力の型を推論します。  
空の開始値はまだ入力型を持たず、オペレーターでも恒等変換でもありません。  

```go
transform := op.Compose().
    Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)

first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform)
second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform)
fmt.Println(seq.ToSlice(first))  // [2 4]
fmt.Println(seq.ToSlice(second)) // [8 10]
```

4種類のOperator関数型はすべて `Pipe` / `Pipe2` メソッドを持ち、単一値を出力する処理は `Pipe`、キーと値の組を出力する処理は `Pipe2` で追加します。  
各段階で合成全体の入力型を保ち、途中の出力型に次の処理を接続できるため、単一値と組を行き来する合成もできます。  
結果はOperator関数型なので、そのまま `seq.Seq` / `seq.Seq2` のメソッドに渡したり、標準イテレーターへ直接適用したりできます。  
`op.Map(strconv.Itoa).Pipe(op.Take(2).Operator)` のように、名前付きのオペレーターから直接合成を始めることもできます。  

先頭の処理だけでは型情報が不足する場合、最初のメソッドで不足する型引数を指定します。  
2段目以降は途中の出力型から推論できます。  

```go
firstTwo := op.Compose().Pipe[int](op.Take(2).Operator)
firstTwoPairs := op.Compose2().Pipe2[int, string](op.Take(2).Operator2)
values := op.Compose2().Pipe[int, string](op.Values)
```

合成の組み立て時には変換関数を呼び出さず、入力も評価しません。  
合成結果を入力へ適用すると接続順に変換関数を呼び出し、各処理の評価・停止の仕様に従います。  
`Cache` / `Cache2` は適用時に呼び出しごとのキャッシュを作り、以降の反復ではそのキャッシュを再利用します。  
入力・コールバック・`Zip` の右側などの追加入力は、それぞれの再利用・状態共有の条件に従います。  

実行可能な利用例は [ExampleCompose](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose)、[ExampleCompose2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2)、[ExampleCompose2_values](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2-Values) を参照してください。  

### 独自オペレーターの実装

入出力に `iter.Seq` / `iter.Seq2` を使う変換関数を実装し、出力型に応じて `Pipe` / `Pipe2` に渡します。  
変換関数は、呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
生成時の引数に型情報を持たない処理には、設定値と型パラメーターを持つ変換メソッドを用意し、そのメソッド値を出力型に応じて `Pipe` / `Pipe2` に渡します。  
独自実装の内部でオペレーターを組み合わせる場合も、入出力を標準イテレーターのまま保ち、`op` 内の合成を使えます。  

### 実行可能な利用例

`Pipe` / `Pipe2` でオペレーターを接続する使い方は、[seq](../seq/example_test.go) / [op](../op/example_test.go) の Go Example にまとめています。  

| 利用例 | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| 要素・キーと値の変換 | [ExampleMap](../op/example_test.go#L13) | [ExampleMap2](../op/example_test.go#L20) |
| 条件による選別 | [ExampleFilter](../op/example_test.go#L31) | [ExampleFilter2](../op/example_test.go#L38) |
| 取得件数の制限と型推論（.Operator / .Operator2） | [ExampleTake](../op/example_test.go#L47) | [ExampleTake_pairs](../op/example_test.go#L53) |
| 先頭の要素・組の読み飛ばし | [ExampleSkip](../op/example_test.go#L208) | [ExampleSkip](../op/example_test.go#L208) |
| 型を変換する処理の接続 | [ExampleSeq_Pipe](../seq/example_test.go#L14) | [ExampleSeq2_Pipe2](../seq/example_test.go#L60) |
| `seq.Seq` と `seq.Seq2` をまたぐ接続 | [ExampleSeq_Pipe2](../seq/example_test.go#L40) | [ExampleSeq2_Pipe](../seq/example_test.go#L52) |
| `seq.Seq2` からキー・値を取り出す | — | [ExampleKeys](../op/example_test.go#L131)、[ExampleValues](../op/example_test.go#L140) |
| `seq.Seq` と `seq.Seq2` をまたぐ関数オペレーターの接続 | [ExampleToSeq2Operator](../op/example_test.go#L166) | [ExampleToSeqOperator](../op/example_test.go#L176) |
| 同じシーケンス型の間での関数オペレーターの接続 | [ExampleOperator](../op/example_test.go#L79) | [ExampleOperator2](../op/example_test.go#L121) |
| 独自オペレーターによる拡張 | [ExampleOperator_custom](../op/example_test.go#L105) | [ExampleOperator2](../op/example_test.go#L121) |

Example のみを実行するには、次のコマンドを使います。  

```sh
go test ./... -run '^Example' -count=1
```
