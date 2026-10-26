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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 238.5 ±2% | 545.75 ±0% | 2.288× |
| 64 | 6 | 332.75 ±2% | 1,139 ±0% | 3.423× |
| 4096 | 2 | 13,579.5 ±1% | 26,350 ±1% | 1.940× |
| 4096 | 6 | 19,271 ±2% | 56,441.5 ±1% | 2.929× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +128.83% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +242.30% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +94.04% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +192.88% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 243.25 ±2% | 545.75 ±0% | 2.244× |
| 64 | 6 | 307.45 ±2% | 1,139 ±0% | 3.705× |
| 4096 | 2 | 13,988 ±1% | 26,350 ±1% | 1.884× |
| 4096 | 6 | 18,763 ±1% | 56,441.5 ±1% | 3.008× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +124.36% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +270.47% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +88.38% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +200.81% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 2,032 ±0% | 7 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 2,032 ±0% | 7 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 2,032 ±0% | 7 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.835 ±1% | 196.35 ±1% | 3.581× |
| 64 | 6 | 73.41 ±1% | 384.75 ±0% | 5.241× |
| 4096 | 2 | 56.155 ±1% | 196.75 ±0% | 3.504× |
| 4096 | 6 | 76.925 ±0% | 385.65 ±0% | 5.013× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +258.07% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +424.11% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +250.37% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +401.33% | p=0.000 n=10 |
| 4096 | 6 | B/op | +327.56% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +757.14% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 57.375 ±1% | 196.35 ±1% | 3.422× |
| 64 | 6 | 72.785 ±1% | 384.75 ±0% | 5.286× |
| 4096 | 2 | 56.13 ±1% | 196.75 ±0% | 3.505× |
| 4096 | 6 | 71.705 ±1% | 385.65 ±0% | 5.378× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +242.22% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +428.61% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +250.53% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +437.83% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 240 ±0% | 4 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 240 ±0% | 4 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 240 ±0% | 4 ±25% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 240 ±0% | 4 ±25% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L180) | 239.5 ±0% | 3.5 ±14% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L95) | 240 ±0% | 4 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L11) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 241.15 ±1% | 560 ±2% | 2.322× |
| 64 | 6 | 336.6 ±1% | 1,200.5 ±0% | 3.567× |
| 4096 | 2 | 13,526.5 ±1% | 26,915 ±2% | 1.990× |
| 4096 | 6 | 19,473 ±5% | 60,963.5 ±1% | 3.131× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +132.22% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +256.65% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.50% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +346.15% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +98.98% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +213.07% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 246.55 ±2% | 560 ±2% | 2.271× |
| 64 | 6 | 310.9 ±1% | 1,200.5 ±0% | 3.861× |
| 4096 | 2 | 14,124 ±1% | 26,915 ±2% | 1.906× |
| 4096 | 6 | 18,746 ±1% | 60,963.5 ±1% | 3.252× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +127.13% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +286.14% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +90.56% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +225.21% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 2,031.5 ±0% | 6.5 ±8% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.815 ±1% | 199.2 ±1% | 3.634× |
| 64 | 6 | 74.165 ±1% | 382.2 ±1% | 5.153× |
| 4096 | 2 | 56.285 ±1% | 198.65 ±1% | 3.529× |
| 4096 | 6 | 76.625 ±1% | 383.3 ±1% | 5.002× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +263.40% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +415.34% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +252.94% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +400.23% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.585 ±1% | 199.2 ±1% | 3.520× |
| 64 | 6 | 71.91 ±1% | 382.2 ±1% | 5.315× |
| 4096 | 2 | 57.07 ±1% | 198.65 ±1% | 3.481× |
| 4096 | 6 | 71.925 ±1% | 383.3 ±1% | 5.329× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +252.04% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +431.50% | p=0.000 n=10 |
| 64 | 6 | B/op | +327.56% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +757.14% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +248.08% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +432.92% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 240 ±0% | 4 ±25% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 239 ±0% | 3 ±33% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 239.5 ±0% | 3.5 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L434) | 240 ±0% | 4 ±25% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L349) | 240 ±0% | 4 ±25% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipeline_scenario_test.go#L265) | 1,024 ±0% | 30 ±0% |

## Channel

### Seq 入力

#### 時間：Seq と ToChannel

| N | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L18) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L53) | 倍率 |
| --- | --- | --- | --- |
| 64 | 547.8 ±1% | 6,691 ±1% | 12.214× |
| 4096 | 26,001 ±1% | 338,057 ±1% | 13.002× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1121.43% | p=0.000 n=10 |
| 64 | B/op | +6.53% | p=0.000 n=10 |
| 64 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | sec/op | +1200.17% | p=0.000 n=10 |
| 4096 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Seq](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L18) | 2,328 ±0% | 17 ±0% |
| 64 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L53) | 2,480 ±0% | 16 ±0% |
| 4096 | [Seq](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L18) | 177,688 ±0% | 25 ±0% |
| 4096 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L53) | 177,840 ±0% | 24 ±0% |

### Channel 入力

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L142) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,375 ±1% | 7,674 ±0% | 1.204× |
| 64 | 64 | 2,290 ±1% | 3,367 ±1% | 1.470× |
| 4096 | 0 | 383,680.5 ±1% | 456,316 ±1% | 1.189× |
| 4096 | 64 | 117,171.5 ±2% | 182,843.5 ±0% | 1.560× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +20.38% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +47.03% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +18.93% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +56.05% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L193) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,375 ±1% | 14,140.5 ±1% | 2.218× |
| 64 | 64 | 2,290 ±1% | 8,728.5 ±0% | 3.812× |
| 4096 | 0 | 383,680.5 ±1% | 784,357 ±1% | 2.044× |
| 4096 | 64 | 117,171.5 ±2% | 691,058.5 ±1% | 5.898× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +121.81% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +281.16% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +104.43% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +489.78% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L142) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L193) | 2,752 ±0% | 19 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L142) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L193) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L142) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L193) | 178,112 ±0% | 27 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L89) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L142) | 178,488 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/channel_scenario_test.go#L193) | 178,640 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 390.15 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 18,685 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 1,378 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 79,428 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L30) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 104 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 4,790.5 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 104.25 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 4,788 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 211.3 ±1% | 398 ±1% | 1.884× |
| 64 | 4 | 774.05 ±1% | 452.85 ±1% | 0.585× |
| 64 | 16 | 3,006 ±1% | 654.45 ±0% | 0.218× |
| 4096 | 1 | 11,122.5 ±1% | 19,999.5 ±1% | 1.798× |
| 4096 | 4 | 44,500.5 ±0% | 23,167 ±1% | 0.521× |
| 4096 | 16 | 178,158.5 ±0% | 34,123 ±1% | 0.192× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +88.36% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -41.50% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -78.23% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +79.81% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -47.94% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -80.85% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,245.5 ±1% | 1,385 ±0% | 1.112× |
| 64 | 4 | 4,885 ±1% | 1,435 ±1% | 0.294× |
| 64 | 16 | 19,382.5 ±1% | 1,633.5 ±1% | 0.084× |
| 4096 | 1 | 76,495.5 ±1% | 80,960.5 ±1% | 1.058× |
| 4096 | 4 | 303,994.5 ±1% | 83,743 ±1% | 0.275× |
| 4096 | 16 | 1,213,763.5 ±1% | 95,248.5 ±1% | 0.078× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +11.20% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -70.62% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.57% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +5.84% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -72.45% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -92.15% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 458.75 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 26,271.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 259,609 ±0% | 26 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 1,431 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 85,507.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L378) | 259,608 ±0% | 26 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 119.2 ±8% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 5,533 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 96.265 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 5,532 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 211.4 ±1% | 466.8 ±1% | 2.208× |
| 64 | 4 | 769.25 ±1% | 523.9 ±1% | 0.681× |
| 64 | 16 | 2,994.5 ±2% | 725.4 ±1% | 0.242× |
| 4096 | 1 | 11,025 ±1% | 27,765.5 ±0% | 2.518× |
| 4096 | 4 | 44,117 ±1% | 31,031.5 ±0% | 0.703× |
| 4096 | 16 | 176,832.5 ±1% | 42,779.5 ±0% | 0.242× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +120.81% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -31.89% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -75.78% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +151.84% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190753.68% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -29.66% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.31% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -75.81% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,561 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,257 ±1% | 1,446.5 ±1% | 1.151× |
| 64 | 4 | 4,925 ±1% | 1,497.5 ±1% | 0.304× |
| 64 | 16 | 19,572.5 ±1% | 1,700.5 ±1% | 0.087× |
| 4096 | 1 | 77,187 ±1% | 86,990 ±1% | 1.127× |
| 4096 | 4 | 307,107.5 ±1% | 89,753 ±1% | 0.292× |
| 4096 | 16 | 1,224,639.5 ±1% | 103,306.5 ±1% | 0.084× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +15.08% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -69.59% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.31% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +12.70% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190752.94% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -70.77% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.31% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -91.56% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
