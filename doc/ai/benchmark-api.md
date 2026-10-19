# 公開 API の測定対応（保守者向け）

対象コミットは [測定環境](../../benchmarks/results/environment.md) に記録する。  
`seq` の公開関数 21 件と公開メソッド 16 件、計 37 件を今回のコードから棚卸しした。  
非公開の設定値型が返す公開メソッドも含む。  
32 種の API/phase を N64 / N4096 の計 64 条件・Go / Library の 128 ケースで比較する。  
設定値の生成 3 関数は関連する Unit で、Cache / Cache2 は利用目的に合う Scenario で測定する。  
この表は今回のスナップショットの対応であり、通常の機能変更時に即時追従を要求する検証ではない。  

| 公開 API | Unit の api キー | 測定範囲・代替理由 |
| --- | --- | --- |
| `Cache` | — | 標準 API の代替ではないため Unit は設けず、Scenario の Fill / Warm / Reuse で測定 |
| `Cache2` | — | Cache と同じく Scenario の Fill / Warm / Reuse で測定 |
| `FromChannel` | `FromChannel` | channel 生成・goroutine 起動・全送受信・終了・チェックサム |
| `ToChannel` | `ToChannel` | channel 生成・goroutine 起動・全送受信・終了・チェックサム |
| `FromSlice` | `FromSlice` | 構築・全件評価・チェックサム |
| `ToSlice` | `ToSlice` | 収集・割り当て・結果保持・チェックサム |
| `FromSeq` | `FromSeq` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `ToSeq` | `ToSeq` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `FromMap` | `FromMap` | 構築・全件評価・チェックサム |
| `ToMap` | `ToMap` | 収集・割り当て・結果保持・チェックサム |
| `FromSeq2` | `FromSeq2` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `ToSeq2` | `ToSeq2` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `Filter` | `Filter` | 構築・全件評価・チェックサム |
| `Filter2` | `Filter2` | 構築・全件評価・チェックサム |
| `Map` | `Map` | 構築・全件評価・チェックサム |
| `Map2` | `Map2` | 構築・全件評価・チェックサム |
| `Keys` | `Keys` | 構築・全件評価・チェックサム |
| `Values` | `Values` | 構築・全件評価・チェックサム |
| `Seq.Pipe` | `Seq.Pipe` | 変換関数への転送なので、対応する Map / Zip / Values / Map2 の接続・評価に含める |
| `Seq.Pipe2` | `Seq.Pipe2` | 変換関数への転送なので、対応する Map / Zip / Values / Map2 の接続・評価に含める |
| `Seq2.Pipe` | `Seq2.Pipe` | 変換関数への転送なので、対応する Map / Zip / Values / Map2 の接続・評価に含める |
| `Seq2.Pipe2` | `Seq2.Pipe2` | 変換関数への転送なので、対応する Map / Zip / Values / Map2 の接続・評価に含める |
| `Skip` | `Skip.Operator / Skip.Operator2` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `skipOperator.Operator` | `Skip.Operator` | 構築・全件評価・チェックサム |
| `skipOperator.Operator2` | `Skip.Operator2` | 構築・全件評価・チェックサム |
| `Take` | `Take.Operator / Take.Operator2` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `takeOperator.Operator` | `Take.Operator` | 設定・接続・先頭 8 件までの評価・チェックサム |
| `takeOperator.Operator2` | `Take.Operator2` | 設定・接続・先頭 8 件までの評価・チェックサム |
| `Zip` | `Zip.Operator` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `zipOperator.AlignBy` | `Zip.AlignBy` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipOperator.WithDefaultValue` | `Zip.WithDefaultValue` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipOperator.WithDefaultKey` | `Zip.WithDefaultKey` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipOperator.Operator` | `Zip.Operator` | 構築・全件評価・チェックサム |
| `zipKeyOperator.AlignBy` | `ZipKey.AlignBy` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipKeyOperator.WithDefaultValue` | `ZipKey.WithDefaultValue` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipKeyOperator.WithDefaultKey` | `ZipKey.WithDefaultKey` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `zipKeyOperator.Operator` | `ZipKey.Operator` | 構築・全件評価・チェックサム |

`Seq` / `Seq2` と `Operator` / `Operator2` / `ToSeqOperator` / `ToSeq2Operator` は関数型で、独立した関数・メソッドを追加しない。  
それぞれ入力・Map / Map2・Zip・Values と Pipe / Pipe2 の測定で利用する。  
`ZipAlign` と `AlignShortest` / `AlignLeft` / `AlignRight` / `AlignLongest` は型・定数であり、対応する Zip ケースに含める。  
各ケースの詳しい条件は [測定方法](../../benchmarks/METHOD.md) を参照。  

## 測定コードの配置

| API・機能 | 測定・結果検証のファイル |
| --- | --- |
| Map / Map2 | [map_test.go](../../benchmarks/map_test.go) |
| Filter / Filter2 | [filter_test.go](../../benchmarks/filter_test.go) |
| Keys / Values | [keys_test.go](../../benchmarks/keys_test.go)、[values_test.go](../../benchmarks/values_test.go) |
| Take / Skip と Operator / Operator2 | [take_test.go](../../benchmarks/take_test.go)、[skip_test.go](../../benchmarks/skip_test.go) |
| FromSlice / ToSlice | [from_slice_test.go](../../benchmarks/from_slice_test.go)、[to_slice_test.go](../../benchmarks/to_slice_test.go) |
| FromSeq / FromSeq2 | [from_seq_test.go](../../benchmarks/from_seq_test.go) |
| ToSeq / ToSeq2 | [to_seq_test.go](../../benchmarks/to_seq_test.go) |
| FromMap / ToMap | [from_map_test.go](../../benchmarks/from_map_test.go)、[to_map_test.go](../../benchmarks/to_map_test.go) |
| FromChannel / ToChannel | [from_channel_test.go](../../benchmarks/from_channel_test.go)、[to_channel_test.go](../../benchmarks/to_channel_test.go) |
| Zip と設定メソッド | [zip_test.go](../../benchmarks/zip_test.go) |
| Seq / Seq2 の Pipe / Pipe2 | [pipe_test.go](../../benchmarks/pipe_test.go) |
| Cache / Cache2、Fill / Warm / Reuse | [cache_scenario_test.go](../../benchmarks/cache_scenario_test.go) |
| Channel を含むパイプライン | [channel_scenario_test.go](../../benchmarks/channel_scenario_test.go) |
| Map / Filter / Take を接続するパイプライン | [pipeline_scenario_test.go](../../benchmarks/pipeline_scenario_test.go) |

[benchmark_helper_test.go](../../benchmarks/benchmark_helper_test.go) は入力データの生成、結果保持用の変数、要素の変換・判定だけを共有する。  
Unit のケースは機能別の `*_test.go`、Scenario のケースは機能別の `*_scenario_test.go` に分離する。  
共有の送信処理は `channel_helper_test.go` に置き、異なる測定分類のケースを同じファイルへ混在させない。  
ケースの登録表、共通のベンチマークループ、テストのディスパッチャーは設けない。  

## 処理と入口

各 API・実装の処理全体を `mapLibrary` / `mapGo` のような型付き関数に置く。  
通常の `BenchmarkMapLibrary` / `TestMapLibrary` は、どちらも `mapLibrary` を直接呼ぶ。  
Seq の生成・演算子の構築・評価・収集・結果の加算は処理本体に含め、入力データだけを呼び出し元で準備する。  
Cache / Cache2 は Scenario のみで扱い、`scenarioCacheSeqFill` / `scenarioCacheSeqWarm` のように段階を分ける。  
Scenario は `scenarioPipelineSeqLibrary` / `scenarioPipelineSeq2Library` のように、種類と実装ごとに分ける。  

ベンチマーク関数内の各 `b.Run` は条件を明記し、`b.Loop()` 内から処理本体を直接呼ぶ。  
`class=Unit` / `class=ScenarioPipeline` / `class=ScenarioChannel` / `class=ScenarioCache` で分類し、入力や段数などの比較条件も名前へ残す。  
テストは各処理につき代表入力 1 件で処理本体を直接呼び、簡単な期待値で疎通を確認する。  
公開表のリンクは測定コミットの関数宣言行から自動生成するため、行番号を手動で転記しない。  
