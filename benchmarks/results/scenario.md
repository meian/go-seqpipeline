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

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 248.75 ±4% | 546.05 ±4% | 2.195× |
| 64 | 6 | 338.85 ±3% | 1,137.5 ±2% | 3.357× |
| 4096 | 2 | 13,801.5 ±2% | 26,466.5 ±1% | 1.918× |
| 4096 | 6 | 20,406.5 ±6% | 57,750 ±3% | 2.830× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +119.52% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +235.69% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +91.77% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +183.00% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +155.17% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 243.25 ±5% | 546.05 ±4% | 2.245× |
| 64 | 6 | 308.55 ±2% | 1,137.5 ±2% | 3.687× |
| 4096 | 2 | 13,779 ±7% | 26,466.5 ±1% | 1.921× |
| 4096 | 6 | 17,838 ±6% | 57,750 ±3% | 3.237× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +124.48% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +268.66% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +92.08% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +223.75% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +155.17% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 177,391.5 ±0% | 14.5 ±3% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 177,391.5 ±0% | 14.5 ±3% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 178,072 ±0% | 37 ±0% |

### Seq / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 57.655 ±2% | 197.8 ±3% | 3.431× |
| 64 | 6 | 80.38 ±2% | 382.85 ±3% | 4.763× |
| 4096 | 2 | 57.195 ±1% | 198.15 ±6% | 3.464× |
| 4096 | 6 | 81.135 ±2% | 383.6 ±2% | 4.728× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +243.08% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +376.30% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +246.45% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +372.79% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 58.935 ±2% | 197.8 ±3% | 3.356× |
| 64 | 6 | 75.445 ±3% | 382.85 ±3% | 5.075× |
| 4096 | 2 | 59.845 ±4% | 198.15 ±6% | 3.311× |
| 4096 | 6 | 75.51 ±2% | 383.6 ±2% | 5.080× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +235.62% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +407.46% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +231.11% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +408.01% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±0% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 1,024 ±0% | 30 ±0% |

### Seq2 / 全件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 250.7 ±6% | 556 ±5% | 2.218× |
| 64 | 6 | 341.4 ±3% | 1,139 ±2% | 3.336× |
| 4096 | 2 | 14,376.5 ±5% | 26,798 ±1% | 1.864× |
| 4096 | 6 | 20,709 ±6% | 60,239 ±1% | 2.909× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +121.78% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +233.63% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +86.40% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +66.67% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +190.88% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 243.65 ±2% | 556 ±5% | 2.282× |
| 64 | 6 | 317 ±2% | 1,139 ±2% | 3.593× |
| 4096 | 2 | 14,443 ±7% | 26,798 ±1% | 1.855× |
| 4096 | 6 | 18,238 ±6% | 60,239 ±1% | 3.303× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +128.20% | p=0.000 n=10 |
| 64 | 2 | B/op | +14.62% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +183.33% | p=0.000 n=10 |
| 64 | 6 | sec/op | +259.31% | p=0.000 n=10 |
| 64 | 6 | B/op | +33.53% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +383.33% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +85.54% | p=0.000 n=10 |
| 4096 | 2 | B/op | +0.17% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +78.57% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +230.29% | p=0.000 n=10 |
| 4096 | 6 | B/op | +0.38% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +164.29% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 2,031 ±0% | 6 ±17% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 2,328 ±0% | 17 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 2,031 ±0% | 6 ±17% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 2,712 ±0% | 29 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 177,392 ±0% | 15 ±7% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 177,391.5 ±0% | 14 ±7% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 177,688 ±0% | 25 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 177,391 ±0% | 14 ±7% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 178,072 ±0% | 37 ±0% |

### Seq2 / 先頭8件

#### 時間：Functions と Pipeline

| N | 段数 | [Functions ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 58.29 ±1% | 199.05 ±3% | 3.415× |
| 64 | 6 | 78.445 ±3% | 377.4 ±4% | 4.811× |
| 4096 | 2 | 59.095 ±6% | 198.95 ±4% | 3.367× |
| 4096 | 6 | 79.455 ±2% | 380.4 ±2% | 4.788× |

<details>
<summary>検定：Functions と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +241.48% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +381.10% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +236.66% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +378.76% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### 時間：Loop と Pipeline

| N | 段数 | [Loop ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | [Pipeline ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 2 | 59.405 ±2% | 199.05 ±3% | 3.351× |
| 64 | 6 | 74.59 ±3% | 377.4 ±4% | 5.060× |
| 4096 | 2 | 60.6 ±6% | 198.95 ±4% | 3.283× |
| 4096 | 6 | 75.525 ±2% | 380.4 ±2% | 5.037× |

<details>
<summary>検定：Loop と Pipeline</summary>

| N | 段数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 2 | sec/op | +235.07% | p=0.000 n=10 |
| 64 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 64 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 64 | 6 | sec/op | +405.97% | p=0.000 n=10 |
| 64 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 64 | 6 | allocs/op | +900.00% | p=0.000 n=10 |
| 4096 | 2 | sec/op | +228.30% | p=0.000 n=10 |
| 4096 | 2 | B/op | +167.78% | p=0.000 n=10 |
| 4096 | 2 | allocs/op | +500.00% | p=0.000 n=10 |
| 4096 | 6 | sec/op | +403.67% | p=0.000 n=10 |
| 4096 | 6 | B/op | +328.45% | p=0.000 n=10 |
| 4096 | 6 | allocs/op | +900.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 段数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 64 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 640 ±0% | 18 ±0% |
| 64 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±0% |
| 64 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 1,024 ±0% | 30 ±0% |
| 4096 | 2 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 4096 | 2 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 640 ±0% | 18 ±0% |
| 4096 | 6 | [Functions](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L63) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 239 ±0% | 3 ±33% |
| 4096 | 6 | [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 1,024 ±0% | 30 ±0% |

## Channel

### 入力条件別

#### 時間：Direct と FromChannel

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | [FromChannel ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,335 ±1% | 7,530 ±1% | 1.189× |
| 64 | 64 | 2,251.5 ±3% | 3,288 ±3% | 1.460× |
| 4096 | 0 | 389,577.5 ±2% | 454,514.5 ±1% | 1.167× |
| 4096 | 64 | 109,611 ±3% | 182,735.5 ±1% | 1.667× |

<details>
<summary>検定：Direct と FromChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +18.86% | p=0.000 n=10 |
| 64 | 0 | B/op | +15.25% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +150.00% | p=0.000 n=10 |
| 64 | 64 | sec/op | +46.04% | p=0.000 n=10 |
| 64 | 64 | B/op | +12.36% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +16.67% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +75.00% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +66.71% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.19% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +75.00% | p=0.000 n=10 |

</details>

#### 時間：Direct と RoundTrip

| N | 入力容量 | [Direct ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | [RoundTrip ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L86) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 6,335 ±1% | 13,881 ±1% | 2.191× |
| 64 | 64 | 2,251.5 ±3% | 8,477 ±1% | 3.765× |
| 4096 | 0 | 389,577.5 ±2% | 776,972.5 ±1% | 1.994× |
| 4096 | 64 | 109,611 ±3% | 663,434.5 ±5% | 6.053× |

<details>
<summary>検定：Direct と RoundTrip</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +119.12% | p=0.000 n=10 |
| 64 | 0 | B/op | +21.99% | p=0.000 n=10 |
| 64 | 0 | allocs/op | +137.50% | p=0.000 n=10 |
| 64 | 64 | sec/op | +276.50% | p=0.000 n=10 |
| 64 | 64 | B/op | +17.82% | p=0.000 n=10 |
| 64 | 64 | allocs/op | +137.50% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +99.44% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | +68.75% | p=0.000 n=10 |
| 4096 | 64 | sec/op | +505.26% | p=0.000 n=10 |
| 4096 | 64 | B/op | +0.28% | p=0.000 n=10 |
| 4096 | 64 | allocs/op | +68.75% | p=0.000 n=10 |

</details>

#### 時間：Seq と ToChannel

| N | 入力容量 | [Seq ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L46) | [ToChannel ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L54) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 0 | 535.8 ±2% | 6,561.5 ±0% | 12.246× |
| 4096 | 0 | 26,010.5 ±2% | 335,918 ±1% | 12.915× |

<details>
<summary>検定：Seq と ToChannel</summary>

| N | 入力容量 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 0 | sec/op | +1124.62% | p=0.000 n=10 |
| 64 | 0 | B/op | +6.53% | p=0.000 n=10 |
| 64 | 0 | allocs/op | -5.88% | p=0.000 n=10 |
| 4096 | 0 | sec/op | +1191.47% | p=0.000 n=10 |
| 4096 | 0 | B/op | +0.09% | p=0.000 n=10 |
| 4096 | 0 | allocs/op | -4.00% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 入力容量 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | 2,256 ±0% | 8 ±0% |
| 64 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 2,600 ±0% | 20 ±0% |
| 64 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L86) | 2,752 ±0% | 19 ±0% |
| 64 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L46) | 2,328 ±0% | 17 ±0% |
| 64 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L54) | 2,480 ±0% | 16 ±0% |
| 64 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | 2,784 ±0% | 8 ±0% |
| 64 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 3,128 ±0% | 20 ±0% |
| 64 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L86) | 3,280 ±0% | 19 ±0% |
| 4096 | 0 | [Direct](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | 177,616 ±0% | 16 ±0% |
| 4096 | 0 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 177,960 ±0% | 28 ±0% |
| 4096 | 0 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L86) | 178,112 ±0% | 27 ±0% |
| 4096 | 0 | [Seq](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L46) | 177,688 ±0% | 25 ±0% |
| 4096 | 0 | [ToChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L54) | 177,840 ±0% | 24 ±0% |
| 4096 | 64 | [Direct](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | 178,144 ±0% | 16 ±0% |
| 4096 | 64 | [FromChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 178,488 ±0% | 28 ±0% |
| 4096 | 64 | [RoundTrip](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L86) | 178,642 ±0% | 27 ±0% |

## Cache

### Seq / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 99.35 ±2% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 4,743.5 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |

### Seq / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 355.2 ±3% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 18,961.5 ±1% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 208.5 ±1% | 460.75 ±1% | 2.210× |
| 64 | 4 | 764.55 ±1% | 760.8 ±2% | 0.995× |
| 64 | 16 | 2,976.5 ±1% | 1,941 ±1% | 0.652× |
| 4096 | 1 | 11,332 ±3% | 23,783 ±1% | 2.099× |
| 4096 | 4 | 45,193.5 ±3% | 37,499.5 ±1% | 0.830× |
| 4096 | 16 | 180,656 ±3% | 94,495.5 ±1% | 0.523× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +120.98% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +240.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | ~ | p=0.123 n=10 |
| 64 | 4 | B/op | +431.25% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +136.36% | p=0.000 n=10 |
| 64 | 16 | sec/op | -34.79% | p=0.000 n=10 |
| 64 | 16 | B/op | +150.00% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +77.14% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +109.87% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +420.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -17.02% | p=0.000 n=10 |
| 4096 | 4 | B/op | +50131.25% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +218.18% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -47.69% | p=0.000 n=10 |
| 4096 | 16 | B/op | +17436.96% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +102.86% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,360 ±0% | 26 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,840 ±0% | 62 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 128,592 ±0% | 35 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 129,072 ±0% | 71 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |

### Seq / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 98.4 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 4,736.5 ±0% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |

### Seq / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 1,346 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 80,630 ±3% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 1,200 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 128,432 ±0% | 23 ±0% |

### Seq / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,243.5 ±1% | 1,454 ±2% | 1.169× |
| 64 | 4 | 4,870 ±1% | 1,761.5 ±2% | 0.362× |
| 64 | 16 | 19,468.5 ±3% | 2,914.5 ±2% | 0.150× |
| 4096 | 1 | 77,027 ±2% | 84,532.5 ±2% | 1.097× |
| 4096 | 4 | 307,426.5 ±2% | 100,147 ±3% | 0.326× |
| 4096 | 16 | 1,220,541.5 ±3% | 156,348.5 ±3% | 0.128× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +16.93% | p=0.000 n=10 |
| 64 | 1 | B/op | +811.76% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +240.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.83% | p=0.000 n=10 |
| 64 | 4 | B/op | +431.25% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +136.36% | p=0.000 n=10 |
| 64 | 16 | sec/op | -85.03% | p=0.000 n=10 |
| 64 | 16 | B/op | +150.00% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +77.14% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +9.74% | p=0.000 n=10 |
| 4096 | 1 | B/op | +94364.71% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +420.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -67.42% | p=0.000 n=10 |
| 4096 | 4 | B/op | +50131.25% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +218.18% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -87.19% | p=0.000 n=10 |
| 4096 | 16 | B/op | +17436.96% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +102.86% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,240 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,360 ±0% | 26 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 1,840 ±0% | 62 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 128,472 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 128,592 ±0% | 35 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 129,072 ±0% | 71 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |

### Seq2 / 負荷=0 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 101.2 ±1% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 5,438 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |

### Seq2 / 負荷=0 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 420.6 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 26,574.5 ±4% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 259,497 ±0% | 23 ±0% |

### Seq2 / 負荷=0 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 208.75 ±1% | 525.4 ±1% | 2.517× |
| 64 | 4 | 761.9 ±1% | 835.25 ±1% | 1.096× |
| 64 | 16 | 2,967.5 ±0% | 2,055 ±1% | 0.693× |
| 4096 | 1 | 11,242 ±0% | 31,322.5 ±3% | 2.786× |
| 4096 | 4 | 44,739 ±0% | 46,746.5 ±1% | 1.045× |
| 4096 | 16 | 178,560.5 ±0% | 111,429 ±2% | 0.624× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +151.69% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +240.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | +9.63% | p=0.000 n=10 |
| 64 | 4 | B/op | +828.12% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +136.36% | p=0.000 n=10 |
| 64 | 16 | sec/op | -30.75% | p=0.000 n=10 |
| 64 | 16 | B/op | +288.04% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +77.14% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +178.62% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190736.03% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +420.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | +4.49% | p=0.000 n=10 |
| 4096 | 4 | B/op | +101328.32% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +218.18% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -37.60% | p=0.000 n=10 |
| 4096 | 16 | B/op | +35244.57% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +102.86% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,376 ±0% | 26 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,856 ±0% | 62 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 259,537 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 259,656.5 ±0% | 35 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 260,136 ±0% | 71 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |

### Seq2 / 負荷=32 / 保存済み反復 Warm

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 101.05 ±2% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 5,483 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |
| 4096 | [Warm](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L93) | 40 ±0% | 3 ±0% |

### Seq2 / 負荷=32 / 初回保存 Fill

#### 時間

| N | 実装 | ns/op |
| --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 1,387.5 ±2% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 85,476 ±2% |

#### メモリ割り当て

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 2,216 ±0% | 14 ±0% |
| 4096 | [Fill](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L86) | 259,496 ±0% | 23 ±0% |

### Seq2 / 負荷=32 / 保存費用込み Reuse

#### 時間：Uncached と Cached

| N | 利用回数 | [Uncached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | [Cached ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 倍率 |
| --- | --- | --- | --- | --- |
| 64 | 1 | 1,252.5 ±2% | 1,495 ±2% | 1.194× |
| 64 | 4 | 4,934.5 ±2% | 1,818.5 ±3% | 0.369× |
| 64 | 16 | 19,736 ±2% | 3,012.5 ±1% | 0.153× |
| 4096 | 1 | 77,590.5 ±2% | 90,320 ±1% | 1.164× |
| 4096 | 4 | 310,575 ±3% | 106,826 ±2% | 0.344× |
| 4096 | 16 | 1,242,283.5 ±3% | 170,579 ±2% | 0.137× |

<details>
<summary>検定：Uncached と Cached</summary>

| N | 利用回数 | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- | --- |
| 64 | 1 | sec/op | +19.36% | p=0.000 n=10 |
| 64 | 1 | B/op | +1558.82% | p=0.000 n=10 |
| 64 | 1 | allocs/op | +240.00% | p=0.000 n=10 |
| 64 | 4 | sec/op | -63.15% | p=0.000 n=10 |
| 64 | 4 | B/op | +828.12% | p=0.000 n=10 |
| 64 | 4 | allocs/op | +136.36% | p=0.000 n=10 |
| 64 | 16 | sec/op | -84.74% | p=0.000 n=10 |
| 64 | 16 | B/op | +288.04% | p=0.000 n=10 |
| 64 | 16 | allocs/op | +77.14% | p=0.000 n=10 |
| 4096 | 1 | sec/op | +16.41% | p=0.000 n=10 |
| 4096 | 1 | B/op | +190735.29% | p=0.000 n=10 |
| 4096 | 1 | allocs/op | +420.00% | p=0.000 n=10 |
| 4096 | 4 | sec/op | -65.60% | p=0.000 n=10 |
| 4096 | 4 | B/op | +101328.12% | p=0.000 n=10 |
| 4096 | 4 | allocs/op | +218.18% | p=0.000 n=10 |
| 4096 | 16 | sec/op | -86.27% | p=0.000 n=10 |
| 4096 | 16 | B/op | +35244.57% | p=0.000 n=10 |
| 4096 | 16 | allocs/op | +102.86% | p=0.000 n=10 |

</details>

#### メモリ割り当て

| N | 利用回数 | 実装 | B/op | allocs/op |
| --- | --- | --- | --- | --- |
| 64 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,256 ±0% | 17 ±0% |
| 64 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 64 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,376 ±0% | 26 ±0% |
| 64 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 64 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 2,856 ±0% | 62 ±0% |
| 64 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |
| 4096 | 1 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 259,536 ±0% | 26 ±0% |
| 4096 | 1 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 136 ±0% | 5 ±0% |
| 4096 | 4 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 259,656 ±0% | 35 ±0% |
| 4096 | 4 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 256 ±0% | 11 ±0% |
| 4096 | 16 | [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 260,136 ±0% | 71 ±0% |
| 4096 | 16 | [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 736 ±0% | 35 ±0% |
