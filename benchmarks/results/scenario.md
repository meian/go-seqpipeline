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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 220.75 ±7% | 531.1 ±2% | 2.406× |
| 64 | 6 | 319.4 ±2% | 1,106.5 ±2% | 3.464× |
| 4096 | 2 | 13,238 ±7% | 25,286 ±2% | 1.910× |
| 4096 | 6 | 17,954 ±4% | 54,992 ±2% | 3.063× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +140.59% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +246.43% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +91.01% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +206.29% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 236.75 ±4% | 531.1 ±2% | 2.243× |
| 64 | 6 | 296.55 ±3% | 1,106.5 ±2% | 3.731× |
| 4096 | 2 | 13,621 ±6% | 25,286 ±2% | 1.856× |
| 4096 | 6 | 17,272 ±9% | 54,992 ±2% | 3.184× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +124.33% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +273.12% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +85.64% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +218.39% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 52.875 ±3% | 191.4 ±2% | 3.620× |
| 64 | 6 | 71.6 ±1% | 373.4 ±2% | 5.215× |
| 4096 | 2 | 56.195 ±2% | 193.35 ±2% | 3.441× |
| 4096 | 6 | 75.56 ±2% | 374 ±1% | 4.950× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +261.99% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +421.51% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +244.07% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +394.97% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 55.455 ±1% | 191.4 ±2% | 3.451× |
| 64 | 6 | 70.24 ±1% | 373.4 ±2% | 5.316× |
| 4096 | 2 | 55.11 ±3% | 193.35 ±2% | 3.508× |
| 4096 | 6 | 70.715 ±2% | 374 ±1% | 5.289× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +245.14% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +431.61% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +250.84% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +428.88% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L179) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L94) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L10) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 225.3 ±6% | 540.95 ±1% | 2.401× |
| 64 | 6 | 326.7 ±2% | 1,157.5 ±2% | 3.543× |
| 4096 | 2 | 12,952 ±2% | 25,452 ±1% | 1.965× |
| 4096 | 6 | 17,817.5 ±5% | 59,216 ±2% | 3.323× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +140.10% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.60% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +161.54% | p=0.000 n=10 |
| 64 | 6 | sec/op | +254.30% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +96.51% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +232.35% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 229.15 ±5% | 540.95 ±1% | 2.361× |
| 64 | 6 | 300.25 ±2% | 1,157.5 ±2% | 3.855× |
| 4096 | 2 | 13,259 ±5% | 25,452 ±1% | 1.920× |
| 4096 | 6 | 17,195 ±4% | 59,216 ±2% | 3.444× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +136.07% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +285.51% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.50% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +346.15% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +91.96% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +244.38% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +155.17% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 2,031.5 ±0% | 6.5 ±8% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 2,032 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 2,031.5 ±0% | 6.5 ±8% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 177,392 ±0% | 14.5 ±3% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 52.83 ±3% | 194.35 ±1% | 3.679× |
| 64 | 6 | 71.965 ±2% | 370.65 ±2% | 5.150× |
| 4096 | 2 | 54.905 ±2% | 196.2 ±2% | 3.573× |
| 4096 | 6 | 75.84 ±2% | 371.15 ±2% | 4.894× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +267.88% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +415.04% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +257.34% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.22% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +389.39% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 55.475 ±3% | 194.35 ±1% | 3.503× |
| 64 | 6 | 71.09 ±3% | 370.65 ±2% | 5.214× |
| 4096 | 2 | 55.85 ±2% | 196.2 ±2% | 3.513× |
| 4096 | 6 | 70.72 ±2% | 371.15 ±2% | 5.248× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +250.34% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +421.38% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +251.30% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +424.82% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 240 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 240 ±0% | 3 ±33% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 239.5 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L433) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L348) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/pipeline_scenario_test.go#L264) | 1,024 ±0% | 30 ±0% |

## Channel

### Seq 入力

#### 時間：Seq と ToChannel

| N | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L17) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L52) | 倍率 |
| --- | --- | --- | --- |
| 64 | 539.3 ±3% | 6,604 ±1% | 12.246× |
| 4096 | 25,397 ±3% | 337,583 ±1% | 13.292× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1124.55% | p=0.000 n=10 |
| 64 | B/op | +6.53% | p=0.000 n=10 |
| 64 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | sec/op | +1229.22% | p=0.000 n=10 |
| 4096 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Seq](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L17) | 2,328 ±0% | 17 ±0% |
| 64 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L52) | 2,480 ±0% | 16 ±0% |
| 4096 | [Seq](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L17) | 177,688 ±0% | 25 ±0% |
| 4096 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L52) | 177,840 ±0% | 24 ±0% |

### Channel 入力

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L141) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,190 ±1% | 7,712 ±2% | 1.246× |
| 64 | 64 | 2,244 ±6% | 3,269.5 ±3% | 1.457× |
| 4096 | 0 | 370,741 ±1% | 462,758 ±2% | 1.248× |
| 4096 | 64 | 110,921 ±5% | 182,578 ±2% | 1.646× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +24.59% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +45.70% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +24.82% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +64.60% | p=0.001 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L192) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,190 ±1% | 14,098.5 ±1% | 2.278× |
| 64 | 64 | 2,244 ±6% | 8,596.5 ±2% | 3.831× |
| 4096 | 0 | 370,741 ±1% | 782,169 ±1% | 2.110× |
| 4096 | 64 | 110,921 ±5% | 654,353.5 ±4% | 5.899× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +127.76% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +283.09% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +110.97% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +489.93% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L141) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L192) | 2,752 ±0% | 19 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L141) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L192) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L141) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L192) | 178,112 ±0% | 27 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L88) | 178,144.5 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L141) | 178,488 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/channel_scenario_test.go#L192) | 178,640 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 455.7 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 17,804 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 1,321 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 78,913 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L29) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 92.785 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 4,735 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 92.775 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 4,768.5 ±8% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L78) | 40 ±0% | 2 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 208.45 ±1% | 452.9 ±2% | 2.173× |
| 64 | 4 | 768.8 ±1% | 742.7 ±2% | 0.966× |
| 64 | 16 | 2,998 ±0% | 1,877 ±2% | 0.626× |
| 4096 | 1 | 11,340.5 ±1% | 22,850 ±3% | 2.015× |
| 4096 | 4 | 45,204.5 ±0% | 36,890 ±3% | 0.816× |
| 4096 | 16 | 180,679.5 ±0% | 93,985 ±5% | 0.520× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +117.27% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -3.39% | p=0.000 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -37.39% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +101.49% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -18.39% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.67% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -47.98% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.93% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,569 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,953 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,241.5 ±2% | 1,451 ±3% | 1.169× |
| 64 | 4 | 4,886.5 ±1% | 1,749.5 ±3% | 0.358× |
| 64 | 16 | 19,488.5 ±3% | 2,849 ±2% | 0.146× |
| 4096 | 1 | 75,823 ±1% | 84,299 ±2% | 1.112× |
| 4096 | 4 | 305,265.5 ±2% | 99,985 ±3% | 0.328× |
| 4096 | 16 | 1,220,563.5 ±4% | 155,948 ±4% | 0.128× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +16.87% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -64.20% | p=0.000 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.38% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +11.18% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -67.25% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.24% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -87.22% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.77% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,568 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L131) | 128,952 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L250) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 415.45 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 24,726 ±6% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 259,497 ±0% | 23 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 1,373.5 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 83,568 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L371) | 259,496 ±0% | 23 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 94.475 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 5,362 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 94.475 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 5,354.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L423) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 209.9 ±1% | 522 ±2% | 2.487× |
| 64 | 4 | 769 ±1% | 806.3 ±2% | 1.049× |
| 64 | 16 | 3,001 ±0% | 1,947.5 ±2% | 0.649× |
| 4096 | 1 | 10,975.5 ±0% | 30,525.5 ±3% | 2.781× |
| 4096 | 4 | 43,789.5 ±1% | 45,964 ±1% | 1.050× |
| 4096 | 16 | 174,936 ±3% | 110,221.5 ±2% | 0.630× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +148.69% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | +4.85% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -35.10% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +178.12% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190736.03% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | +4.97% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.78% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -36.99% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.55% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 259,537 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 259,633 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 260,017 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,257 ±1% | 1,478.5 ±1% | 1.176× |
| 64 | 4 | 4,938.5 ±3% | 1,786.5 ±2% | 0.362× |
| 64 | 16 | 19,729 ±3% | 2,921 ±2% | 0.148× |
| 4096 | 1 | 77,500.5 ±2% | 89,967 ±2% | 1.161× |
| 4096 | 4 | 308,741.5 ±2% | 106,871 ±2% | 0.346× |
| 4096 | 16 | 1,223,020 ±2% | 169,878 ±2% | 0.139× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +17.62% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.83% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.19% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +16.09% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190735.29% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -65.38% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.34% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -86.11% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.39% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 259,536 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 259,632 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L476) | 260,016 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/6dd1dd4f0afebc6d4f562883f4ef9a6a457ecde2/benchmarks/cache_scenario_test.go#L595) | 616 ±0% | 36 ±0% |
