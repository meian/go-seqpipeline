# 性能概要

[API 別の全結果](results/unit.md) / [利用場面別の全結果](results/scenario.md) / [測定条件](METHOD.md) / [実行環境](results/environment.md)。  
時間・割り当て量は中央値で、時間比は各表の先頭の実装を 1 とする（`~` は差が不明確）。  

## パイプライン

Seq / N=4096 / 段数=2 / 全件。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Loop](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L94) | 13,601.5 | 1× | 177,391 |
| [Pipeline](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/pipeline_scenario_test.go#L10) | 25,241.5 | 1.856× | 177,688 |

## Channel から読み込む

N=4096 / 入力容量=64。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Direct](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L88) | 114,584.5 | 1× | 178,144 |
| [FromChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L141) | 187,568.5 | 1.637× | 178,488.5 |

## Channel へ書き出す

N=4096 / 入力容量=0。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Seq](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L17) | 26,019 | 1× | 177,688 |
| [ToChannel](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/channel_scenario_test.go#L52) | 337,299.5 | 12.964× | 177,840.5 |

## キャッシュ：軽い処理を 1 回使う

Seq / N=4096 / 負荷=0 / 利用回数=1 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L250) | 11,415 | 1× | 136 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L131) | 23,138.5 | 2.027× | 128,472 |

## キャッシュ：重い処理を 16 回使う

Seq2 / N=4096 / 負荷=32 / 利用回数=16 / 保存費用込み Reuse。  

| 実装 | ns/op | 時間比 | B/op |
| --- | --- | --- | --- |
| [Uncached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L595) | 1,247,180.5 | 1× | 616 |
| [Cached](https://github.com/meian/go-seqpipeline/blob/b57287562a850088ae8c5b15690600755f58f4eb/benchmarks/cache_scenario_test.go#L476) | 174,471.5 | 0.140× | 260,016 |

軽い変換では、パイプライン接続や channel 変換の追加費用が目立つ。  
キャッシュは初回保存の費用と保持メモリを要するが、重い上流を繰り返し使うと再評価を省ける。  
この整数処理での結果を、外部 I/O や他の型の速度保証には使わない。  
