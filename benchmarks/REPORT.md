# 性能概要

[API 別の全結果](results/unit.md) / [利用場面別の全結果](results/scenario.md) / [測定条件](METHOD.md) / [実行環境](results/environment.md)。  
時間・割り当て量は中央値で、時間比は各表の先頭の実装を 1 とする（`~` は差が不明確）。  

## パイプライン

Seq / N=4096 / 段数=2 / 全件。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Loop](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L276) | 14,359.5 | 1× | 177,391 |
| [Pipeline](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/pipeline_test.go#L10) | 25,624.5 | 1.784× | 177,688 |

## Channel から読み込む

N=4096 / 入力容量=64。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Direct](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L225) | 110,092.5 | 1× | 178,144 |
| [FromChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L370) | 182,121.5 | 1.654× | 178,488 |

## Channel へ書き出す

N=4096 / 入力容量=0。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Seq](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L28) | 26,133 | 1× | 177,689 |
| [ToChannel](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/channel_test.go#L126) | 331,992.5 | 12.704× | 177,841 |

## キャッシュ：軽い処理を 1 回使う

Seq / N=4096 / 負荷=0 / 利用回数=1 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1417) | 11,038 | 1× | 136 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L1064) | 23,230 | 2.105× | 128,472 |

## キャッシュ：重い処理を 16 回使う

Seq2 / N=4096 / 負荷=32 / 利用回数=16 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2444) | 1,238,235 | 1× | 616 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/2469d92303644cfd5dbf1e8b265b18782c9f2ce7/benchmarks/cache_test.go#L2091) | 171,877 | 0.139× | 260,016 |

軽い変換では、パイプライン接続や channel 変換の追加費用が目立つ。  
キャッシュは初回保存の費用と保持メモリを要するが、重い上流を繰り返し使うと再評価を省ける。  
この整数処理での結果を、外部 I/O や他の型の速度保証には使わない。  
