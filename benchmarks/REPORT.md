# 性能概要

[API 別の全結果](results/unit.md) / [利用場面別の全結果](results/scenario.md) / [測定条件](METHOD.md) / [実行環境](results/environment.md)。  
時間・割り当て量は中央値で、時間比は各表の先頭の実装を 1 とする（`~` は差が不明確）。  

## パイプライン

Seq / N=4096 / 段数=2 / 全件。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Loop](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L42) | 13,779 | 1× | 177,392 |
| [Pipeline](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/pipeline_test.go#L12) | 26,466.5 | 1.921× | 177,688 |

## Channel から読み込む

N=4096 / 入力容量=64。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Direct](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L63) | 109,611 | 1× | 178,144 |
| [FromChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L76) | 182,735.5 | 1.667× | 178,488 |

## Channel へ書き出す

N=4096 / 入力容量=0。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Seq](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L46) | 26,010.5 | 1× | 177,688 |
| [ToChannel](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/channel_test.go#L54) | 335,918 | 12.915× | 177,840 |

## キャッシュ：軽い処理を 1 回使う

Seq / N=4096 / 負荷=0 / 利用回数=1 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 11,332 | 1× | 136 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 23,783 | 2.099× | 128,472 |

## キャッシュ：重い処理を 16 回使う

Seq2 / N=4096 / 負荷=32 / 利用回数=16 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L116) | 1,242,283.5 | 1× | 736 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/aea977304b759c54fd5fbf16aa1170f75e606e4f/benchmarks/cache_test.go#L107) | 170,579 | 0.137× | 260,136 |

軽い変換では、パイプライン接続や channel 変換の追加費用が目立つ。  
キャッシュは初回保存の費用と保持メモリを要するが、重い上流を繰り返し使うと再評価を省ける。  
この整数処理での結果を、外部 I/O や他の型の速度保証には使わない。  
