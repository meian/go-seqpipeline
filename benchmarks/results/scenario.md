# 利用場面別の性能

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

倍率は各表の右側の実装の時間 / 左側の実装の時間で、1 を超えると右側が遅い。  
検定の `~` は差が不明確で、増減率の正値は右側の時間・割り当ての増加を表す。  
Fill / Warm は保存費用込みの Reuse と計測範囲が異なる。  

## Pipeline

### Seq / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 234.3 ±5% | 543.05 ±2% | 2.318× |
| 64 | 6 | 333.8 ±5% | 1,143.5 ±2% | 3.426× |
| 4096 | 2 | 13,121 ±6% | 25,241.5 ±2% | 1.924× |
| 4096 | 6 | 18,977.5 ±7% | 56,625.5 ±1% | 2.984× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +131.78% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +242.57% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +92.37% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +198.38% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 235.05 ±3% | 543.05 ±2% | 2.310× |
| 64 | 6 | 304.25 ±3% | 1,143.5 ±2% | 3.758× |
| 4096 | 2 | 13,601.5 ±6% | 25,241.5 ±2% | 1.856× |
| 4096 | 6 | 18,094.5 ±4% | 56,625.5 ±1% | 3.129× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +131.04% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +275.84% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +85.58% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +212.94% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 2,031 ±0% | 6 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 2,031 ±0% | 6 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 177,391 ±0% | 14 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 177,391 ±0% | 14 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.035 ±1% | 197.6 ±2% | 3.657× |
| 64 | 6 | 73.87 ±3% | 382.65 ±1% | 5.180× |
| 4096 | 2 | 55.75 ±2% | 197.75 ±1% | 3.547× |
| 4096 | 6 | 77.76 ±3% | 383.4 ±1% | 4.931× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +265.69% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +418.00% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +254.71% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +393.06% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.905 ±2% | 197.6 ±2% | 3.472× |
| 64 | 6 | 72.425 ±2% | 382.65 ±1% | 5.283× |
| 4096 | 2 | 56.91 ±2% | 197.75 ±1% | 3.475× |
| 4096 | 6 | 72.195 ±3% | 383.4 ±1% | 5.311× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +247.25% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +428.34% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +247.48% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +431.06% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 236 ±3% | 557.05 ±2% | 2.360× |
| 64 | 6 | 334.35 ±2% | 1,198 ±3% | 3.583× |
| 4096 | 2 | 13,338 ±4% | 26,674 ±3% | 2.000× |
| 4096 | 6 | 18,423.5 ±7% | 60,081.5 ±2% | 3.261× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +136.04% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +258.31% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +99.99% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +226.11% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 240.65 ±3% | 557.05 ±2% | 2.315× |
| 64 | 6 | 307.25 ±2% | 1,198 ±3% | 3.899× |
| 4096 | 2 | 14,290 ±6% | 26,674 ±3% | 1.867× |
| 4096 | 6 | 18,103.5 ±3% | 60,081.5 ±2% | 3.319× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +131.48% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +289.91% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +86.66% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +231.88% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.59 ±3% | 197.9 ±1% | 3.625× |
| 64 | 6 | 74.07 ±2% | 379.25 ±2% | 5.120× |
| 4096 | 2 | 55.66 ±1% | 198.6 ±2% | 3.568× |
| 4096 | 6 | 78.24 ±2% | 381 ±1% | 4.870× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +262.52% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +412.02% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +256.81% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +386.96% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.22 ±2% | 197.9 ±1% | 3.520× |
| 64 | 6 | 72.92 ±2% | 379.25 ±2% | 5.201× |
| 4096 | 2 | 57.29 ±1% | 198.6 ±2% | 3.467× |
| 4096 | 6 | 72.81 ±2% | 381 ±1% | 5.233× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +252.01% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +420.09% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +246.66% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +423.28% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 239 ±0% | 3 ±33% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L433) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L264) | 1,024 ±0% | 30 ±0% |

## Channel

### 入力条件別

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,406.5 ±1% | 7,671.5 ±1% | 1.197× |
| 64 | 64 | 2,289 ±3% | 3,406.5 ±2% | 1.488× |
| 4096 | 0 | 389,562 ±1% | 458,205.5 ±1% | 1.176× |
| 4096 | 64 | 114,584.5 ±4% | 187,568.5 ±2% | 1.637× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +19.75% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +48.82% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +17.62% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +63.69% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L192) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,406.5 ±1% | 14,184.5 ±1% | 2.214× |
| 64 | 64 | 2,289 ±3% | 8,794 ±2% | 3.842× |
| 4096 | 0 | 389,562 ±1% | 787,046.5 ±2% | 2.020× |
| 4096 | 64 | 114,584.5 ±4% | 701,894.5 ±3% | 6.126× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +121.41% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +284.19% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +102.03% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +512.56% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### 時間：Seq と ToChannel

| N | 入力容量 | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L17) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L52) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 549 ±2% | 6,714.5 ±2% | 12.230× |
| 4096 | 0 | 26,019 ±2% | 337,299.5 ±1% | 12.964× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +1123.04% | p=0.000 n=10 |
| 64 | 0 | B/op | +6.53% | p=0.000 n=10 |
| 64 | 0 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +1196.36% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L192) | 2,752 ±0% | 19 ±0% |
| 64 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L17) | 2,328 ±0% | 17 ±0% |
| 64 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L52) | 2,480 ±0% | 16 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L192) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L192) | 178,112 ±0% | 27 ±0% |
| 4096 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L17) | 177,688 ±0% | 25 ±0% |
| 4096 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L52) | 177,840.5 ±0% | 24 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 178,488.5 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L192) | 178,640 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 453.65 ±15% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 18,102 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 1,364.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 81,224 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L29) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 93.565 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 4,758 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 93.01 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 4,765.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 212.6 ±1% | 468.5 ±2% | 2.204× |
| 64 | 4 | 777.75 ±0% | 759.05 ±1% | 0.976× |
| 64 | 16 | 3,028.5 ±1% | 1,911.5 ±1% | 0.631× |
| 4096 | 1 | 11,415 ±0% | 23,138.5 ±2% | 2.027× |
| 4096 | 4 | 45,496 ±1% | 37,376 ±2% | 0.822× |
| 4096 | 16 | 181,880.5 ±1% | 94,796.5 ±1% | 0.521× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +120.37% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -2.40% | p=0.000 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -36.88% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +102.70% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -17.85% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.24% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -47.88% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.77% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,568 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,952 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,262.5 ±1% | 1,491 ±2% | 1.181× |
| 64 | 4 | 4,962.5 ±1% | 1,787.5 ±2% | 0.360× |
| 64 | 16 | 19,756.5 ±2% | 2,936 ±1% | 0.149× |
| 4096 | 1 | 77,128 ±2% | 86,685 ±2% | 1.124× |
| 4096 | 4 | 308,096 ±2% | 101,955 ±2% | 0.331× |
| 4096 | 16 | 1,228,119.5 ±3% | 159,229.5 ±2% | 0.130× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +18.10% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.98% | p=0.000 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.14% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +12.39% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -66.91% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.24% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -87.03% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.77% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,568 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 128,952 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 431.15 ±5% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 25,257 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 259,497 ±0% | 23 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 1,423 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 86,299.5 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L371) | 259,496 ±0% | 23 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 95.515 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 5,362.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 95.21 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 5,391.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 212.3 ±1% | 541.15 ±28% | 2.549× |
| 64 | 4 | 774.15 ±0% | 830.85 ±2% | 1.073× |
| 64 | 16 | 3,018 ±2% | 1,983.5 ±1% | 0.657× |
| 4096 | 1 | 11,049 ±1% | 30,849.5 ±2% | 2.792× |
| 4096 | 4 | 43,807.5 ±1% | 47,124 ±2% | 1.076× |
| 4096 | 16 | 175,821 ±2% | 111,986 ±2% | 0.637× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +154.90% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | +7.32% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -34.28% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +179.21% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190735.66% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | +7.57% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.34% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -36.31% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.47% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 259,536.5 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 259,632 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 260,016.5 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,271.5 ±2% | 1,538.5 ±2% | 1.210× |
| 64 | 4 | 5,001.5 ±2% | 1,834 ±2% | 0.367× |
| 64 | 16 | 19,907 ±1% | 2,983.5 ±1% | 0.150× |
| 4096 | 1 | 77,883.5 ±0% | 92,977.5 ±3% | 1.194× |
| 4096 | 4 | 311,552 ±1% | 110,001.5 ±2% | 0.353× |
| 4096 | 16 | 1,247,180.5 ±2% | 174,471.5 ±2% | 0.140× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +21.00% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.33% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.01% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +19.38% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190735.29% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -64.69% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.34% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -86.01% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.39% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 259,536 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 259,632 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 260,016 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
