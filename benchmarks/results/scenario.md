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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 230.05 ±3% | 537.65 ±2% | 2.337× |
| 64 | 6 | 321.8 ±2% | 1,120 ±1% | 3.480× |
| 4096 | 2 | 12,296.5 ±6% | 25,019.5 ±1% | 2.035× |
| 4096 | 6 | 18,108 ±5% | 55,190 ±1% | 3.048× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +133.71% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +248.04% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +103.47% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +204.78% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 233.95 ±3% | 537.65 ±2% | 2.298× |
| 64 | 6 | 295.25 ±2% | 1,120 ±1% | 3.793× |
| 4096 | 2 | 13,426 ±7% | 25,019.5 ±1% | 1.864× |
| 4096 | 6 | 17,320.5 ±7% | 55,190 ±1% | 3.186× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +129.81% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +279.34% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +86.35% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +218.64% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +155.17% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 177,391.5 ±0% | 14.5 ±3% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 53.86 ±2% | 194.65 ±1% | 3.614× |
| 64 | 6 | 73.065 ±1% | 376.35 ±1% | 5.151× |
| 4096 | 2 | 55.22 ±1% | 195.25 ±2% | 3.536× |
| 4096 | 6 | 76.21 ±2% | 379.2 ±1% | 4.976× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +261.40% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +415.09% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +253.59% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +397.57% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.195 ±1% | 194.65 ±1% | 3.464× |
| 64 | 6 | 71.53 ±2% | 376.35 ±1% | 5.261× |
| 4096 | 2 | 56.325 ±2% | 195.25 ±2% | 3.466× |
| 4096 | 6 | 84.785 ±1% | 379.2 ±1% | 4.472× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +246.38% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +426.14% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +246.65% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.22% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +347.25% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 239.5 ±0% | 3 ±33% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L180) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L95) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L11) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 232.95 ±3% | 551.3 ±1% | 2.367× |
| 64 | 6 | 326.75 ±2% | 1,179.5 ±1% | 3.610× |
| 4096 | 2 | 13,254 ±6% | 26,106.5 ±1% | 1.970× |
| 4096 | 6 | 18,542 ±7% | 59,696 ±0% | 3.220× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +136.66% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.57% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +142.86% | p=0.000 n=10 |
| 64 | 6 | sec/op | +260.98% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.50% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +96.97% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +221.95% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 239.15 ±3% | 551.3 ±1% | 2.305× |
| 64 | 6 | 299.65 ±3% | 1,179.5 ±1% | 3.936× |
| 4096 | 2 | 13,094.5 ±6% | 26,106.5 ±1% | 1.994× |
| 4096 | 6 | 17,704 ±4% | 59,696 ±0% | 3.372× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +130.52% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.60% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +293.63% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +99.37% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +237.19% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 2,032 ±0% | 7 ±14% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 2,031.5 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 2,031.5 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.05 ±1% | 196.85 ±2% | 3.642× |
| 64 | 6 | 73.5 ±1% | 375.55 ±1% | 5.110× |
| 4096 | 2 | 55.955 ±2% | 197.35 ±1% | 3.527× |
| 4096 | 6 | 77.335 ±2% | 375.9 ±1% | 4.861× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +264.20% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +410.95% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +252.69% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +386.07% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.385 ±2% | 196.85 ±2% | 3.491× |
| 64 | 6 | 72.12 ±2% | 375.55 ±1% | 5.207× |
| 4096 | 2 | 56.785 ±1% | 197.35 ±1% | 3.475× |
| 4096 | 6 | 71.98 ±1% | 375.9 ±1% | 5.222× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +249.12% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +420.73% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +247.54% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.22% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +414.29% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +422.23% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 239 ±0% | 3 ±33% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 240 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 239.5 ±0% | 3.5 ±14% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L434) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L349) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/pipeline_scenario_test.go#L265) | 1,024 ±0% | 30 ±0% |

## Channel

### Seq 入力

#### 時間：Seq と ToChannel

| N | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L18) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L53) | 倍率 |
| --- | --- | --- | --- |
| 64 | 542.7 ±1% | 6,633 ±1% | 12.222× |
| 4096 | 25,277.5 ±3% | 334,130 ±1% | 13.218× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1122.22% | p=0.000 n=10 |
| 64 | B/op | +6.53% | p=0.000 n=10 |
| 64 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | sec/op | +1221.85% | p=0.000 n=10 |
| 4096 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Seq](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L18) | 2,328 ±0% | 17 ±0% |
| 64 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L53) | 2,480 ±0% | 16 ±0% |
| 4096 | [Seq](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L18) | 177,688 ±0% | 25 ±0% |
| 4096 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L53) | 177,840 ±0% | 24 ±0% |

### Channel 入力

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L142) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,331 ±1% | 7,608 ±0% | 1.202× |
| 64 | 64 | 2,242.5 ±1% | 3,334.5 ±2% | 1.487× |
| 4096 | 0 | 386,532.5 ±2% | 456,508 ±1% | 1.181× |
| 4096 | 64 | 110,736.5 ±3% | 179,100.5 ±1% | 1.617× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +20.17% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +48.70% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +18.10% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +61.74% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L193) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,331 ±1% | 14,005 ±0% | 2.212× |
| 64 | 64 | 2,242.5 ±1% | 8,584.5 ±0% | 3.828× |
| 4096 | 0 | 386,532.5 ±2% | 779,104 ±1% | 2.016× |
| 4096 | 64 | 110,736.5 ±3% | 660,512 ±1% | 5.965× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +121.21% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +282.81% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +101.56% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +496.47% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L142) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L193) | 2,752 ±0% | 19 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L142) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L193) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L142) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L193) | 178,112 ±0% | 27 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L89) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L142) | 178,488 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/channel_scenario_test.go#L193) | 178,649 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 411.5 ±5% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 18,214 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 1,354.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 78,318 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 1,312 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L30) | 128,544 ±0% | 26 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 105.05 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 5,117 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 105.05 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 5,113.5 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L85) | 88 ±0% | 3 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 209.6 ±1% | 394.7 ±2% | 1.883× |
| 64 | 4 | 771.15 ±0% | 449.05 ±2% | 0.582× |
| 64 | 16 | 3,011 ±0% | 646.05 ±1% | 0.215× |
| 4096 | 1 | 11,133.5 ±3% | 19,394.5 ±1% | 1.742× |
| 4096 | 4 | 44,380 ±2% | 22,694 ±1% | 0.511× |
| 4096 | 16 | 177,214.5 ±3% | 33,662 ±1% | 0.190× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +88.31% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -41.77% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -78.54% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +74.20% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -48.86% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -81.00% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,234.5 ±1% | 1,356 ±1% | 1.098× |
| 64 | 4 | 4,850 ±1% | 1,410 ±1% | 0.291× |
| 64 | 16 | 19,232 ±2% | 1,609.5 ±1% | 0.084× |
| 4096 | 1 | 75,580 ±1% | 79,652 ±1% | 1.054× |
| 4096 | 4 | 299,655.5 ±1% | 81,996 ±1% | 0.274× |
| 4096 | 16 | 1,199,316 ±1% | 93,600.5 ±1% | 0.078× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +9.84% | p=0.000 n=10 |
| 64 | 1 | B/op | +829.41% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -70.93% | p=0.000 n=10 |
| 64 | 4 | B/op | +444.83% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.63% | p=0.000 n=10 |
| 64 | 16 | B/op | +105.19% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +5.39% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94382.35% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -72.64% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55286.21% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -92.20% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20759.74% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 1,264 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L138) | 128,496 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L257) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 452.5 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 24,866.5 ±4% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 259,609 ±0% | 26 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 1,406.5 ±1% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 83,481 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 2,328 ±0% | 17 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L378) | 259,608 ±0% | 26 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 99.63 ±4% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 5,724 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 100.33 ±4% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 5,751.5 ±4% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L435) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 210.3 ±0% | 463.95 ±2% | 2.206× |
| 64 | 4 | 770.6 ±0% | 517.05 ±1% | 0.671× |
| 64 | 16 | 3,002.5 ±1% | 716.15 ±1% | 0.239× |
| 4096 | 1 | 11,022.5 ±0% | 26,904.5 ±4% | 2.441× |
| 4096 | 4 | 43,845 ±0% | 30,587 ±2% | 0.698× |
| 4096 | 16 | 175,062 ±4% | 42,661 ±1% | 0.244× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +120.61% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -32.90% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -76.15% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +144.09% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190753.68% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -30.24% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.31% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -75.63% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,561 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,245.5 ±1% | 1,415.5 ±1% | 1.136× |
| 64 | 4 | 4,877 ±1% | 1,469 ±2% | 0.301× |
| 64 | 16 | 19,395.5 ±2% | 1,674 ±1% | 0.086× |
| 4096 | 1 | 76,394 ±1% | 85,663 ±1% | 1.121× |
| 4096 | 4 | 304,602.5 ±2% | 88,103 ±1% | 0.289× |
| 4096 | 16 | 1,207,812.5 ±1% | 101,540 ±1% | 0.084× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +13.65% | p=0.000 n=10 |
| 64 | 1 | B/op | +1576.47% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -69.88% | p=0.000 n=10 |
| 64 | 4 | B/op | +882.76% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +25.00% | p=0.000 n=10 |
| 64 | 16 | sec/op | -91.37% | p=0.000 n=10 |
| 64 | 16 | B/op | +270.13% | p=0.000 n=10 |
| 64 | 16 | allocs/op | -58.33% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +12.13% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190752.94% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +300.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -71.08% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111779.31% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -91.59% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42036.36% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | -33.33% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 2,280 ±0% | 15 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L488) | 259,560 ±0% | 24 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/45e5daaf139c8b0cc89cec101ae4602bacfbf0f5/benchmarks/cache_scenario_test.go#L607) | 616 ±0% | 36 ±0% |
