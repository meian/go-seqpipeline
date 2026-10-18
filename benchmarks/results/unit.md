# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
倍率は Library / Go で、1 を超えると Library が遅い。  
検定の `~` は差が不明確で、増減率の正値は Library の時間・割り当ての増加を表す。  

## 入力変換・出力収集

### FromMap：構築・評価

Go：map を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L193) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L183) | 倍率 |
| --- | --- | --- | --- |
| 64 | 323.95 ±3% | 316.2 ±8% | 0.976× |
| 4096 | 22,316 ±1% | 22,063 ±4% | 0.989× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L193) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L183) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L193) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L183) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.315 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.105 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L105) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L95) | 倍率 |
| --- | --- | --- | --- |
| 64 | 93.65 ±1% | 92.98 ±1% | 0.993× |
| 4096 | 4,738.5 ±7% | 4,745.5 ±0% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L105) | 40 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L95) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L105) | 40 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L95) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.063 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.928 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L221) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L209) | 倍率 |
| --- | --- | --- | --- |
| 64 | 92.985 ±1% | 92.725 ±1% | 0.997× |
| 4096 | 4,722 ±2% | 4,718 ±2% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L221) | 40 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L209) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L221) | 40 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L209) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.529 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.424 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L85) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L73) | 倍率 |
| --- | --- | --- | --- |
| 64 | 22.2 ±5% | 17.265 ±1% | 0.778× |
| 4096 | 997.3 ±3% | 995.85 ±3% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L85) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L73) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L85) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L73) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -22.23% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.542 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L286) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L269) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,297.5 ±2% | 2,301.5 ±2% | 1.002× |
| 4096 | 151,737 ±3% | 151,577 ±3% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L286) | 4,688 ±0% | 13 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L269) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L286) | 296,160 ±0% | 50 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L269) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.753 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.796 n=10 |
| 4096 | B/op | ~ | p=0.474 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L137) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L121) | 倍率 |
| --- | --- | --- | --- |
| 64 | 93.055 ±1% | 95.6 ±1% | 1.027× |
| 4096 | 4,741 ±1% | 5,093.5 ±0% | 1.074× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L137) | 40 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L121) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L137) | 40 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L121) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2.73% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +7.44% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L253) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L237) | 倍率 |
| --- | --- | --- | --- |
| 64 | 92.37 ±1% | 92.3 ±1% | 0.999× |
| 4096 | 4,715.5 ±3% | 4,723 ±2% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L253) | 40 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L237) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L253) | 40 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L237) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.781 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.325 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L170) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L153) | 倍率 |
| --- | --- | --- | --- |
| 64 | 252.6 ±5% | 251.1 ±3% | 0.994× |
| 4096 | 19,825.5 ±4% | 20,770.5 ±3% | 1.048× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L170) | 1,080 ±0% | 10 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L153) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L170) | 128,312 ±0% | 19 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/conversion_test.go#L153) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.481 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +4.77% | p=0.019 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 変換・選別・射影

### Filter：構築・評価

Go：`for range` 内で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L90) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L78) | 倍率 |
| --- | --- | --- | --- |
| 64 | 27.915 ±0% | 180.75 ±2% | 6.475× |
| 4096 | 2,050 ±1% | 9,643.5 ±3% | 4.704× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L90) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L78) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L90) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L78) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +547.50% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +370.41% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Filter2：構築・評価

Go：`for range` 内で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L148) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L136) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.42 ±0% | 195.55 ±2% | 5.851× |
| 4096 | 2,038.5 ±1% | 10,639.5 ±1% | 5.219× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L148) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L136) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L148) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L136) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +485.13% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +421.93% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Keys：構築・評価

Go：`for range` で添字を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L176) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L166) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.355 ±2% | 167.95 ±1% | 10.269× |
| 4096 | 944.45 ±2% | 9,425 ±0% | 9.979× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L176) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L166) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L176) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L166) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +926.90% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +897.94% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Map：構築・評価

Go：`for range` 内で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L62) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L52) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.555 ±2% | 194.2 ±1% | 11.062× |
| 4096 | 1,033 ±4% | 11,005.5 ±0% | 10.654× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L62) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L52) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L62) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L52) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1006.24% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +965.39% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Map2：構築・評価

Go：`for range` 内で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L120) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L108) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.705 ±0% | 197.7 ±2% | 10.569× |
| 4096 | 1,196 ±0% | 11,194 ±0% | 9.360× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L120) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L108) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L120) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L108) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +956.94% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +835.95% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Values：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L202) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L192) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.305 ±3% | 168.75 ±2% | 9.752× |
| 4096 | 999.5 ±3% | 9,443.5 ±0% | 9.448× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L202) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L192) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L202) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/transform_test.go#L192) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +875.15% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +844.82% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除いた部分スライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L85) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L73) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.725 ±3% | 173.25 ±0% | 11.017× |
| 4096 | 997.5 ±2% | 10,029 ±1% | 10.054× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L85) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L73) | 112 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L85) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L73) | 112 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1001.75% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +905.41% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き、元の添字・値を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L135) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L123) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.85 ±1% | 172.85 ±1% | 10.905× |
| 4096 | 1,163.5 ±0% | 10,166.5 ±0% | 8.738× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L135) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L123) | 112 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L135) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L123) | 112 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +990.54% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +773.79% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件の部分スライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L63) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L51) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.2935 ±5% | 49.525 ±2% | 15.037× |
| 4096 | 3.267 ±12% | 50.025 ±2% | 15.312× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L63) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L51) | 112 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L63) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L51) | 112 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1403.72% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1431.22% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件の添字・値を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L107) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L95) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.29 ±11% | 49.24 ±1% | 14.967× |
| 4096 | 3.637 ±13% | 50.035 ±2% | 13.757× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L107) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L95) | 112 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L107) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/limit_test.go#L95) | 112 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1396.66% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1275.72% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L122) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L110) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.155 ±3% | 3,833.5 ±4% | 146.569× |
| 4096 | 1,588 ±2% | 227,944.5 ±3% | 143.542× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L122) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L110) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L122) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L110) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14556.85% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14254.19% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L100) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.305 ±2% | 4,865 ±2% | 281.133× |
| 4096 | 1,023 ±2% | 296,623.5 ±3% | 289.955× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L100) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L88) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L100) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L88) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28013.26% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +28895.45% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L170) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L158) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.325 ±3% | 3,757.5 ±3% | 142.735× |
| 4096 | 1,588.5 ±1% | 225,936.5 ±3% | 142.233× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L170) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L158) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L170) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L158) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14173.50% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14123.26% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L146) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L134) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.085 ±2% | 3,821.5 ±2% | 146.502× |
| 4096 | 1,584.5 ±3% | 229,522.5 ±3% | 144.855× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L146) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L134) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L146) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L134) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14550.18% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14385.48% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L218) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L206) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.055 ±3% | 3,751.5 ±2% | 143.984× |
| 4096 | 1,589 ±2% | 227,388.5 ±2% | 143.102× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L218) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L206) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L218) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L206) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14298.39% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14210.16% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L194) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L182) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.26 ±3% | 4,868.5 ±3% | 282.068× |
| 4096 | 1,007.5 ±2% | 297,729 ±2% | 295.513× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L194) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L182) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L194) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L182) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28106.84% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29451.27% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L266) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L254) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.87 ±2% | 3,787 ±2% | 146.386× |
| 4096 | 1,572 ±1% | 227,648 ±3% | 144.814× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L266) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L254) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L266) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L254) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14538.58% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14381.42% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L242) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L230) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.11 ±3% | 3,808 ±2% | 145.845× |
| 4096 | 1,586.5 ±2% | 224,458.5 ±2% | 141.480× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L242) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L230) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L242) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/zip_test.go#L230) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14484.45% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14048.03% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

## パイプ接続

### Seq.Pipe：構築・評価

Go：`for range` 内で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L56) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L44) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.57 ±1% | 194.25 ±1% | 11.056× |
| 4096 | 1,046 ±2% | 11,009 ±0% | 10.525× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L56) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L44) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L56) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L44) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1005.58% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +952.49% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L84) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L72) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.44 ±4% | 4,905.5 ±3% | 281.279× |
| 4096 | 1,018.5 ±3% | 298,039.5 ±3% | 292.626× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L84) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L72) | 416 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L84) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L72) | 416 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28027.87% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29162.59% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L106) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L94) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.335 ±4% | 168.65 ±1% | 9.729× |
| 4096 | 983.5 ±3% | 9,445 ±0% | 9.603× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L106) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L94) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L106) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L94) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +872.89% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +860.35% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`for range` 内で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L128) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L116) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.69 ±1% | 197.3 ±1% | 10.556× |
| 4096 | 1,196 ±1% | 11,213.5 ±1% | 9.376× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L128) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L116) | 104 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L128) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipe_test.go#L116) | 104 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +955.64% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +837.58% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L208) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L193) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,625.5 ±2% | 6,630 ±1% | 1.179× |
| 4096 | 356,779 ±1% | 400,397 ±2% | 1.122× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L208) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L193) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L208) | 272 ±1% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L193) | 360 ±5% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +17.86% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +12.23% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L242) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L226) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,535.5 ±1% | 6,978.5 ±1% | 1.261× |
| 4096 | 338,991 ±2% | 413,741 ±2% | 1.221× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L242) | 136 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L226) | 256 ±0% | 4 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L242) | 136 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L226) | 256 ±0% | 4 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +26.07% | p=0.000 n=10 |
| 64 | B/op | +88.24% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +22.05% | p=0.000 n=10 |
| 4096 | B/op | +88.24% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

## キャッシュ

### Cache：保存済み反復 Warm

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L311) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L301) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.285 ±2% | 93.135 ±1% | 5.388× |
| 4096 | 1,014 ±3% | 4,740.5 ±0% | 4.675× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L311) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L301) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L311) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L301) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +438.82% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +367.50% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Cache：初回保存 Fill

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L287) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L273) | 倍率 |
| --- | --- | --- | --- |
| 64 | 240.3 ±1% | 243.4 ±2% | 1.013× |
| 4096 | 19,716 ±4% | 20,371 ±6% | 1.033× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L287) | 1,112 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L273) | 1,120 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L287) | 128,344 ±0% | 20 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L273) | 128,352 ±0% | 21 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.063 n=10 |
| 64 | B/op | +0.72% | p=0.000 n=10 |
| 64 | allocs/op | +9.09% | p=0.000 n=10 |
| 4096 | sec/op | ~ | p=0.143 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +5.00% | p=0.000 n=10 |

</details>

### Cache2：初回保存 Fill

Go：標準 iterator の組を `append` で保存。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L335) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L321) | 倍率 |
| --- | --- | --- | --- |
| 64 | 304.8 ±2% | 309.95 ±2% | 1.017× |
| 4096 | 28,809.5 ±8% | 30,187 ±5% | 1.048× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L335) | 2,120 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L321) | 2,136 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L335) | 259,402 ±0% | 20 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L321) | 259,417 ±0% | 21 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.69% | p=0.000 n=10 |
| 64 | B/op | +0.75% | p=0.000 n=10 |
| 64 | allocs/op | +9.09% | p=0.000 n=10 |
| 4096 | sec/op | +4.78% | p=0.045 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +5.00% | p=0.000 n=10 |

</details>

### Cache2：保存済み反復 Warm

Go：保存済みの組のスライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L359) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L349) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.585 ±1% | 96.805 ±1% | 5.505× |
| 4096 | 1,027 ±8% | 5,461 ±1% | 5.317× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L359) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L349) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L359) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L349) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +450.50% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +431.74% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
