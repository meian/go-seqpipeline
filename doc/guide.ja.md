# 利用ガイド

[English](guide.md) | 日本語  

[README](../README.ja.md) の基本例から、自分のデータでパイプラインを組む手順を説明します。  
個別 API の契約は [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) の GoDoc を参照してください。  

## seq・op と Seq・Seq2 の役割

`seq` は入力・結果の変換と、処理を接続する `seq.Seq` / `seq.Seq2` を提供します。  
`op` は要素の変換や選別、取得件数の制限などのオペレーターを提供します。  
オペレーターは `iter.Seq` / `iter.Seq2` を変換して新しいシーケンスを生成する処理で、`seq` の `Pipe` / `Pipe2` に接続して使います。  

`seq.Seq[T]` / `seq.Seq2[K, V]` は、[`iter`](https://pkg.go.dev/iter) の関数型に、加工処理を接続する `Pipe` / `Pipe2` メソッドを追加した型です。  
結果は標準イテレーターと同様に、`for v := range pipeline` / `for k, v := range pipeline` で直接反復できます。  
`seq.Seq2[K, V]` の `K` と `V` に型の制限はありませんが、map との相互変換を伴う `FromMap` / `ToMap` では、`K` が `comparable` に制限されます。  
`seq.Seq2` から単一値の `seq.Seq` へ変換し、そのままパイプラインを続けることもできます。  

単一値とキー・値の組に別々の関数で同じ操作を提供する場合は、`Filter` と `Filter2` のように、キー・値の組を扱う名前に `2` を付けています。  

## 入力と結果の取り出し

`From` で始まる関数は、コレクション、標準イテレーター、channel をパイプラインの入力へ変換します。  
`Range(start, stop, step)` は `stop` にちょうど達した場合にそれを含む整数列を入力として生成します。  
`To` で始まる関数は、結果をコレクションへ収集するか、標準イテレーターや channel へ変換します。  

| 操作 | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| コレクションから入力 | `FromSlice` | `FromMap` |
| 標準イテレーターから入力 | `FromSeq` | `FromSeq2` |
| コレクションへ収集 | `ToSlice` | `ToMap`（キーは `comparable`） |
| 標準イテレーターへ変換 | `ToSeq` | `ToSeq2` |
| channel から入力 / channel へ出力 | `FromChannel` / `ToChannel` | — |

`iter.Seq` / `iter.Seq2` を受け取る関数と組み合わせる場合は、`ToSeq` / `ToSeq2` で変換して渡します。  

## Pipe / Pipe2 による処理の接続

入力が `seq.Seq` / `seq.Seq2` のどちらでも、単一値を出力する処理には `Pipe`、キーと値の組を出力する処理には `Pipe2` を使います。  
要素の型を変える場合も、直前の出力型に合う処理を続けて接続できます。  

```go
pipeline := seq.FromSlice([]int{1, 22, 333}).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [1 22]
```

コールバックなどが型情報を持つ `Map` / `Filter` は、生成関数の戻り値をそのまま渡します。  
`Take` / `Skip` のように生成時の引数だけでは型が決まらない処理は、単一値には `.Operator`、キーと値の組には `.Operator2` を渡し、接続時に型を推論します。  

`Map` / `Map2` はシーケンスの形を変えずに値を変換します。  
`MapToSeq2` は各要素からキーと値の組を1つ、`MapToSeq` は各組から単一値を1つ生成します。  
どちらもコールバックから型を推論し、`Compose` / `Compose2` の合成を含め、`Pipe2` / `Pipe` へ接続できます。  
キーと値に `comparable` 制約はなく、順序と重複を保持します。  
利用例は [ExampleMapToSeq2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-MapToSeq2) / [ExampleMapToSeq](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-MapToSeq) を参照してください。  

次の例では `Pipe2` で単一値をキーと値の組にし、`Pipe` で値だけのシーケンスへ戻します。  
`Zip` の追加入力には `iter.Seq` を渡すため、ライブラリ型から渡す場合は `ToSeq` で変換してください。  
設定メソッドの名前は各 API に従い、`Zip` では組を出力する場合も `.Operator` を `Pipe2` に渡します。  

```go
pairs := seq.FromSlice([]int{1, 2, 3}).
    Pipe2(op.Zip(slices.Values([]string{"a", "b", "c"})).Operator).
    Pipe2(op.Take(2).Operator2)
values := pairs.Pipe(op.Values)
fmt.Println(seq.ToSlice(values)) // [a b]
```

型を変える接続は [ExampleSeq_Pipe](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq.Pipe) / [ExampleSeq2_Pipe2](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq2.Pipe2)、単一値と組の間の接続は [ExampleSeq_Pipe2](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq.Pipe2) / [ExampleSeq2_Pipe](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq2.Pipe) を参照してください。  
変換・選別・件数制限・射影・Zip の利用例は [op の Example 索引](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples) から参照できます。  

### 接続の関数型

4種類の関数型で入出力の組み合わせを表します。  

| 入力 | メソッド | 変換関数 | 出力 |
| --- | --- | --- | --- |
| `seq.Seq[I]` | `Pipe` | `op.Operator[I, O]` | `seq.Seq[O]` |
| `seq.Seq[I]` | `Pipe2` | `op.ToSeq2Operator[I, K, V]` | `seq.Seq2[K, V]` |
| `seq.Seq2[K, V]` | `Pipe` | `op.ToSeqOperator[K, V, O]` | `seq.Seq[O]` |
| `seq.Seq2[K, V]` | `Pipe2` | `op.Operator2[K, V, K2, V2]` | `seq.Seq2[K2, V2]` |

`Pipe` / `Pipe2` は入力を標準型へ変換して変換関数をその場で呼び出し、戻り値をライブラリ型へ戻します。  
この型変換では要素を評価・コピーせず、nil を保持します。  
入力の評価と停止伝播は、変換関数の実装に従います。  

## 評価と停止の基本

キャッシュ以外の組み込みの入力変換とオペレーターは遅延評価し、収集関数の呼び出しや反復で処理を開始します。  
`Cache` / `Cache2` は、`Pipe` / `Pipe2` での接続時を含め、呼び出し時に入力を全件評価します。  
channel への出力は、呼び出し時に goroutine で評価を開始します。  
`seq.Seq` / `seq.Seq2` の反復中に下流の `yield` が `false` を返すか、`range` を `break` すると、上流へ停止を伝えます。  
入力やコールバック内部の待機は強制中断できないため、呼び出し側で停止手段を用意してください。  

## Cache による評価結果の再利用

`Pipe(op.Cache)` / `Pipe2(op.Cache2)` で途中までの結果を保存し、同じキャッシュから複数の後続処理を組み立てられます。  
キャッシュは接続時に上流を一度だけ評価し、順序と重複を保持し、nil・空入力にも非 nil のシーケンスを返します。  
以降の反復では途中終了後も保存済みの先頭から読み出し、キャッシュより後の処理は反復ごとに実行します。  

利用例は [ExampleCache](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache) / [ExampleCache2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache2) を参照してください。  

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

## channel との相互変換

channel の通信待ちを含めて途中の中断をサポートするため、channel と相互変換する関数にはコンテキストを渡します。  
途中で処理をやめる場合は、変換時に渡したコンテキストをキャンセルしてください。  

入力の利用例は [ExampleFromChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-FromChannel)、キャンセルを含む出力の利用例は [ExampleToChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-ToChannel) を参照してください。  

キャンセルは channel の通信待ちを解除しますが、入力やコールバック内部の待機は強制中断できません。  
ただし、送信とキャンセルが競合した場合は送信が選ばれることがあり、タイミングによっては次の値を受信します。  

`FromChannel(ctx, ch)` は反復時に受信し、入力を close しません。  
`ToChannel(ctx, src)` はバッファなしの channel を返し、評価終了時に close します。  
上流の通信待ちも中断する場合は、上流にも同じコンテキストを渡してください。  

## Compose によるオペレーターの合成

入力に処理をつなぐ場合は、`seq.Seq` / `seq.Seq2` のメソッドを使います。  
既存オペレーターをまとめたプリセットを入力へ適用する前に作るには、単一値入力なら `op.Compose()`、キーと値の組の入力なら `op.Compose2()` を使います。  
最初の `Pipe` / `Pipe2` に渡す処理から、入出力の型を推論します。  
空の開始値はまだ入力型を持たず、最初の接続で専用の合成値を返します。  

```go
transform := op.Compose().
    Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)

first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform.Operator)
second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform.Operator)
fmt.Println(seq.ToSlice(first))  // [2 4]
fmt.Println(seq.ToSlice(second)) // [8 10]
```

合成値の `Pipe` / `Pipe2` で、単一値を出力する処理は `Pipe`、キーと値の組を出力する処理は `Pipe2` で追加します。  
各段階で合成全体の入力型を保ち、途中の出力型に次の処理を接続できるため、単一値と組を行き来する合成もできます。  
合成値の `.Operator` メソッド値を、最終出力の型に応じて `seq.Seq` / `seq.Seq2` のメソッドに渡します。  
このメソッド値は対応する `op.Operator`、`op.ToSeq2Operator`、`op.ToSeqOperator`、`op.Operator2` の関数型にも合うため、別の合成の変換処理としても使えます。  

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

## 独自オペレーターの実装

既存オペレーターの合成では必要な処理を表せない場合は、独自オペレーターを実装します。  

入出力に `iter.Seq` / `iter.Seq2` を使う変換関数を実装し、出力型に応じて `Pipe` / `Pipe2` に渡します。  
変換関数は、呼び出し時に入力を評価せず、実行時に処理するイテレーターを返してください。  
生成時の引数に型情報を持たない処理には、設定値と型パラメーターを持つ変換メソッドを用意し、そのメソッド値を出力型に応じて `Pipe` / `Pipe2` に渡します。  
独自実装の内部でオペレーターを組み合わせる場合も、入出力を標準イテレーターのまま保ち、`op` 内の合成を使えます。  

利用例は [ExampleOperator_custom](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Operator-Custom) / [ExampleOperator2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Operator2)、単一値と組をまたぐ関数型は [ExampleToSeq2Operator](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-ToSeq2Operator) / [ExampleToSeqOperator](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-ToSeqOperator) を参照してください。  

## 共通仕様の詳細

### 入力・順序・再評価

入出力の変換は入力の反復順序と重複を保ちますが、map の順序は不定で、`ToMap` は同じキーの最後の値を採用します。  
`FromSlice` / `FromMap` は入力をコピーせず、評価時点の要素を参照します。  
キャッシュ以外の組み込みオペレーターの内部状態は評価ごとに初期化し、再評価や並行評価の可否は入力とコールバックに従います。  
channel との相互変換以外の組み込み処理は、goroutine や channel を使いません。  

### nil と空の入力

`seq.Seq` / `seq.Seq2` のゼロ値は nil で、標準イテレーターとの相互変換でも nil を保持します。  
組み込みオペレーターは nil を空として扱いますが、直接の呼び出し・反復や `ToSlice` / `ToMap` への入力は panic します。  
空の非 nil イテレーターでは、`ToSlice` は nil、`ToMap` は空の非 nil の map を返します。  
独自オペレーターの評価・停止・nil の扱いは、その実装に従います。  
各 API の境界条件は、GoDoc と [seq](../seq) / [op](../op) のテストを参照してください。  

### Example 索引

- [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples)：入出力と `Pipe` / `Pipe2`（[ソース](../seq/example_test.go)）
- [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples)：組み込み・合成・独自オペレーター（[ソース](../op/example_test.go)）
