# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
倍率は Library / Go で、1 を超えると Library が遅い。  
検定の `~` は差が不明確で、増減率の正値は Library の時間・割り当ての増加を表す。  

## パイプ接続

### Compose.Seq：構築・評価

Go：標準 iterator の反復中に同じ値変換・単一値と組の変換を実行。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 20.06 ±0% | 600.15 ±1% | 29.918× |
| 4096 | 1,974 ±2% | 28,075 ±0% | 14.222× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L51) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L11) | 472 ±0% | 15 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L51) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L11) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2891.77% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1322.24% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：構築・評価

Go：標準 iterator の反復中に同じ値変換・単一値と組の変換を実行。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 20.055 ±0% | 338.5 ±0% | 16.879× |
| 4096 | 1,954 ±1% | 16,752.5 ±1% | 8.573× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L88) | 280 ±0% | 9 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L88) | 280 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1587.86% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +757.34% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：構築・評価

Go：標準 iterator の反復中に同じ値変換・単一値と組の変換を実行。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L202) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L162) | 倍率 |
| --- | --- | --- | --- |
| 64 | 21.82 ±1% | 598.4 ±0% | 27.424× |
| 4096 | 1,992.5 ±2% | 29,481.5 ±0% | 14.796× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L202) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L162) | 472 ±0% | 15 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L202) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L162) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2642.44% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1379.62% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：構築・評価

Go：標準 iterator の反復中に同じ値変換・単一値と組の変換を実行。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L277) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L239) | 倍率 |
| --- | --- | --- | --- |
| 64 | 20.96 ±1% | 337.45 ±1% | 16.100× |
| 4096 | 1,975.5 ±2% | 17,366.5 ±0% | 8.791× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L277) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L239) | 280 ±0% | 9 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L277) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/compose_test.go#L239) | 280 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1509.97% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +779.09% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq.Pipe：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.32 ±1% | 17.545 ±1% | 1.013× |
| 4096 | 1,017 ±4% | 1,023 ±2% | 1.006× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.30% | p=0.005 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.425 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：`slices.Values` の値と変換した値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.655 ±1% | 18.52 ±0% | 0.993× |
| 4096 | 1,149.5 ±5% | 1,157 ±8% | 1.007× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.72% | p=0.005 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.288 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：`slices.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L202) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L163) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.625 ±1% | 16.86 ±2% | 1.014× |
| 4096 | 973.15 ±2% | 963.2 ±2% | 0.990× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.41% | p=0.022 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.353 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L278) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L239) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.65 ±0% | 18.565 ±1% | 0.995× |
| 4096 | 1,212 ±3% | 1,205.5 ±1% | 0.995× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.46% | p=0.024 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=1.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 変換・選別・射影

### Filter：構築・評価

Go：`slices.Values` で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.46 ±1% | 28.68 ±1% | 0.857× |
| 4096 | 2,035 ±0% | 1,695.5 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -14.29% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.68% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`slices.All` で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L128) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L89) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.435 ±0% | 28.705 ±0% | 0.859× |
| 4096 | 2,037.5 ±0% | 1,697 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L128) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L89) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L128) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/filter_test.go#L89) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -14.15% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.71% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`maps.All` からキーを取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 317.15 ±9% | 304.7 ±7% | 0.961× |
| 4096 | 22,049 ±2% | 21,852.5 ±4% | 0.991× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/keys_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.529 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=1.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.465 ±2% | 17.53 ±1% | 1.004× |
| 4096 | 1,021.5 ±7% | 1,021 ±3% | 1.000× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.128 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.754 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.71 ±1% | 18.595 ±1% | 0.994× |
| 4096 | 1,177.5 ±2% | 1,205.5 ±0% | 1.024× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/map_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.183 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +2.38% | p=0.003 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：`maps.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 320.9 ±9% | 310.55 ±10% | 0.968× |
| 4096 | 22,085.5 ±2% | 22,124 ±3% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/values_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.579 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.986 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,746 ±2% | 6,884 ±4% | 1.198× |
| 4096 | 343,505.5 ±2% | 410,107.5 ±2% | 1.194× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.81% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +19.39% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,585.5 ±1% | 7,089 ±2% | 1.269× |
| 4096 | 353,994 ±4% | 417,820.5 ±1% | 1.180× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +26.92% | p=0.000 n=10 |
| 64 | B/op | +90.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +18.03% | p=0.000 n=10 |
| 4096 | B/op | +90.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：`maps.All` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 317.25 ±5% | 326.8 ±3% | 1.030× |
| 4096 | 22,255 ±3% | 22,189 ±2% | 0.997× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +3.01% | p=0.023 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.912 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Signed：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.305 ±8% | 101.2 ±1% | 6.207× |
| 4096 | 928.65 ±1% | 4,808 ±1% | 5.177× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +520.67% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +417.74% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Signed：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.835 ±1% | 103.35 ±2% | 6.527× |
| 4096 | 927.55 ±0% | 4,876.5 ±9% | 5.257× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +552.67% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +425.74% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Unsigned：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.095 ±1% | 107.25 ±1% | 6.664× |
| 4096 | 937.65 ±1% | 5,043 ±1% | 5.378× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +566.36% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +437.83% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Unsigned：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.745 ±2% | 106.5 ±3% | 6.764× |
| 4096 | 926.8 ±0% | 5,083.5 ±0% | 5.485× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +576.41% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +448.50% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.745 ±2% | 16.725 ±2% | 0.999× |
| 4096 | 970.25 ±2% | 971.85 ±1% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.853 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.699 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.12 ±1% | 17.155 ±1% | 1.002× |
| 4096 | 1,070 ±2% | 1,076 ±3% | 1.006× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.209 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.591 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L48) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.77 ±1% | 16.68 ±2% | 0.995× |
| 4096 | 969.3 ±2% | 959.95 ±2% | 0.990× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.670 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.165 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L52) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,235.5 ±1% | 2,335 ±0% | 1.045× |
| 4096 | 146,740.5 ±1% | 151,260 ±1% | 1.031× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L52) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L13) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L52) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_map_test.go#L13) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +4.45% | p=0.000 n=10 |
| 64 | B/op | +0.86% | p=0.000 n=10 |
| 64 | allocs/op | +18.18% | p=0.000 n=10 |
| 4096 | sec/op | +3.08% | p=0.000 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +4.17% | p=0.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.915 ±1% | 17.14 ±1% | 1.013× |
| 4096 | 981.65 ±2% | 985.55 ±1% | 1.004× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.33% | p=0.024 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.105 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.055 ±1% | 17.18 ±1% | 1.007× |
| 4096 | 1,074 ±3% | 1,076 ±1% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.288 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.781 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 153.55 ±2% | 255.6 ±2% | 1.665× |
| 4096 | 9,693.5 ±6% | 20,281 ±4% | 2.092× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L51) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L12) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L51) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/to_slice_test.go#L12) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +66.46% | p=0.000 n=10 |
| 64 | B/op | +6.30% | p=0.000 n=10 |
| 64 | allocs/op | +42.86% | p=0.000 n=10 |
| 4096 | sec/op | +109.22% | p=0.000 n=10 |
| 4096 | B/op | +0.05% | p=0.000 n=10 |
| 4096 | allocs/op | +18.75% | p=0.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除き `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.3 ±1% | 30.365 ±1% | 1.985× |
| 4096 | 982.45 ±0% | 1,942.5 ±1% | 1.977× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +98.46% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +97.72% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き `slices.All` で元の添字・値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.825 ±1% | 30.58 ±2% | 1.932× |
| 4096 | 1,155 ±0% | 1,940.5 ±2% | 1.680× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/skip_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +93.24% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +68.01% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件を `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.0605 ±13% | 3.4655 ±0% | 1.132× |
| 4096 | 2.74 ±21% | 3.5335 ±1% | 1.290× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +13.23% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +28.96% | p=0.001 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件を `slices.All` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.035 ±11% | 3.471 ±1% | 1.144× |
| 4096 | 3.1065 ±12% | 3.399 ±0% | 1.094× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/take_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14.37% | p=0.002 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +9.42% | p=0.005 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L128) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.67 ±1% | 1,651 ±1% | 64.316× |
| 4096 | 1,569.5 ±2% | 92,910.5 ±2% | 59.198× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L128) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L88) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L128) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L88) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6331.63% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5819.75% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.99 ±1% | 2,859.5 ±1% | 168.305× |
| 4096 | 1,015 ±3% | 170,165 ±2% | 167.650× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L51) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L11) | 376 ±0% | 11 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L51) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L11) | 376 ±0% | 11 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16730.49% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +16665.02% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L293) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L252) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.875 ±2% | 1,640.5 ±4% | 63.401× |
| 4096 | 1,566 ±2% | 92,157.5 ±1% | 58.849× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L293) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L252) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L293) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L252) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6240.10% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5784.90% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L210) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L170) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.69 ±1% | 1,642 ±1% | 63.916× |
| 4096 | 1,559 ±2% | 92,189 ±2% | 59.133× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L210) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L170) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L210) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L170) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6291.59% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5813.34% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L453) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L412) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.73 ±2% | 1,651.5 ±2% | 64.186× |
| 4096 | 1,560 ±2% | 92,089 ±2% | 59.031× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L453) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L412) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L453) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L412) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6318.58% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5803.14% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L375) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L335) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.845 ±2% | 2,848.5 ±1% | 169.101× |
| 4096 | 995.3 ±5% | 169,691 ±1% | 170.492× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L375) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L335) | 440 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L375) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L335) | 440 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16810.06% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +16949.23% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L618) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L577) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.56 ±1% | 2,785 ±1% | 108.959× |
| 4096 | 1,561.5 ±2% | 160,167.5 ±2% | 102.573× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L618) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L577) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L618) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L577) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +10795.93% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +10157.28% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L535) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L495) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.84 ±1% | 1,645 ±2% | 63.661× |
| 4096 | 1,566 ±1% | 92,731 ±1% | 59.215× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L535) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L495) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L535) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/zip_test.go#L495) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6266.10% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5821.52% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
