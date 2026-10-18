# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
倍率は Library / Go で、1 を超えると Library が遅い。  
検定の `~` は差が不明確で、増減率の正値は Library の時間・割り当ての増加を表す。  

## キャッシュ

### Cache：保存済み反復 Warm

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L248) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L162) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.735 ±12% | 93.335 ±1% | 5.263× |
| 4096 | 994.9 ±4% | 4,741.5 ±1% | 4.766× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L248) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L162) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L248) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L162) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +426.28% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +376.58% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Cache：初回保存 Fill

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L88) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 132.3 ±3% | 177.85 ±3% | 1.344× |
| 4096 | 8,009 ±2% | 7,880 ±5% | 0.984× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L88) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L12) | 1,064 ±0% | 8 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L88) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L12) | 128,296 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +34.43% | p=0.000 n=10 |
| 64 | B/op | +4.72% | p=0.000 n=10 |
| 64 | allocs/op | +14.29% | p=0.000 n=10 |
| 4096 | sec/op | ~ | p=0.197 n=10 |
| 4096 | B/op | +0.04% | p=0.000 n=10 |
| 4096 | allocs/op | +6.25% | p=0.000 n=10 |

</details>

### Cache2：初回保存 Fill

Go：標準 iterator の組を `append` で保存。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L454) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L336) | 倍率 |
| --- | --- | --- | --- |
| 64 | 185.25 ±3% | 202.25 ±4% | 1.092× |
| 4096 | 14,818 ±2% | 14,631.5 ±2% | 0.987× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L454) | 1,984 ±0% | 5 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L336) | 2,080 ±0% | 8 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L454) | 259,264 ±0% | 14 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L336) | 259,360 ±0% | 17 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +9.18% | p=0.000 n=10 |
| 64 | B/op | +4.84% | p=0.000 n=10 |
| 64 | allocs/op | +60.00% | p=0.000 n=10 |
| 4096 | sec/op | ~ | p=0.280 n=10 |
| 4096 | B/op | +0.04% | p=0.000 n=10 |
| 4096 | allocs/op | +21.43% | p=0.000 n=10 |

</details>

### Cache2：保存済み反復 Warm

Go：保存済みの組のスライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L642) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L556) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.3 ±2% | 95.42 ±1% | 5.516× |
| 4096 | 1,025 ±0% | 5,360.5 ±0% | 5.230× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L642) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L556) | 40 ±0% | 2 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L642) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L556) | 40 ±0% | 2 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +451.56% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +422.98% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

## 変換・選別・射影

### Filter：構築・評価

Go：`for range` 内で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 32.98 ±0% | 28.8 ±1% | 0.873× |
| 4096 | 2,055.5 ±0% | 1,696 ±0% | 0.825× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -12.67% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -17.49% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`for range` 内で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L266) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L181) | 倍率 |
| --- | --- | --- | --- |
| 64 | 32.99 ±0% | 29.235 ±1% | 0.886× |
| 4096 | 2,054.5 ±0% | 1,731.5 ±0% | 0.843× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L266) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L181) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L266) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/filter_test.go#L181) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -11.38% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -15.72% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`for range` で添字を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.83 ±1% | 16.225 ±2% | 1.025× |
| 4096 | 929.85 ±2% | 929.5 ±1% | 1.000× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/keys_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2.50% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.753 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`for range` 内で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.245 ±1% | 17.74 ±1% | 1.029× |
| 4096 | 1,142 ±1% | 1,026 ±4% | 0.898× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2.87% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -10.16% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`for range` 内で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L264) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L179) | 倍率 |
| --- | --- | --- | --- |
| 64 | 19.22 ±0% | 18.615 ±0% | 0.969× |
| 4096 | 1,216 ±0% | 1,207 ±0% | 0.993× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L264) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L179) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L264) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/map_test.go#L179) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -3.15% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -0.74% | p=0.001 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.135 ±2% | 17.43 ±20% | 1.017× |
| 4096 | 1,020 ±1% | 1,012 ±2% | 0.992× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/values_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.72% | p=0.011 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -0.78% | p=0.033 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L96) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,758 ±1% | 6,852.5 ±2% | 1.190× |
| 4096 | 342,053.5 ±1% | 408,547 ±2% | 1.194× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L96) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L96) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.01% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +19.44% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,598 ±2% | 7,007.5 ±2% | 1.252× |
| 4096 | 339,626.5 ±1% | 421,100 ±1% | 1.240× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L95) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L10) | 304 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L95) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_channel_test.go#L10) | 304 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +25.18% | p=0.000 n=10 |
| 64 | B/op | +90.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +23.99% | p=0.000 n=10 |
| 4096 | B/op | +90.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：map を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L119) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 320.45 ±6% | 321.65 ±3% | 1.004× |
| 4096 | 22,068.5 ±1% | 21,886 ±2% | 0.992× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L119) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L9) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L119) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_map_test.go#L9) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.529 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.315 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.16 ±2% | 17.47 ±3% | 1.018× |
| 4096 | 999.15 ±2% | 1,011 ±1% | 1.012× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.165 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.109 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L265) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L180) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.48 ±2% | 17.36 ±3% | 0.993× |
| 4096 | 1,080 ±4% | 1,087.5 ±3% | 1.007× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L265) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L180) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L265) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_seq_test.go#L180) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.470 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.564 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L93) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.99 ±1% | 17.095 ±3% | 1.006× |
| 4096 | 1,014 ±2% | 994.15 ±3% | 0.980× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L93) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L9) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L93) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/from_slice_test.go#L9) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.566 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.052 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L99) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,237 ±1% | 2,345 ±1% | 1.048× |
| 4096 | 146,013.5 ±1% | 151,096.5 ±1% | 1.035× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L99) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L13) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L99) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_map_test.go#L13) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +4.83% | p=0.000 n=10 |
| 64 | B/op | +0.86% | p=0.000 n=10 |
| 64 | allocs/op | +18.18% | p=0.000 n=10 |
| 4096 | sec/op | +3.48% | p=0.000 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +4.17% | p=0.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.315 ±2% | 17.305 ±2% | 0.999× |
| 4096 | 1,000.4 ±2% | 998.3 ±3% | 0.998× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.810 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.782 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L265) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L180) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.465 ±2% | 17.47 ±3% | 1.000× |
| 4096 | 1,083 ±2% | 1,094.5 ±2% | 1.011× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L265) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L180) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L265) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_seq_test.go#L180) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.725 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.128 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L98) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 156.85 ±4% | 259.1 ±3% | 1.652× |
| 4096 | 10,313.5 ±8% | 19,665.5 ±4% | 1.907× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L98) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L12) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L98) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/to_slice_test.go#L12) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +65.19% | p=0.000 n=10 |
| 64 | B/op | +6.30% | p=0.000 n=10 |
| 64 | allocs/op | +42.86% | p=0.000 n=10 |
| 4096 | sec/op | +90.68% | p=0.000 n=10 |
| 4096 | B/op | +0.05% | p=0.000 n=10 |
| 4096 | allocs/op | +18.75% | p=0.000 n=10 |

</details>

## パイプ接続

### Seq.Pipe：構築・評価

Go：`for range` 内で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.305 ±1% | 17.815 ±2% | 1.029× |
| 4096 | 1,140 ±1% | 1,041 ±0% | 0.913× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +2.95% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -8.68% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L264) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L179) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.15 ±2% | 4,894.5 ±3% | 285.394× |
| 4096 | 1,012 ±1% | 297,447.5 ±2% | 293.920× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L264) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L179) | 464 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L264) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L179) | 464 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28439.36% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29292.05% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：スライスを `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L433) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L348) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.925 ±1% | 17.07 ±2% | 1.009× |
| 4096 | 1,013.5 ±1% | 986.2 ±2% | 0.973× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L433) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L348) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L433) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L348) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.289 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -2.69% | p=0.001 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`for range` 内で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L602) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L517) | 倍率 |
| --- | --- | --- | --- |
| 64 | 19.2 ±0% | 18.605 ±1% | 0.969× |
| 4096 | 1,216 ±0% | 1,207 ±0% | 0.993× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L602) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L517) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L602) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipe_test.go#L517) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -3.10% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -0.74% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除いた部分スライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.155 ±1% | 30.705 ±2% | 2.026× |
| 4096 | 1,020 ±1% | 1,984 ±1% | 1.945× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +102.61% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +94.51% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き、元の添字・値を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L265) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L180) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.785 ±0% | 30.805 ±2% | 1.732× |
| 4096 | 2,016.5 ±4% | 1,975.5 ±2% | 0.980× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L265) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L180) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L265) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/skip_test.go#L180) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +73.21% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -2.03% | p=0.014 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件の部分スライスを反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L95) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2.6725 ±8% | 3.4635 ±0% | 1.296× |
| 4096 | 2.6955 ±5% | 3.536 ±5% | 1.312× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L95) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L95) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +29.60% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +31.18% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件の添字・値を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L265) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L180) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2.8565 ±11% | 3.4665 ±0% | 1.214× |
| 4096 | 2.9305 ±11% | 3.4705 ±5% | 1.184× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L265) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L180) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L265) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/take_test.go#L180) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +21.35% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +18.43% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L264) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L178) | 倍率 |
| --- | --- | --- | --- |
| 64 | 26.09 ±1% | 3,833 ±1% | 146.915× |
| 4096 | 1,587.5 ±1% | 230,551 ±2% | 145.229× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L264) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L178) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L264) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L178) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14591.45% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14422.90% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L94) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.24 ±2% | 4,866 ±3% | 282.251× |
| 4096 | 1,022.5 ±2% | 295,199 ±2% | 288.703× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L94) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L9) | 464 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L94) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L9) | 464 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28125.06% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +28770.32% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L614) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L528) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.915 ±1% | 3,755 ±3% | 144.897× |
| 4096 | 1,582 ±1% | 226,427 ±2% | 143.127× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L614) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L528) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L614) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L528) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14389.68% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14212.71% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L439) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L353) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.615 ±1% | 3,817 ±2% | 149.014× |
| 4096 | 1,585 ±2% | 229,714 ±2% | 144.930× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L439) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L353) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L439) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L353) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14801.42% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14393.00% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L958) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L872) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.555 ±3% | 3,779 ±2% | 147.877× |
| 4096 | 1,587 ±1% | 226,857.5 ±2% | 142.947× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L958) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L872) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L958) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L872) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14687.71% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14194.74% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L788) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L703) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17 ±2% | 4,888 ±3% | 287.529× |
| 4096 | 1,017 ±2% | 296,838 ±3% | 291.876× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L788) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L703) | 464 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L788) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L703) | 464 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +28652.94% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +29087.61% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1308) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1222) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.9 ±3% | 3,872.5 ±2% | 149.517× |
| 4096 | 1,583.5 ±1% | 234,131.5 ±2% | 147.857× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1308) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1222) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1308) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1222) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14851.74% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14685.70% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1133) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1047) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.88 ±1% | 3,811 ±3% | 147.257× |
| 4096 | 1,584 ±1% | 227,553.5 ±1% | 143.658× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1133) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1047) | 512 ±0% | 14 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1133) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/zip_test.go#L1047) | 512 ±0% | 14 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +14625.66% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +14265.75% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
