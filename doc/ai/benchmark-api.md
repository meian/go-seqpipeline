# 公開 API の測定対応（保守者向け）

対象コミットは [測定環境](../../benchmarks/results/environment.md) に記録する。  
`seq` の公開関数 21 件と公開メソッド 16 件、計 37 件を今回のコードから棚卸しした。  
非公開の設定値型が返す公開メソッドも含む。  
34 種の API 名・36 種の API/phase を N64 / N4096 の計 72 条件を Go / Library の 144 ケースで比較し、設定値の生成 3 関数は関連する評価ケースで代替する。  
この表は今回のスナップショットの対応であり、通常の機能変更時に即時追従を要求する検証ではない。  

| 公開 API | Unit の api キー | 測定範囲・代替理由 |
| --- | --- | --- |
| `Cache` | `Cache` | Fill（作成・全件保存） / Warm（保存済みの反復）を分離 |
| `Cache2` | `Cache2` | Fill（作成・全件保存） / Warm（保存済みの反復）を分離 |
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
| FromSlice / FromSeq / ToSeq / ToSlice / FromMap / FromSeq2 / ToSeq2 / ToMap | [conversion_test.go](../../benchmarks/conversion_test.go) |
| Map / Map2 / Filter / Filter2 / Keys / Values | [transform_test.go](../../benchmarks/transform_test.go) |
| Take / Skip と Operator / Operator2 | [limit_test.go](../../benchmarks/limit_test.go) |
| Zip と設定メソッド | [zip_test.go](../../benchmarks/zip_test.go) |
| Seq / Seq2 の Pipe / Pipe2 | [pipe_test.go](../../benchmarks/pipe_test.go) |
| FromChannel / ToChannel、相互変換 Scenario | [channel_test.go](../../benchmarks/channel_test.go) |
| Cache / Cache2、Fill / Warm / Reuse Scenario | [cache_test.go](../../benchmarks/cache_test.go) |
| Map / Filter / Take を接続する Pipeline Scenario | [pipeline_test.go](../../benchmarks/pipeline_test.go) |

共通の Unit ランナーは [benchmark_helper_test.go](../../benchmarks/benchmark_helper_test.go) に置き、全件を従来の `Unit/api=...` という名前に `ref`（一般的な Go の書き方）と `impl=Go / Library` を付けて実行する。  

Go 側の実装も同じ機能ファイルの `*UnitReferences` に置く。  
各 API・phase の比較相手は登録時に検証し、対応漏れ・重複・不要な登録をテストで検出する。  
公開結果には比較相手の書き方を示し、一般的な Go と Library の時間・割り当て・検定を併記する。  
