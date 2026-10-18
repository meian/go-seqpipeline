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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 229.7 ±9% | 539.45 ±2% | 2.348× |
| 64 | 6 | 322.4 ±6% | 1,120.5 ±1% | 3.475× |
| 4096 | 2 | 13,546 ±5% | 25,624.5 ±4% | 1.892× |
| 4096 | 6 | 17,973 ±5% | 55,699 ±1% | 3.099× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +134.85% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +247.55% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +89.17% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +209.90% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +155.17% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 235.55 ±6% | 539.45 ±2% | 2.290× |
| 64 | 6 | 305.35 ±4% | 1,120.5 ±1% | 3.670× |
| 4096 | 2 | 14,359.5 ±4% | 25,624.5 ±4% | 1.784× |
| 4096 | 6 | 18,025 ±7% | 55,699 ±1% | 3.090× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +129.02% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +266.96% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +78.45% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +209.01% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 2,031 ±0% | 6 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 177,391 ±0% | 14 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 177,392 ±0% | 14.5 ±3% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.46 ±2% | 194.95 ±1% | 3.580× |
| 64 | 6 | 71.865 ±3% | 376.3 ±2% | 5.236× |
| 4096 | 2 | 55.04 ±2% | 193.9 ±2% | 3.523× |
| 4096 | 6 | 75.485 ±2% | 375.75 ±1% | 4.978× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +257.97% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +423.62% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +252.29% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +397.78% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.975 ±3% | 194.95 ±1% | 3.546× |
| 64 | 6 | 69.575 ±3% | 376.3 ±2% | 5.409× |
| 4096 | 2 | 55.8 ±3% | 193.9 ±2% | 3.475× |
| 4096 | 6 | 71.34 ±2% | 375.75 ±1% | 5.267× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +254.62% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +440.86% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +247.49% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +426.70% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 239 ±0% | 3 ±0% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L543) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 237.85 ±4% | 549.3 ±1% | 2.309× |
| 64 | 6 | 326 ±3% | 1,172.5 ±1% | 3.597× |
| 4096 | 2 | 13,324 ±5% | 26,140.5 ±4% | 1.962× |
| 4096 | 6 | 19,121 ±5% | 59,594.5 ±1% | 3.117× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +130.94% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +259.66% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.50% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +96.19% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +211.67% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 236.35 ±4% | 549.3 ±1% | 2.324× |
| 64 | 6 | 308.05 ±2% | 1,172.5 ±1% | 3.806× |
| 4096 | 2 | 13,733.5 ±4% | 26,140.5 ±4% | 1.903× |
| 4096 | 6 | 18,502 ±5% | 59,594.5 ±1% | 3.221× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +132.41% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +280.62% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.46% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +314.29% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +90.34% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +222.10% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +146.67% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 2,031.5 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 2,032 ±0% | 7 ±14% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 177,392 ±0% | 15 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 54.02 ±2% | 194.9 ±1% | 3.608× |
| 64 | 6 | 72.635 ±2% | 369.95 ±2% | 5.093× |
| 4096 | 2 | 55.005 ±2% | 195.85 ±1% | 3.561× |
| 4096 | 6 | 76.235 ±3% | 371.55 ±1% | 4.874× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +260.79% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.22% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +409.33% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +256.06% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +387.37% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 56.26 ±5% | 194.9 ±1% | 3.464× |
| 64 | 6 | 73.84 ±2% | 369.95 ±2% | 5.010× |
| 4096 | 2 | 56.465 ±2% | 195.85 ±1% | 3.469× |
| 4096 | 6 | 72.96 ±3% | 371.55 ±1% | 5.093× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +246.43% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +401.02% | p=0.000 n=10 |
| 64 | 6 | B/op | +326.67% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +246.85% | p=0.000 n=10 |
| 4096 | 2 | B/op | +166.67% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +350.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +409.25% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 239.5 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 239 ±0% | 3 ±33% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 240 ±0% | 3 ±33% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 239 ±0% | 3 ±0% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 240 ±0% | 4 ±25% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1343) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L1076) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L810) | 1,024 ±0% | 30 ±0% |

## Channel

### 入力条件別

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,248 ±1% | 7,444 ±1% | 1.191× |
| 64 | 64 | 2,272 ±7% | 3,281 ±3% | 1.444× |
| 4096 | 0 | 379,952 ±1% | 447,356.5 ±1% | 1.177× |
| 4096 | 64 | 110,092.5 ±2% | 182,121.5 ±2% | 1.654× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +19.14% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +44.41% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +17.74% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +65.43% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L512) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,248 ±1% | 13,883 ±1% | 2.222× |
| 64 | 64 | 2,272 ±7% | 8,502 ±1% | 3.742× |
| 4096 | 0 | 379,952 ±1% | 768,771 ±1% | 2.023× |
| 4096 | 64 | 110,092.5 ±2% | 661,335.5 ±8% | 6.007× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +122.20% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +274.21% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +102.33% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +500.71% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### 時間：Seq と ToChannel

| N | 入力容量 | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L28) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L126) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 539.7 ±2% | 6,578 ±2% | 12.188× |
| 4096 | 0 | 26,133 ±5% | 331,992.5 ±1% | 12.704× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +1118.83% | p=0.000 n=10 |
| 64 | 0 | B/op | +6.53% | p=0.000 n=10 |
| 64 | 0 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +1170.40% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L512) | 2,752 ±0% | 19 ±0% |
| 64 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L28) | 2,328 ±0% | 17 ±0% |
| 64 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L126) | 2,480 ±0% | 16 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L512) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L512) | 178,112 ±0% | 27 ±0% |
| 4096 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L28) | 177,689 ±0% | 25 ±0% |
| 4096 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L126) | 177,841 ±0% | 24 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 178,488 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L512) | 178,644.5 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 356.45 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 18,285.5 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 1,346 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 79,347.5 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L745) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 92.42 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 4,740.5 ±4% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 40 ±0% | 2 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 91.81 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 4,733.5 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L906) | 40 ±0% | 2 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 211.45 ±1% | 456.35 ±2% | 2.158× |
| 64 | 4 | 773.95 ±1% | 754.45 ±2% | 0.975× |
| 64 | 16 | 3,010 ±1% | 1,891.5 ±3% | 0.628× |
| 4096 | 1 | 11,038 ±1% | 23,230 ±3% | 2.105× |
| 4096 | 4 | 44,065.5 ±1% | 37,694 ±4% | 0.855× |
| 4096 | 16 | 176,308 ±1% | 96,337.5 ±8% | 0.546× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +115.82% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -2.52% | p=0.008 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -37.16% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +110.45% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -14.46% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.24% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -45.36% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.77% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,568 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,952 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 616 ±0% | 36 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,247.5 ±2% | 1,451.5 ±3% | 1.164× |
| 64 | 4 | 4,887.5 ±2% | 1,753 ±2% | 0.359× |
| 64 | 16 | 19,659.5 ±3% | 2,884.5 ±3% | 0.147× |
| 4096 | 1 | 76,637.5 ±2% | 84,758.5 ±3% | 1.106× |
| 4096 | 4 | 305,682.5 ±2% | 100,446 ±3% | 0.329× |
| 4096 | 16 | 1,228,140.5 ±2% | 156,493 ±5% | 0.127× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +16.35% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -64.13% | p=0.000 n=10 |
| 64 | 4 | B/op | +475.86% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.33% | p=0.000 n=10 |
| 64 | 16 | B/op | +179.22% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +10.60% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -67.14% | p=0.000 n=10 |
| 4096 | 4 | B/op | +55317.24% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -87.26% | p=0.000 n=10 |
| 4096 | 16 | B/op | +20833.77% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,336 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 1,720 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,568 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 128,952 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 420.85 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 25,797.5 ±6% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 259,497 ±0% | 23 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 1,373.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 84,844 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1772) | 259,496 ±0% | 23 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 94.84 ±4% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 5,372 ±6% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 94.94 ±2% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 5,366.5 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 40 ±0% | 2 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1933) | 40 ±0% | 2 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 210.4 ±1% | 529.4 ±4% | 2.516× |
| 64 | 4 | 770.45 ±1% | 815.2 ±3% | 1.058× |
| 64 | 16 | 3,009.5 ±1% | 1,965 ±2% | 0.653× |
| 4096 | 1 | 11,254.5 ±0% | 31,615 ±4% | 2.809× |
| 4096 | 4 | 44,834.5 ±0% | 47,006.5 ±3% | 1.048× |
| 4096 | 16 | 178,677.5 ±3% | 111,400.5 ±2% | 0.623× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +151.62% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | +5.81% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -34.71% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +180.91% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190736.03% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | +4.84% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.78% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -37.65% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.55% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 259,537 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 259,633 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 260,017 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 616 ±0% | 36 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,259.5 ±3% | 1,494 ±3% | 1.186× |
| 64 | 4 | 4,938 ±2% | 1,795 ±3% | 0.364× |
| 64 | 16 | 19,718 ±3% | 2,927 ±2% | 0.148× |
| 4096 | 1 | 77,566 ±3% | 90,860 ±2% | 1.171× |
| 4096 | 4 | 310,084.5 ±3% | 107,966.5 ±3% | 0.348× |
| 4096 | 16 | 1,238,235 ±3% | 171,877 ±2% | 0.139× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +18.62% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.65% | p=0.000 n=10 |
| 64 | 4 | B/op | +913.79% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +91.67% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.16% | p=0.000 n=10 |
| 64 | 16 | B/op | +344.16% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +30.56% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +17.14% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190735.29% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +333.33% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -65.18% | p=0.000 n=10 |
| 4096 | 4 | B/op | +111810.34% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +166.67% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -86.12% | p=0.000 n=10 |
| 4096 | 16 | B/op | +42110.39% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +55.56% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 136 ±0% | 6 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,352 ±0% | 23 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 232 ±0% | 12 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 2,736 ±0% | 47 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 616 ±0% | 36 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 259,536 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 136 ±0% | 6 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 259,632 ±0% | 32 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 232 ±0% | 12 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 260,016 ±0% | 56 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 616 ±0% | 36 ±0% |
