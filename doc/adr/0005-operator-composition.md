# ADR-0005: Operator 関数型と合成専用型の役割を分ける

- 決定日: 2026-10-21
- 状態: 採用
- 日付の根拠: [PR #58 のレビュー合意](https://github.com/meian/go-seqpipeline/pull/58#discussion_r4328913359) の日付（日本時間）

## 背景

[Issue #51](https://github.com/meian/go-seqpipeline/issues/51) では、4 種類の Operator 関数型に `Pipe` / `Pipe2` を持たせ、既存のオペレーターからも合成を始める案だった。  
合成結果も Operator 関数型として `seq` 側に直接渡す想定だった。  
しかし [PR #58 のレビュー](https://github.com/meian/go-seqpipeline/pull/58#discussion_r4328832967) で、Operator と Seq の双方に `Pipe` があると、通常の入力への接続と再利用する処理の合成の使い分けが曖昧になると指摘された。  

## 決定

`Operator` / `Operator2` / `ToSeqOperator` / `ToSeq2Operator` は変換関数型とし、`Pipe` / `Pipe2` を持たせない。  
通常の入力には `seq.Seq` / `seq.Seq2` の `Pipe` / `Pipe2` で処理を接続する。  
入力へ適用する前に再利用可能な処理を組み立てる場合は、`op.Compose()` / `op.Compose2()` から得る合成専用型の `Pipe` / `Pipe2` を使う。  
合成値を入力へ適用するときは、その `.Operator` メソッド値を `seq` 側の `Pipe` / `Pipe2` に渡す。  

## 理由

入力への接続と入力なしの合成を、`Pipe` / `Pipe2` の受け手で区別できる。  
Operator 関数型は変換だけを担い、合成の途中で元の入力型を保つ役割は合成専用型に集められる。  

## 採用しなかった選択肢

- 4 種類の Operator 関数型にも `Pipe` / `Pipe2` を持たせる案: `op.Map(...).Pipe(...)` から合成を始められるが、`seq` 側の接続メソッドとの使い分けが曖昧になる

## 影響

合成は `op.Compose()` / `op.Compose2()` から始め、`op.Map(...).Pipe(...)` では始めない。  
合成結果を適用する記法は、当初の `Pipe(transform)` から `Pipe(transform.Operator)` に変わる。  
単一値とキー・値の組をまたぐ場合を含む接続方法は [利用ガイド](../guide.ja.md#compose-によるオペレーターの合成) と [英語版](../guide.md#composing-operators-with-compose) を参照する。  

## 関連 ADR

- [ADR-0001](0001-operator-package.md): オペレーターの配置と `seq` / `op` の責務

## 根拠

- [Issue #51](https://github.com/meian/go-seqpipeline/issues/51)
- [PR #58 の変更提案](https://github.com/meian/go-seqpipeline/pull/58#discussion_r4328876425)・[合意](https://github.com/meian/go-seqpipeline/pull/58#discussion_r4328913359)
- [Issue #71](https://github.com/meian/go-seqpipeline/issues/71)
