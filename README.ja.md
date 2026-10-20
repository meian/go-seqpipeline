# go-seqpipeline

[English](README.md) | 日本語  

`iter.Seq` / `iter.Seq2` を使い、型安全なパイプラインをメソッドチェーンで構築するライブラリです。  

## 環境要件とインストール

Go 1.27.1 以降が必要です。  

```sh
go get github.com/meian/go-seqpipeline
```

## 使い方

`From` で始まる関数から `seq.Seq` / `seq.Seq2` を作り、`Pipe` / `Pipe2` でオペレーターを接続します。  
`seq.Seq[T]` / `seq.Seq2[K, V]` は、`iter.Seq` / `iter.Seq2` に接続メソッドを追加した型です。  
どちらの入力でも、`Pipe` は `seq.Seq`、`Pipe2` は `seq.Seq2` を返します。  
オペレーターの入出力は標準イテレーター型で、`Pipe` / `Pipe2` への接続時に型が推論されます。  
結果は `for range` で直接反復するか、`ToSlice` / `ToMap` でコレクションへ収集します。  

### 単一値のパイプライン

```go
pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
    Pipe(op.Map(func(v int) string { return strconv.Itoa(v) })).
    Pipe(op.Filter(func(v string) bool { return len(v) > 2 })).
    Pipe(op.Take(1).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [333]
```

### キーと値のパイプライン

```go
pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
    Pipe2(op.Map2(func(k int, v string) (string, int) {
        return strconv.Itoa(k), len(v)
    })).
    Pipe2(op.Filter2(func(k string, v int) bool { return v > 1 })).
    Pipe2(op.Take(1).Operator2)
for k, v := range pairs {
    fmt.Println(k, v) // 1 3
}
```

### 型推論

コールバックなどの引数が型情報を持つ場合は、生成関数の戻り値をそのまま `Pipe` / `Pipe2` に渡します。  
生成時の引数に型情報を持たない処理は、設定値のメソッドで入力から型を推論します。  
単一値には `.Operator` を `Pipe` に、キーと値の組には `.Operator2` を `Pipe2` に渡すため、型引数の明示は不要です。  

## 利用上の注意

`Cache` / `Cache2` は入力への適用時に全件評価するため、無限入力はキャッシュより前で制限してください。  
詳細は [キャッシュの注意点](doc/guide.ja.md#評価結果の再利用) を参照してください。  
channel の受信を途中でやめる場合は、渡したコンテキストをキャンセルしてください。  
通信待ちやキャンセル競合の注意は [channel との相互変換](doc/guide.ja.md#channel-との相互変換) を参照してください。  

## 参照先

- [利用ガイド](doc/guide.ja.md)：評価・停止・再評価・nil、入出力、接続型、合成、独自オペレーター
- GoDoc：[seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) の個別 API の契約と境界条件
- Go Example：[seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples) の利用例（[seq のソース](seq/example_test.go)、[op のソース](op/example_test.go)）
- [性能概要](benchmarks/REPORT.md)：測定結果と適用条件

## ライセンス

[MIT License](LICENSE)  
