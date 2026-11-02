# 性能概要

[API 別の全結果](results/unit.md) / [利用場面別の全結果](results/scenario.md) / [測定条件](METHOD.md) / [実行環境](results/environment.md)。  
時間・割り当て量は中央値で、時間比は各表の先頭の実装を 1 とする（`~` は差が不明確）。  

## パイプライン

Seq / N=4096 / 段数=2 / 全件。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Loop](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L28) | 13,502.5 | 1× | 177,392 |
| [Pipeline](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/pipeline_scenario_test.go#L12) | 25,419 | 1.883× | 177,688 |

## Channel から読み込む

N=4096 / 入力容量=64。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Direct](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L61) | 112,686.5 | 1× | 178,144 |
| [FromChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L29) | 125,503 | 1.114× | 178,464 |

## Channel へ書き出す

N=4096。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Seq](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L52) | 25,203.5 | 1× | 177,688 |
| [ToChannel](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/channel_scenario_test.go#L19) | 292,317 | 11.598× | 177,801.5 |

## キャッシュ：軽い処理を 1 回使う

Seq / N=4096 / 負荷=0 / 利用回数=1 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L129) | 11,363.5 | 1× | 136 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L118) | 19,530 | 1.719× | 128,496 |

## キャッシュ：重い処理を 16 回使う

Seq2 / N=4096 / 負荷=32 / 利用回数=16 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L307) | 1,205,696.5 | 1× | 616 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/202662048ff6cedfe62770faab9d717b860c441d/benchmarks/cache_scenario_test.go#L296) | 101,250 | 0.084× | 259,560 |

軽い変換では、パイプライン接続や channel 変換の追加費用が目立つ。  
キャッシュは初回保存の費用と保持メモリを要するが、重い上流を繰り返し使うと再評価を省ける。  
この整数処理での結果を、外部 I/O や他の型の速度保証には使わない。  
