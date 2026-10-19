# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
倍率は Library / Go で、1 を超えると Library が遅い。  
検定の `~` は差が不明確で、増減率の正値は Library の時間・割り当ての増加を表す。  

## 変換・選別・射影

### Filter：構築・評価

Go：`slices.Values` で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.43 ±11% | 28.82 ±0% | 0.862× |
| 4096 | 2,037 ±0% | 1,696.5 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -13.79% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.72% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`slices.All` で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.43 ±0% | 28.775 ±1% | 0.861× |
| 4096 | 2,039.5 ±0% | 1,697 ±0% | 0.832× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/filter_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -13.92% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.79% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`maps.All` からキーを取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 308.35 ±5% | 321.3 ±9% | 1.042× |
| 4096 | 21,903 ±3% | 21,550.5 ±2% | 0.984× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.190 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -1.61% | p=0.023 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.525 ±1% | 17.715 ±1% | 1.011× |
| 4096 | 1,025.5 ±9% | 1,024.5 ±1% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.08% | p=0.004 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.467 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L125) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L86) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.7 ±0% | 18.605 ±0% | 0.995× |
| 4096 | 1,181.5 ±1% | 1,206.5 ±0% | 1.021× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L125) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L86) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L125) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/map_test.go#L86) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.51% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +2.12% | p=0.002 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：`maps.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 312.8 ±5% | 317.8 ±9% | 1.016× |
| 4096 | 21,607.5 ±2% | 21,854.5 ±3% | 1.011× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.481 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.404 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,620.5 ±2% | 6,965 ±3% | 1.239× |
| 4096 | 346,226 ±1% | 408,994.5 ±1% | 1.181× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +23.92% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +18.13% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,655 ±1% | 7,283.5 ±3% | 1.288× |
| 4096 | 340,875.5 ±2% | 424,954.5 ±1% | 1.247× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28.80% | p=0.000 n=10 |
| 64 | B/op | +90.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +24.67% | p=0.000 n=10 |
| 4096 | B/op | +90.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：`maps.All` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 316.65 ±4% | 316.3 ±6% | 0.999× |
| 4096 | 21,797 ±5% | 21,779 ±2% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.895 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.869 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.125 ±4% | 17.405 ±2% | 1.016× |
| 4096 | 1,000.9 ±2% | 1,001 ±3% | 1.000× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.101 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.912 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.42 ±1% | 17.5 ±3% | 1.005× |
| 4096 | 1,089 ±4% | 1,067.5 ±2% | 0.980× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.986 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.491 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L48) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.035 ±2% | 16.985 ±2% | 0.997× |
| 4096 | 978.9 ±2% | 986.85 ±2% | 1.008× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.670 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.280 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L52) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,236.5 ±1% | 2,323 ±1% | 1.039× |
| 4096 | 145,304.5 ±2% | 149,569.5 ±1% | 1.029× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L52) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L13) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L52) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_map_test.go#L13) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +3.87% | p=0.000 n=10 |
| 64 | B/op | +0.86% | p=0.000 n=10 |
| 64 | allocs/op | +18.18% | p=0.000 n=10 |
| 4096 | sec/op | +2.94% | p=0.000 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +4.17% | p=0.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.355 ±4% | 17.39 ±1% | 1.002× |
| 4096 | 999.75 ±3% | 1,014 ±1% | 1.014× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.362 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +1.43% | p=0.037 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.435 ±3% | 17.345 ±1% | 0.995× |
| 4096 | 1,094.5 ±1% | 1,077.5 ±2% | 0.984× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.971 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.183 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 152.55 ±4% | 253.55 ±3% | 1.662× |
| 4096 | 9,639 ±8% | 19,948.5 ±4% | 2.070× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L51) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L12) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L51) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/to_slice_test.go#L12) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +66.21% | p=0.000 n=10 |
| 64 | B/op | +6.30% | p=0.000 n=10 |
| 64 | allocs/op | +42.86% | p=0.000 n=10 |
| 4096 | sec/op | +106.96% | p=0.000 n=10 |
| 4096 | B/op | +0.05% | p=0.000 n=10 |
| 4096 | allocs/op | +18.75% | p=0.000 n=10 |

</details>

## パイプ接続

### Seq.Pipe：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.565 ±1% | 17.76 ±2% | 1.011× |
| 4096 | 1,044 ±10% | 1,025.5 ±5% | 0.982× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.072 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.566 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：`slices.Values` の値と変換した値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L133) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L86) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.72 ±0% | 18.855 ±0% | 1.007× |
| 4096 | 1,153 ±9% | 1,162.5 ±9% | 1.008× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L133) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L86) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L133) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L86) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +0.72% | p=0.008 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.183 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：`slices.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L209) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L170) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.065 ±3% | 17.135 ±4% | 1.004× |
| 4096 | 989.3 ±3% | 986.1 ±2% | 0.997× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L209) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L170) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L209) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L170) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.516 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.739 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L285) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L246) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.66 ±0% | 18.595 ±0% | 0.997× |
| 4096 | 1,202 ±2% | 1,203.5 ±0% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L285) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L246) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L285) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipe_test.go#L246) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.35% | p=0.024 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.897 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除き `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.585 ±1% | 30.915 ±3% | 1.984× |
| 4096 | 999.55 ±2% | 1,963.5 ±2% | 1.964× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +98.36% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +96.44% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き `slices.All` で元の添字・値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.985 ±1% | 31.035 ±3% | 1.942× |
| 4096 | 1,156.5 ±0% | 1,981 ±3% | 1.713× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/skip_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +94.15% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +71.29% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件を `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2.855 ±20% | 3.465 ±0% | 1.214× |
| 4096 | 3.347 ±11% | 3.516 ±5% | 1.050× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +21.37% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +5.05% | p=0.001 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件を `slices.All` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.0955 ±12% | 3.459 ±0% | 1.117× |
| 4096 | 3.406 ±11% | 3.412 ±2% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/take_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +11.74% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.565 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.175 ±2% | 1,644.5 ±3% | 62.827× |
| 4096 | 1,589 ±1% | 92,084.5 ±2% | 57.951× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L87) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L87) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6182.71% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5695.12% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.33 ±2% | 2,815.5 ±3% | 162.464× |
| 4096 | 1,019.5 ±2% | 166,876.5 ±4% | 163.685× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L10) | 376 ±0% | 11 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L10) | 376 ±0% | 11 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16146.39% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +16268.46% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L292) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L251) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.215 ±1% | 1,623.5 ±2% | 61.930× |
| 4096 | 1,586 ±1% | 91,038.5 ±3% | 57.401× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L292) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L251) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L292) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L251) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6093.02% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5640.13% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L209) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L169) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.235 ±2% | 1,610.5 ±2% | 61.387× |
| 4096 | 1,587.5 ±1% | 90,958.5 ±1% | 57.297× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L209) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L169) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L209) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L169) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6038.75% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5629.67% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L452) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L411) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.01 ±1% | 1,629.5 ±3% | 62.649× |
| 4096 | 1,588.5 ±0% | 91,016.5 ±5% | 57.297× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L452) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L411) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L452) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L411) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6164.90% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5629.71% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L374) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L334) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.03 ±3% | 2,800 ±2% | 164.416× |
| 4096 | 1,008.5 ±1% | 164,113.5 ±1% | 162.730× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L374) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L334) | 440 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L374) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L334) | 440 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16341.57% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +16173.03% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L617) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L576) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.06 ±2% | 2,724 ±1% | 104.528× |
| 4096 | 1,588 ±2% | 156,402 ±3% | 98.490× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L617) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L576) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L617) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L576) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +10352.80% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +9748.99% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L534) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L494) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.175 ±2% | 1,618.5 ±1% | 61.834× |
| 4096 | 1,591.5 ±1% | 90,470 ±2% | 56.846× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L494) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/zip_test.go#L494) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6083.38% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5584.57% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
