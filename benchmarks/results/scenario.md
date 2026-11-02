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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 233.8 ±2% | 538.65 ±3% | 2.304× |
| 64 | 6 | 320.3 ±1% | 1,133.5 ±1% | 3.539× |
| 4096 | 2 | 13,006.5 ±4% | 25,419 ±2% | 1.954× |
| 4096 | 6 | 19,237.5 ±5% | 56,289 ±1% | 2.926× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +130.39% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +253.89% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +95.43% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +192.60% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 237.65 ±1% | 538.65 ±3% | 2.267× |
| 64 | 6 | 296.65 ±1% | 1,133.5 ±1% | 3.821× |
| 4096 | 2 | 13,502.5 ±8% | 25,419 ±2% | 1.883× |
| 4096 | 6 | 17,712 ±5% | 56,289 ±1% | 3.178× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +126.66% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +282.10% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +88.25% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +217.80% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 2,032 ±0% | 7 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 55.13 ±1% | 198.3 ±4% | 3.597× |
| 64 | 6 | 77.16 ±2% | 384.1 ±2% | 4.978× |
| 4096 | 2 | 56.145 ±5% | 199.55 ±3% | 3.554× |
| 4096 | 6 | 78.55 ±2% | 387.2 ±1% | 4.929× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +259.70% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +397.80% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +255.42% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +392.93% | p=0.000 n=10 |
| 4096 | 6 | B/op | +327.56% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +757.14% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.47 ±3% | 198.3 ±4% | 3.512× |
| 64 | 6 | 71.53 ±2% | 384.1 ±2% | 5.370× |
| 4096 | 2 | 57.725 ±2% | 199.55 ±3% | 3.457× |
| 4096 | 6 | 72.685 ±1% | 387.2 ±1% | 5.327× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +251.16% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +436.98% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +245.69% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +432.71% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 240 ±0% | 4 ±25% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 240 ±0% | 4 ±25% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 240 ±0% | 4 ±25% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 240 ±0% | 4 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L45) | 239.5 ±0% | 3.5 ±14% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 240 ±0% | 4 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 233.75 ±2% | 556.85 ±3% | 2.382× |
| 64 | 6 | 323.85 ±2% | 1,158 ±4% | 3.576× |
| 4096 | 2 | 13,128 ±4% | 25,688.5 ±2% | 1.957× |
| 4096 | 6 | 18,400 ±5% | 58,120.5 ±4% | 3.159× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +138.22% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +257.57% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +95.68% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +215.87% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 235.75 ±3% | 556.85 ±3% | 2.362× |
| 64 | 6 | 303.35 ±2% | 1,158 ±4% | 3.817× |
| 4096 | 2 | 13,384.5 ±3% | 25,688.5 ±2% | 1.919× |
| 4096 | 6 | 17,760 ±4% | 58,120.5 ±4% | 3.273× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +136.20% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +281.74% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +91.93% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +227.26% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 2,032 ±0% | 7 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 177,392 ±0% | 15 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 177,392 ±0% | 15 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 55.625 ±1% | 198.2 ±5% | 3.563× |
| 64 | 6 | 76.9 ±1% | 381.4 ±2% | 4.960× |
| 4096 | 2 | 55.71 ±2% | 201.3 ±3% | 3.613× |
| 4096 | 6 | 77.605 ±2% | 381.95 ±3% | 4.922× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +256.31% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +395.97% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +261.34% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +392.17% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.97 ±3% | 198.2 ±5% | 3.479× |
| 64 | 6 | 73.485 ±3% | 381.4 ±2% | 5.190× |
| 4096 | 2 | 57.38 ±1% | 201.3 ±3% | 3.508× |
| 4096 | 6 | 74.135 ±3% | 381.95 ±3% | 5.152× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +247.90% | p=0.000 n=10 |
| 64 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +419.02% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +650.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +250.82% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +415.21% | p=0.000 n=10 |
| 4096 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +650.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 240 ±0% | 4 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 240 ±0% | 4 ±25% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 240 ±0% | 4 ±25% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 240 ±0% | 4 ±25% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 240 ±0% | 4 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 240 ±0% | 4 ±0% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L168) | 240 ±0% | 4 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L151) | 240 ±0% | 4 ±25% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L135) | 1,024 ±0% | 30 ±0% |

## Channel

### Seq 入力

#### 時間：Seq と ToChannel

| N | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L52) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L19) | 倍率 |
| --- | --- | --- | --- |
| 64 | 533.55 ±2% | 4,976 ±0% | 9.326× |
| 4096 | 25,203.5 ±3% | 292,317 ±1% | 11.598× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +832.62% | p=0.000 n=10 |
| 64 | B/op | +4.81% | p=0.000 n=10 |
| 64 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | sec/op | +1059.83% | p=0.000 n=10 |
| 4096 | B/op | +0.06% | p=0.000 n=10 |
| 4096 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Seq](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L52) | 2,328 ±0% | 17 ±0% |
| 64 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L19) | 2,440 ±0% | 16 ±0% |
| 4096 | [Seq](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L52) | 177,688 ±0% | 25 ±0% |
| 4096 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L19) | 177,801.5 ±0% | 24 ±0% |

### Channel 入力

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,384.5 ±2% | 6,787 ±1% | 1.063× |
| 64 | 64 | 2,264.5 ±3% | 2,659.5 ±1% | 1.174× |
| 4096 | 0 | 386,919 ±2% | 402,423.5 ±1% | 1.040× |
| 4096 | 64 | 112,686.5 ±4% | 125,503 ±1% | 1.114× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +6.30% | p=0.000 n=10 |
| 64 | 0 | B/op | +14.18% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +17.44% | p=0.000 n=10 |
| 64 | 64 | B/op | +11.49% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +4.01% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.18% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +11.37% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.18% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L40) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,384.5 ±2% | 11,316 ±0% | 1.772× |
| 64 | 64 | 2,264.5 ±3% | 6,444.5 ±1% | 2.846× |
| 4096 | 0 | 386,919 ±2% | 683,467.5 ±1% | 1.766× |
| 4096 | 64 | 112,686.5 ±4% | 561,893.5 ±5% | 4.986× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +77.24% | p=0.000 n=10 |
| 64 | 0 | B/op | +19.15% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +184.59% | p=0.000 n=10 |
| 64 | 64 | B/op | +15.52% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +76.64% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.24% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +398.63% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.24% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 2,576 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L40) | 2,688 ±0% | 19 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 3,104 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L40) | 3,216 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 177,936 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L40) | 178,048 ±0% | 27 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 178,464 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L40) | 178,576 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 398.45 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 18,213 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 1,348.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 77,550.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L31) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 107.3 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 5,376 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 88 ±0% | 3 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 107.4 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 5,374 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L75) | 88 ±0% | 3 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 210.5 ±1% | 394.25 ±2% | 1.873× |
| 64 | 4 | 772.7 ±1% | 447.35 ±2% | 0.579× |
| 64 | 16 | 3,003 ±1% | 645.25 ±2% | 0.215× |
| 4096 | 1 | 11,363.5 ±1% | 19,530 ±3% | 1.719× |
| 4096 | 4 | 45,330 ±1% | 22,673.5 ±4% | 0.500× |
| 4096 | 16 | 181,122.5 ±1% | 33,418 ±1% | 0.185× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +87.29% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -42.11% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -78.51% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +71.87% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -49.98% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -81.55% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,223 ±2% | 1,354.5 ±2% | 1.108× |
| 64 | 4 | 4,812.5 ±1% | 1,406.5 ±2% | 0.292× |
| 64 | 16 | 19,106.5 ±1% | 1,602.5 ±2% | 0.084× |
| 4096 | 1 | 75,425 ±2% | 78,970.5 ±2% | 1.047× |
| 4096 | 4 | 298,786.5 ±1% | 81,831 ±2% | 0.274× |
| 4096 | 16 | 1,195,231.5 ±2% | 93,096.5 ±2% | 0.078× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +10.75% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -70.77% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.61% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +4.70% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -72.61% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -92.21% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 451.4 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 25,400.5 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 259,609 ±0% | 26 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 1,397.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 83,233 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L207) | 259,608 ±0% | 26 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 97.215 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 5,526 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 97.205 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 5,531.5 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L253) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 209.7 ±1% | 458.1 ±2% | 2.185× |
| 64 | 4 | 767.3 ±1% | 512.6 ±3% | 0.668× |
| 64 | 16 | 2,991.5 ±1% | 711.2 ±2% | 0.238× |
| 4096 | 1 | 11,277 ±2% | 26,926.5 ±3% | 2.388× |
| 4096 | 4 | 44,874 ±2% | 30,370.5 ±1% | 0.677× |
| 4096 | 16 | 179,100.5 ±2% | 41,464.5 ±2% | 0.232× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +118.45% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -33.19% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -76.23% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +138.77% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190753.68% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -32.32% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.53% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -76.85% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,561 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,560.5 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,237 ±2% | 1,408 ±2% | 1.138× |
| 64 | 4 | 4,874.5 ±2% | 1,460 ±2% | 0.300× |
| 64 | 16 | 19,308 ±1% | 1,664.5 ±2% | 0.086× |
| 4096 | 1 | 76,005 ±2% | 84,705 ±2% | 1.114× |
| 4096 | 4 | 303,405.5 ±1% | 87,972.5 ±2% | 0.290× |
| 4096 | 16 | 1,205,696.5 ±2% | 101,250 ±2% | 0.084× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +13.82% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -70.05% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.38% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +11.45% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190752.94% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -71.00% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.31% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -91.60% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,560 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,560 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 616 ±0% | 36 ±0% |
