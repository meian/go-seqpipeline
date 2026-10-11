# go-seqpipeline

`iter.Seq` / `iter.Seq2` を使い、型安全なパイプラインをメソッドチェーンで構築するライブラリです。  

## 環境要件

- Go 1.27.1以降

## インストール

```sh
go get github.com/meian/go-seqpipeline
```

## Seqのパイプライン

入力は `From` で始まる関数、出力への変換は `To` で始まる関数で扱います。  

`Pipe` でオペレーターを接続し、型変換を含むパイプラインを構築します。  
型が一致しない接続はコンパイルエラーになります。  
`ToSlice` でスライスへ収集し、`ToSeq` で標準の `iter.Seq` を取得できます。  

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	values := []int{1, 22, 333, 4444}
	pipeline := seq.FromSlice(values).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Take(1).Type)

	fmt.Println(seq.ToSlice(pipeline)) // [333]
	for v := range seq.ToSeq(pipeline) {
		fmt.Println(v) // 333
		break
	}
}
```

生成関数の引数から型を推論できるオペレーターは、戻り値をそのまま `Pipe` に渡します。  
型を推論できないオペレーターは、戻り値の `.Type` を `Pipe` に渡し、入力の要素型から補います。  

既存の `iter.Seq[T]` は `FromSeq` で変換できます。  
`Seq[T]` は `iter.Seq[T]` を基にした定義型なので、`for v := range pipeline` でも反復できます。  
`slices.Sorted` など標準の `iter.Seq[T]` を受け取る関数には、`ToSeq` で変換して渡します。  

## Seq2のパイプライン

`Map2` はキーと値を同時に変換し、`Filter2` はキーと値を使って選別します。  
`Take2(n).Type` は、入力からキーと値の型を推論し、先頭から最大n組を取り出します。  
`ToMap` は、キーが `comparable` の場合にmapへ収集します。  
標準の `maps.Collect` を使い、同じキーでは後の値を採用します。  
`ToSeq2` で標準の `iter.Seq2` を取得できます。  
比較できないキーも、`ToSeq2` や直接の反復で扱えます。  
`Seq2[K,V]` も関数型であり、`for k, v := range pipeline` で直接反復できます。  

```go
package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe(seq.Map2(func(k int, v string) (string, int) {
			return strconv.Itoa(k), len(v)
		})).
		Pipe(seq.Filter2(func(k string, v int) bool { return v >= 2 })).
		Pipe(seq.Take2(2).Type)

	fmt.Println(seq.ToMap(pipeline)) // map[1:3 2:2]
	for k, v := range seq.ToSeq2(pipeline) {
		fmt.Println(k, v)
	}

	// FromMapの反復順序はGo標準のmapに従い、保証しません。
	for k, v := range seq.ToSeq2(seq.FromMap(map[string]int{"a": 1, "b": 2})) {
		fmt.Println(k, v)
	}
}
```

## 独自オペレーター

`Pipe` は、イテレーターを変換する関数を直接受け取ります。  
`Seq` では `func(iter.Seq[I]) iter.Seq[O]`、`Seq2` では `func(iter.Seq2[K,V]) iter.Seq2[K2,V2]` です。  

生成関数が専用の関数型を返せば、`Pipe(Multiply(10))` のように接続できます。  
オペレーターの追加に `Seq` / `Seq2` の変更は不要です。  

```go
package main

import (
	"fmt"
	"iter"

	"github.com/meian/go-seqpipeline/seq"
)

// 生成関数は型が確定した専用の関数型を返します。
func Multiply(factor int) multiplyOperator {
	return func(src iter.Seq[int]) iter.Seq[int] {
		return func(yield func(int) bool) {
			if src == nil {
				return
			}
			for v := range src {
				if !yield(v * factor) {
					return
				}
			}
		}
	}
}

type multiplyOperator func(iter.Seq[int]) iter.Seq[int]

func (m multiplyOperator) Apply(src iter.Seq[int]) iter.Seq[int] {
	return m(src)
}

var _ seq.Operator[int, int] = multiplyOperator(nil)

func main() {
	fmt.Println(seq.ToSlice(seq.FromSlice([]int{1, 2, 3}).Pipe(Multiply(10)))) // [10 20 30]
}
```

専用の関数型には、必要に応じて公開の設定メソッドを追加できます。  
型自体が非公開でも、公開メソッドは外部パッケージから呼び出せます。  
入力型が未確定のオペレーターには、型パラメーターを持つ `Type` メソッドを用意し、そのメソッド値を `Pipe` に渡します。  

### 変換関数や既存オペレーターの接続

`OperatorFunc` / `Operator2Func` は、イテレーター変換関数を `Pipe` に接続できる専用の関数型へ包みます。  
`Operator` / `Operator2` インターフェイスは、変換を行う `Apply` メソッドの契約を表します。  
構造体で実装したオペレーターは、`Pipe(op.Apply)` で直接接続できます。  

```go
package main

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	text := seq.OperatorFunc[int, string](func(src iter.Seq[int]) iter.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	fmt.Println(seq.ToSlice(seq.FromSlice([]int{1, 2}).Pipe(text))) // [1 2]
}
```

## 評価と途中終了

パイプラインの構築や `ToSeq` / `ToSeq2` の呼び出しでは、要素を評価しません。  
`ToSlice` / `ToMap` の呼び出し、またはイテレーターの実行時に評価し、入力の反復順序で処理します。  
`yield(false)`、`range` の `break`、`Take` / `Take2` の上限到達で、上流の反復も停止します。  
`Take(n)` / `Take2(n)` は `n <= 0` なら上流を実行しません。  

独自オペレーターのイテレーター変換関数では入力を評価せず、返すイテレーターの実行時に処理してください。  
下流の `yield` が `false` を返したら、上流の反復を停止してください。  

### 再評価と入力の参照

パイプラインを再評価できるかどうかは、元のイテレーターに従います。  
`Take` / `Take2` の残り件数は、評価ごとに初期化します。  
`FromSlice` / `FromMap` は入力をコピーせず、評価時点の要素を参照します。  
`FromMap` の反復順序は保証しません。  

コールバックが状態を保持する場合は、呼び出し側で状態の再利用や並行アクセスを管理してください。  
ライブラリはgoroutineやchannelを使用しません。  
`Seq` と `Seq2` をまたぐ変換や、非同期イベント処理は初期実装の対象外です。  

### nilの扱い

各組み込みオペレーターの `Apply` は、nilの入力イテレーターを空として扱います。  
`ToSeq` / `ToSeq2` はnilをそのまま返します。  
`ToSlice` / `ToMap` は、標準の収集関数と同様にnilの入力でpanicします。  
空のイテレーターを収集した場合は、`ToSlice` はnil、`ToMap` は空の非nilのmapを返します。  

`Seq` / `Seq2` のゼロ値や、nilのイテレーターから変換した値はnilです。  
nil値を直接呼び出したり `range` で反復したりするとpanicします。  

## 開発

次のコマンドで整形と検証を行います。  
`go` が選択したツールチェーンに付属する `gofmt` を使い、Goのバージョンを揃えます。  

```sh
"$(go env GOROOT)/bin/gofmt" -w seq
go vet ./...
go build ./...
go test ./... -count=1
```

CIでは `gofmt`、`go vet`、`go build`、`go test` を個別のジョブで実行します。  
