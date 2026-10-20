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
`seq.Seq[T]` / `seq.Seq2[K, V]` は、標準の `iter.Seq` / `iter.Seq2` に接続メソッドを追加した型です。  
どちらの入力でも、`Pipe` は `seq.Seq`、`Pipe2` は `seq.Seq2` を返します。  
オペレーターの入出力は標準イテレーター型で、`Pipe` / `Pipe2` への接続時に型が推論されます。  
結果は `for range` で直接反復するか、`ToSlice` / `ToMap` でコレクションへ収集します。  

### 単一値のパイプライン

次の例は `fmt`、`strconv`、`github.com/meian/go-seqpipeline/seq` をインポートして使います。  

```go
pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
    Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
    Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
    Pipe(seq.Take(1).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [333]
```

### キーと値のパイプライン

次の例は `fmt`、`slices`、`strconv`、`github.com/meian/go-seqpipeline/seq` をインポートして使います。  

```go
pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
    Pipe2(seq.Map2(func(k int, v string) (string, int) {
        return strconv.Itoa(k), len(v)
    })).
    Pipe2(seq.Filter2(func(k string, v int) bool { return v > 1 })).
    Pipe2(seq.Take(1).Operator2)
for k, v := range pairs {
    fmt.Println(k, v) // 1 3
}
```

### 型推論

コールバックなどの引数が型情報を持つ場合は、生成関数の戻り値をそのまま `Pipe` / `Pipe2` に渡します。  
生成時の引数に型情報を持たない処理は、設定値のメソッドで入力から型を推論します。  
単一値には `.Operator` を `Pipe` に、キーと値の組には `.Operator2` を `Pipe2` に渡すため、型引数の明示は不要です。  

## 利用上の注意

`Cache` / `Cache2` は接続時に入力を全件評価するため、無限入力はキャッシュより前で制限してください。  
詳細は [キャッシュの注意点](doc/guide.ja.md#評価結果の再利用) を参照してください。  
channel の受信を途中でやめる場合は、渡したコンテキストをキャンセルしてください。  
通信待ちやキャンセル競合の注意は [channel との相互変換](doc/guide.ja.md#channel-との相互変換) を参照してください。  

## 参照先

- [利用ガイド](doc/guide.ja.md)：評価・停止・再評価・nil、入出力、接続型、独自オペレーター
- [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq)：API ごとの仕様と境界条件
- [Go Example](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples)：期待出力を持つ利用例（[ソース](seq/example_test.go)）
- [性能概要](benchmarks/REPORT.md)：測定結果と適用条件

## ライセンス

[MIT License](LICENSE)  
