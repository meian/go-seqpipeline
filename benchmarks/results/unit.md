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

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.64 ±0% | 28.995 ±1% | 0.862× |
| 4096 | 2,048 ±0% | 1,705.5 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -13.81% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.72% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`slices.All` で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.76 ±1% | 28.895 ±0% | 0.856× |
| 4096 | 2,048 ±0% | 1,706 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/filter_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -14.41% | p=0.002 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.70% | p=0.001 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`maps.All` からキーを取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 313.65 ±6% | 307.6 ±8% | 0.981× |
| 4096 | 21,947 ±2% | 21,981.5 ±4% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.971 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.853 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.62 ±0% | 17.85 ±1% | 1.013× |
| 4096 | 1,031.5 ±1% | 1,033 ±1% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.31% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.342 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L125) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L86) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.8 ±1% | 18.72 ±0% | 0.996× |
| 4096 | 1,191.5 ±2% | 1,214.5 ±0% | 1.019× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L125) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L86) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L125) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/map_test.go#L86) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.43% | p=0.002 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +1.93% | p=0.006 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：`maps.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 312.95 ±8% | 322.9 ±9% | 1.032× |
| 4096 | 21,914.5 ±2% | 22,112.5 ±1% | 1.009× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.353 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.631 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,757 ±3% | 6,841.5 ±2% | 1.188× |
| 4096 | 346,683 ±2% | 407,304.5 ±2% | 1.175× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +18.84% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +17.49% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,603 ±2% | 7,033.5 ±2% | 1.255× |
| 4096 | 352,741 ±4% | 421,658 ±2% | 1.195× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +25.53% | p=0.000 n=10 |
| 64 | B/op | +90.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +19.54% | p=0.000 n=10 |
| 4096 | B/op | +90.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：`maps.All` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 310.95 ±4% | 322.7 ±3% | 1.038× |
| 4096 | 22,007 ±2% | 22,112.5 ±1% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +3.78% | p=0.011 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.579 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.525 ±2% | 17.57 ±2% | 1.003× |
| 4096 | 1,014 ±2% | 1,019 ±2% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.270 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.183 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.59 ±2% | 17.555 ±1% | 0.998× |
| 4096 | 1,097.5 ±3% | 1,099.5 ±1% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.897 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.591 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L48) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.385 ±3% | 17.435 ±3% | 1.003× |
| 4096 | 1,007 ±1% | 1,012.5 ±1% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.617 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.754 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L52) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,280.5 ±1% | 2,371.5 ±1% | 1.040× |
| 4096 | 148,232 ±1% | 152,910.5 ±1% | 1.032× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L52) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L13) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L52) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_map_test.go#L13) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +3.99% | p=0.000 n=10 |
| 64 | B/op | +0.86% | p=0.000 n=10 |
| 64 | allocs/op | +18.18% | p=0.000 n=10 |
| 4096 | sec/op | +3.16% | p=0.000 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +4.17% | p=0.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.435 ±1% | 17.59 ±2% | 1.009× |
| 4096 | 1,012 ±1% | 1,018.5 ±2% | 1.006× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +0.89% | p=0.041 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.697 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.605 ±1% | 17.58 ±1% | 0.999× |
| 4096 | 1,096.5 ±1% | 1,094 ±2% | 0.998× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.446 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.868 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 156.05 ±1% | 257.05 ±2% | 1.647× |
| 4096 | 9,609 ±6% | 20,170 ±4% | 2.099× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L51) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L12) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L51) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/to_slice_test.go#L12) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +64.72% | p=0.000 n=10 |
| 64 | B/op | +6.30% | p=0.000 n=10 |
| 64 | allocs/op | +42.86% | p=0.000 n=10 |
| 4096 | sec/op | +109.91% | p=0.000 n=10 |
| 4096 | B/op | +0.05% | p=0.000 n=10 |
| 4096 | allocs/op | +18.75% | p=0.000 n=10 |

</details>

## パイプ接続

### Seq.Pipe：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.615 ±1% | 17.835 ±1% | 1.012× |
| 4096 | 1,036 ±2% | 1,031.5 ±1% | 0.996× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.25% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.323 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L86) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.38 ±1% | 4,989.5 ±1% | 287.083× |
| 4096 | 1,025 ±2% | 298,787.5 ±2% | 291.500× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L86) | 464 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L86) | 464 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28608.29% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29050.00% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L202) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L163) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.335 ±1% | 17.38 ±1% | 1.003× |
| 4096 | 1,004.5 ±2% | 1,007 ±2% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.697 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.643 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L278) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L239) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.825 ±0% | 18.685 ±0% | 0.993× |
| 4096 | 1,193.5 ±4% | 1,212 ±0% | 1.016× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.74% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +1.55% | p=0.021 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除き `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.025 ±1% | 31.375 ±2% | 1.958× |
| 4096 | 1,019.5 ±1% | 2,003 ±2% | 1.965× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +95.79% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +96.47% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き `slices.All` で元の添字・値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.12 ±1% | 31.62 ±2% | 1.962× |
| 4096 | 1,163 ±0% | 2,019 ±2% | 1.736× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/skip_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +96.15% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +73.60% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件を `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.2035 ±10% | 3.483 ±0% | 1.087× |
| 4096 | 3.385 ±2% | 3.5325 ±0% | 1.044× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +8.72% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +4.36% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件を `slices.All` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.178 ±7% | 3.491 ±1% | 1.098× |
| 4096 | 3.3655 ±12% | 3.437 ±1% | 1.021× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/take_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +9.85% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.342 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.3 ±1% | 3,935.5 ±2% | 149.639× |
| 4096 | 1,595 ±0% | 232,511.5 ±2% | 145.775× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L87) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L87) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14863.88% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14477.52% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.565 ±2% | 5,032.5 ±2% | 286.507× |
| 4096 | 1,031 ±2% | 303,154 ±2% | 294.039× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L10) | 464 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L10) | 464 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28550.73% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29303.88% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L292) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L251) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.24 ±1% | 3,925.5 ±2% | 149.600× |
| 4096 | 1,596.5 ±1% | 231,413.5 ±2% | 144.951× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L292) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L251) | 616 ±0% | 17 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L292) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L251) | 616 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14859.98% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14395.05% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L209) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L169) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.31 ±1% | 3,907.5 ±1% | 148.518× |
| 4096 | 1,594.5 ±1% | 232,351.5 ±2% | 145.721× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L209) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L169) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L209) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L169) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14751.77% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14472.06% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L452) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L411) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.46 ±0% | 3,913.5 ±2% | 147.902× |
| 4096 | 1,598 ±2% | 230,694.5 ±1% | 144.365× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L452) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L411) | 616 ±0% | 17 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L452) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L411) | 616 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14690.25% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14336.45% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L374) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L334) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.49 ±23% | 5,018.5 ±1% | 286.935× |
| 4096 | 1,026 ±2% | 300,996.5 ±1% | 293.369× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L374) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L334) | 568 ±0% | 16 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L374) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L334) | 568 ±0% | 16 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28593.54% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29236.89% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L617) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L576) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.46 ±1% | 3,953.5 ±2% | 149.414× |
| 4096 | 1,595.5 ±1% | 231,918.5 ±1% | 145.358× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L617) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L576) | 616 ±0% | 17 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L617) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L576) | 616 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14841.42% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14435.79% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L534) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L494) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.235 ±2% | 3,960 ±2% | 150.943× |
| 4096 | 1,591 ±1% | 232,691.5 ±1% | 146.255× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L494) | 616 ±0% | 17 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L534) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/zip_test.go#L494) | 616 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14994.34% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14525.49% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
