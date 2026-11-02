# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
Compose の Pipe は個別接続、Composed は合成したオペレーターを接続する処理。  
倍率は各表の右側の実装の時間 / 左側の実装の時間で、1 を超えると右側が遅い。  
検定の `~` は差が不明確で、増減率の正値は右側の時間・割り当ての増加を表す。  

## パイプ接続

### Compose.Seq：構築・評価

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L26) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 504.1 ±4% | 600.6 ±0% | 1.191× |
| 4096 | 31,914.5 ±0% | 32,271.5 ±0% | 1.011× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L26) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L12) | 472 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L26) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L12) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.14% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.12% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：構築・評価

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L101) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 523.9 ±1% | 605.85 ±0% | 1.156× |
| 4096 | 32,554.5 ±0% | 33,530 ±0% | 1.030× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L101) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L87) | 472 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L101) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L87) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +15.64% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +3.00% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：構築・評価

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L176) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L162) | 倍率 |
| --- | --- | --- | --- |
| 64 | 507.6 ±0% | 599.8 ±0% | 1.182× |
| 4096 | 31,843 ±0% | 32,187 ±0% | 1.011× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L176) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L162) | 472 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L176) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L162) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +18.16% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.08% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：構築・評価

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L251) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L237) | 倍率 |
| --- | --- | --- | --- |
| 64 | 508.95 ±0% | 599.05 ±0% | 1.177× |
| 4096 | 31,908.5 ±1% | 32,240 ±0% | 1.010× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L251) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L237) | 472 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L251) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/compose_test.go#L237) | 472 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +17.70% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.04% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq.Pipe：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.29 ±2% | 17.5 ±1% | 1.012× |
| 4096 | 1,010.5 ±2% | 1,061.5 ±2% | 1.050× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.21% | p=0.008 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +5.05% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：`slices.Values` の値と変換した値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L90) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L80) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.69 ±1% | 18.56 ±1% | 0.993× |
| 4096 | 1,151.5 ±0% | 1,155.5 ±0% | 1.003× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L90) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L80) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L90) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L80) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.70% | p=0.011 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.070 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：`slices.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L158) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L148) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.705 ±2% | 16.635 ±2% | 0.996× |
| 4096 | 962.65 ±0% | 963.75 ±1% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L158) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L148) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L158) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L148) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.643 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.739 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L226) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L216) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.735 ±1% | 18.635 ±1% | 0.995× |
| 4096 | 1,180.5 ±1% | 1,130 ±2% | 0.957× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L226) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L216) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L226) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipe_test.go#L216) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.53% | p=0.010 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -4.28% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 変換・選別・射影

### Filter：構築・評価

Go：`slices.Values` で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 36.545 ±9% | 29.04 ±1% | 0.795× |
| 4096 | 2,059 ±0% | 1,696 ±0% | 0.824× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -20.54% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -17.63% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`slices.All` で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L92) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L82) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.415 ±0% | 28.73 ±1% | 0.860× |
| 4096 | 2,039 ±0% | 1,696 ±0% | 0.832× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L92) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L82) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L92) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/filter_test.go#L82) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -14.02% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.82% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`maps.All` からキーを取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 316.5 ±5% | 326.75 ±4% | 1.032× |
| 4096 | 21,972 ±1% | 21,963 ±1% | 1.000× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/keys_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.123 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=1.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.26 ±1% | 17.615 ±2% | 1.021× |
| 4096 | 1,059 ±1% | 1,043 ±2% | 0.985× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2.06% | p=0.003 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -1.51% | p=0.003 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L90) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L80) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.725 ±1% | 18.62 ±1% | 0.994× |
| 4096 | 1,182.5 ±1% | 1,205 ±1% | 1.019× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L90) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L80) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L90) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/map_test.go#L80) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.065 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +1.90% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：`maps.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 315.4 ±2% | 323 ±6% | 1.024× |
| 4096 | 21,976 ±3% | 22,437 ±3% | 1.021× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/values_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.247 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.481 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L23) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,921 ±1% | 5,986.5 ±1% | 1.011× |
| 4096 | 359,936.5 ±3% | 362,523 ±2% | 1.007× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L23) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L12) | 336 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L23) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_channel_test.go#L12) | 336 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.11% | p=0.030 n=10 |
| 64 | B/op | +23.53% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | ~ | p=0.436 n=10 |
| 4096 | B/op | +23.53% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,877.5 ±1% | 6,211 ±1% | 1.057× |
| 4096 | 356,331.5 ±2% | 377,341.5 ±2% | 1.059× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L22) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L12) | 264 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L22) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_channel_test.go#L12) | 264 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +5.67% | p=0.000 n=10 |
| 64 | B/op | +65.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +5.90% | p=0.000 n=10 |
| 4096 | B/op | +65.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：`maps.All` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L20) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 316.5 ±9% | 316.5 ±3% | 1.000× |
| 4096 | 22,019.5 ±1% | 21,882 ±1% | 0.994× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L20) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L20) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_map_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=1.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.448 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Signed：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.215 ±2% | 17.22 ±1% | 1.062× |
| 4096 | 929.85 ±2% | 971.3 ±2% | 1.045× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6.20% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +4.46% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Signed：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.035 ±1% | 17.195 ±2% | 1.072× |
| 4096 | 928.6 ±1% | 958.25 ±1% | 1.032× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L20) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +7.23% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +3.19% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Unsigned：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.125 ±5% | 16.965 ±3% | 1.052× |
| 4096 | 928 ±0% | 967.5 ±2% | 1.043× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +5.21% | p=0.018 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +4.26% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Unsigned：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16 ±1% | 17.015 ±1% | 1.063× |
| 4096 | 928.2 ±1% | 953.25 ±2% | 1.027× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L101) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_range_test.go#L92) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6.34% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +2.70% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L21) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.625 ±1% | 16.615 ±2% | 0.999× |
| 4096 | 968.2 ±2% | 962.5 ±2% | 0.994× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L21) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L21) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.927 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.271 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L90) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L80) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.955 ±2% | 16.95 ±1% | 1.000× |
| 4096 | 1,085.5 ±2% | 1,082.5 ±1% | 0.997× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L90) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L80) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L90) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_seq_test.go#L80) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.726 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.669 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L20) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.645 ±1% | 16.72 ±3% | 1.005× |
| 4096 | 965.85 ±2% | 963.35 ±1% | 0.997× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L20) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L20) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/from_slice_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.724 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.853 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L24) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L14) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,228.5 ±1% | 2,238.5 ±1% | 1.004× |
| 4096 | 146,444 ±1% | 145,746.5 ±1% | 0.995× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L24) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L14) | 4,648 ±0% | 11 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L24) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_map_test.go#L14) | 296,120 ±0% | 48 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.210 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.529 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L21) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.89 ±2% | 17.065 ±2% | 1.010× |
| 4096 | 975.8 ±1% | 980.1 ±1% | 1.004× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L21) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L21) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.04% | p=0.034 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.149 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L90) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L80) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.1 ±2% | 17.095 ±1% | 1.000× |
| 4096 | 1,092.5 ±3% | 1,092 ±1% | 1.000× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L90) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L80) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L90) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_seq_test.go#L80) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.868 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.697 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L23) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 152.3 ±1% | 153.9 ±3% | 1.011× |
| 4096 | 9,497.5 ±4% | 9,536 ±3% | 1.004× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L23) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L13) | 1,016 ±0% | 7 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L23) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/to_slice_test.go#L13) | 128,248 ±0% | 16 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.753 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.631 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除き `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.3 ±1% | 30.44 ±1% | 1.990× |
| 4096 | 979.3 ±1% | 1,929 ±2% | 1.970× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +98.95% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +96.98% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き `slices.All` で元の添字・値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L91) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L81) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.795 ±1% | 30.55 ±1% | 1.934× |
| 4096 | 1,157 ±1% | 1,954 ±1% | 1.689× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L91) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L81) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L91) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/skip_test.go#L81) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +93.42% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +68.89% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件を `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L22) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2.97 ±10% | 3.543 ±1% | 1.193× |
| 4096 | 3.007 ±5% | 3.55 ±1% | 1.181× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L22) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L12) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L22) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L12) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.29% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +18.06% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件を `slices.All` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L91) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L81) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2.926 ±14% | 3.4325 ±1% | 1.173× |
| 4096 | 2.9725 ±6% | 3.432 ±0% | 1.155× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L91) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L81) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L91) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/take_test.go#L81) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +17.31% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +15.46% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L92) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L81) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.76 ±2% | 1,624.5 ±2% | 63.063× |
| 4096 | 1,641.5 ±1% | 91,969.5 ±2% | 56.028× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L92) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L81) | 432 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L92) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L81) | 432 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6206.29% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5502.77% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L23) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17 ±2% | 2,888.5 ±2% | 169.912× |
| 4096 | 998.8 ±1% | 171,218.5 ±1% | 171.424× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L23) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L12) | 360 ±0% | 11 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L23) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L12) | 360 ±0% | 11 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16891.18% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +17042.42% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L241) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L229) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.63 ±1% | 1,632 ±1% | 63.675× |
| 4096 | 1,633 ±1% | 91,256.5 ±1% | 55.883× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L241) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L229) | 496 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L241) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L229) | 496 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6267.54% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5488.27% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L166) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L155) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.61 ±1% | 1,614.5 ±2% | 63.042× |
| 4096 | 1,638.5 ±1% | 91,299 ±2% | 55.721× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L166) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L155) | 432 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L166) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L155) | 432 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6204.18% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5472.11% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L385) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L373) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.915 ±1% | 1,631.5 ±2% | 62.956× |
| 4096 | 1,635 ±1% | 91,201.5 ±2% | 55.781× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L385) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L373) | 496 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L385) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L373) | 496 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6195.58% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5478.07% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L315) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L304) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.11 ±2% | 2,877 ±3% | 168.147× |
| 4096 | 990.2 ±1% | 170,876.5 ±2% | 172.568× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L315) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L304) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L315) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L304) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16714.73% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +17156.77% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L534) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L522) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.685 ±1% | 2,826.5 ±2% | 110.045× |
| 4096 | 1,632 ±1% | 163,995 ±2% | 100.487× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L522) | 496 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L522) | 496 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +10904.48% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +9948.71% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L459) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L448) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.475 ±1% | 1,641 ±1% | 64.416× |
| 4096 | 1,635.5 ±1% | 92,289 ±1% | 56.429× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L459) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L448) | 496 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L459) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/zip_test.go#L448) | 496 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6341.61% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5542.86% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
