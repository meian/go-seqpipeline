# 公開 API の測定対応（保守者向け）

対象コミットは [測定環境](../../benchmarks/results/environment.md) に記録する。  
`seq` と `op` の公開関数 26 件と公開メソッド 32 件、計 58 件を棚卸しした。  
`seq` は入出力変換・生成の 11 関数と接続の 4 メソッド、`op` はオペレーター・合成の 15 関数と設定値・合成値の 28 メソッドを扱う。  
非公開の設定値型が返す公開メソッドも含む。  
40 種の API/phase を N64 / N4096 の計 80 条件・160 ケースで比較する。  
Compose / Compose2 は Pipe / Composed、その他は Go / Library を対照とする。  
設定値・合成開始値の生成 5 関数と MapToSeq / MapToSeq2 は関連する Unit で、Cache / Cache2 は利用目的に合う Scenario で測定する。  
この表は今回のスナップショットの対応であり、通常の機能変更時に即時追従を要求する検証ではない。  

| 公開 API | Unit の api キー | 測定範囲・代替理由 |
| --- | --- | --- |
| `op.Cache` | — | 標準 API の代替ではないため Unit は設けず、Scenario の Fill / Warm / Reuse で測定 |
| `op.Cache2` | — | Cache と同じく Scenario の Fill / Warm / Reuse で測定 |
| `seq.FromChannel` | `FromChannel` | channel 生成・goroutine 起動・全送受信・終了・チェックサム |
| `seq.ToChannel` | `ToChannel` | channel 生成・goroutine 起動・全送受信・終了・チェックサム |
| `seq.FromRange` | `FromRange.Signed / FromRange.Unsigned` | int / uint の昇順・降順を step=±1 で構築・全件評価・加算 |
| `seq.FromSlice` | `FromSlice` | 構築・全件評価・チェックサム |
| `seq.ToSlice` | `ToSlice` | 収集・割り当て・結果保持・チェックサム |
| `seq.FromSeq` | `FromSeq` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `seq.ToSeq` | `ToSeq` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `seq.FromMap` | `FromMap` | 構築・全件評価・チェックサム |
| `seq.ToMap` | `ToMap` | 収集・割り当て・結果保持・チェックサム |
| `seq.FromSeq2` | `FromSeq2` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `seq.ToSeq2` | `ToSeq2` | 型変換だけの時間は意味が薄いため、変換後の全件評価を含める |
| `op.Filter` | `Filter` | 構築・全件評価・チェックサム |
| `op.Filter2` | `Filter2` | 構築・全件評価・チェックサム |
| `op.Map` | `Map` | 構築・全件評価・チェックサム |
| `op.Map2` | `Map2` | 構築・全件評価・チェックサム |
| `op.MapToSeq2` | `Seq.Pipe2` | 接続・全件反復時の変換を測り、単独ケースの重複を避ける |
| `op.MapToSeq` | `Seq2.Pipe` | 接続・全件反復時の変換を測り、単独ケースの重複を避ける |
| `op.Keys` | `Keys` | 構築・全件評価・チェックサム |
| `op.Values` | `Values` | 構築・全件評価・チェックサム |
| `seq.Seq.Pipe` | `Seq.Pipe` | Map の接続・全件評価 |
| `seq.Seq.Pipe2` | `Seq.Pipe2` | MapToSeq2 による値と変換値の組への接続・全件評価 |
| `seq.Seq2.Pipe` | `Seq2.Pipe` | slices.All を入力とし、MapToSeq で値を取り出す接続・全件評価 |
| `seq.Seq2.Pipe2` | `Seq2.Pipe2` | Map2 の接続・全件評価 |
| `op.Skip` | `Skip.Operator / Skip.Operator2` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `op.skipOperator.Operator` | `Skip.Operator` | 構築・全件評価・チェックサム |
| `op.skipOperator.Operator2` | `Skip.Operator2` | 構築・全件評価・チェックサム |
| `op.Take` | `Take.Operator / Take.Operator2` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `op.takeOperator.Operator` | `Take.Operator` | 設定・接続・先頭 8 件までの評価・チェックサム |
| `op.takeOperator.Operator2` | `Take.Operator2` | 設定・接続・先頭 8 件までの評価・チェックサム |
| `op.Zip` | `Zip.Operator` | 設定値の生成だけでは実利用の評価費用を表さないため、Operator の生成・評価に含める |
| `op.zipOperator.AlignBy` | `Zip.AlignBy` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipOperator.WithDefaultValue` | `Zip.WithDefaultValue` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipOperator.WithDefaultKey` | `Zip.WithDefaultKey` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipOperator.Operator` | `Zip.Operator` | 構築・全件評価・チェックサム |
| `op.zipKeyOperator.AlignBy` | `ZipKey.AlignBy` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipKeyOperator.WithDefaultValue` | `ZipKey.WithDefaultValue` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipKeyOperator.WithDefaultKey` | `ZipKey.WithDefaultKey` | 設定だけの時間は意味が薄いため、不等長入力の設定・接続・評価で代替 |
| `op.zipKeyOperator.Operator` | `ZipKey.Operator` | 構築・全件評価・チェックサム |
| `op.Compose` | `Compose.Seq / Compose.Seq2` | 開始値だけの構築は意味が薄いため、合成・接続・評価に含める |
| `op.Compose2` | `Compose2.Seq / Compose2.Seq2` | Compose と同じく合成・接続・評価に含める |
| `op.compose.Pipe` | `Compose.Seq` | Map から単一値入力の合成を開始 |
| `op.compose.Pipe2` | `Compose.Seq2` | MapToSeq2 から単一値入力の合成を開始 |
| `op.compose2.Pipe` | `Compose2.Seq` | MapToSeq から組入力の合成を開始 |
| `op.compose2.Pipe2` | `Compose2.Seq2` | Map2 から組入力の合成を開始 |
| `op.composed.Pipe` | `Compose.Seq` | 単一値から単一値への操作を追加・評価 |
| `op.composed.Pipe2` | `Compose.Seq` | 単一値から組への操作を追加・評価 |
| `op.composed.Operator` | `Compose.Seq` | 単一値入力から単一値出力までの合成を接続・評価 |
| `op.composedToSeq2.Pipe` | `Compose.Seq` | 組から単一値への操作を追加・評価 |
| `op.composedToSeq2.Pipe2` | `Compose.Seq2` | 組から組への操作を追加・評価 |
| `op.composedToSeq2.Operator` | `Compose.Seq2` | 単一値入力から組出力までの合成を接続・評価 |
| `op.composedToSeq.Pipe` | `Compose2.Seq` | 単一値から単一値への操作を追加・評価 |
| `op.composedToSeq.Pipe2` | `Compose2.Seq` | 単一値から組への操作を追加・評価 |
| `op.composedToSeq.Operator` | `Compose2.Seq` | 組入力から単一値出力までの合成を接続・評価 |
| `op.composed2.Pipe` | `Compose2.Seq` | 組から単一値への操作を追加・評価 |
| `op.composed2.Pipe2` | `Compose2.Seq2` | 組から組への操作を追加・評価 |
| `op.composed2.Operator` | `Compose2.Seq2` | 組入力から組出力までの合成を接続・評価 |

`Seq` / `Seq2` と `Operator` / `Operator2` / `ToSeqOperator` / `ToSeq2Operator` は関数型で、独立した関数・メソッドを追加しない。  
それぞれ入力・Map / Map2・MapToSeq2 / MapToSeq と Pipe / Pipe2 の測定で利用する。  
合成値の各メソッドは単独の費用を分離せず、4 種類の入出力を扱う合成全体で測る。  
4 種類とも 4 段を合成し、構築・入力への接続・全件反復を含める。  
対照の Pipe は、同じオペレーターを同じ順序で `Seq` / `Seq2` の Pipe / Pipe2 に直接接続する。  
FromRange は符号付き・符号なしの代表として int / uint を選び、1 から N までと N から 1 までを stop を含めて生成する。  
`ZipAlign` と `AlignShortest` / `AlignLeft` / `AlignRight` / `AlignLongest` は型・定数であり、対応する Zip ケースに含める。  
各ケースの詳しい条件は [測定方法](../../benchmarks/METHOD.md) を参照。  

## 測定コードの配置

| API・機能 | 測定・結果検証のファイル |
| --- | --- |
| Map / Map2 | [map_test.go](../../benchmarks/map_test.go) |
| Filter / Filter2 | [filter_test.go](../../benchmarks/filter_test.go) |
| Keys / Values | [keys_test.go](../../benchmarks/keys_test.go)、[values_test.go](../../benchmarks/values_test.go) |
| Take / Skip と Operator / Operator2 | [take_test.go](../../benchmarks/take_test.go)、[skip_test.go](../../benchmarks/skip_test.go) |
| FromRange | [from_range_test.go](../../benchmarks/from_range_test.go) |
| FromSlice / ToSlice | [from_slice_test.go](../../benchmarks/from_slice_test.go)、[to_slice_test.go](../../benchmarks/to_slice_test.go) |
| FromSeq / FromSeq2 | [from_seq_test.go](../../benchmarks/from_seq_test.go) |
| ToSeq / ToSeq2 | [to_seq_test.go](../../benchmarks/to_seq_test.go) |
| FromMap / ToMap | [from_map_test.go](../../benchmarks/from_map_test.go)、[to_map_test.go](../../benchmarks/to_map_test.go) |
| FromChannel / ToChannel | [from_channel_test.go](../../benchmarks/from_channel_test.go)、[to_channel_test.go](../../benchmarks/to_channel_test.go) |
| Zip と設定メソッド | [zip_test.go](../../benchmarks/zip_test.go) |
| Seq / Seq2 の Pipe / Pipe2、MapToSeq / MapToSeq2 | [pipe_test.go](../../benchmarks/pipe_test.go) |
| Compose / Compose2 と合成値の全メソッド | [compose_test.go](../../benchmarks/compose_test.go) |
| Cache / Cache2、Fill / Warm / Reuse | [cache_scenario_test.go](../../benchmarks/cache_scenario_test.go) |
| Channel を含むパイプライン | [channel_scenario_test.go](../../benchmarks/channel_scenario_test.go) |
| Map / Filter / Take を接続するパイプライン | [pipeline_scenario_test.go](../../benchmarks/pipeline_scenario_test.go) |

[benchmark_helper_test.go](../../benchmarks/benchmark_helper_test.go) は入力データの生成、結果保持用の変数、要素の変換・判定だけを共有する。  
Unit のケースは機能別の `*_test.go`、Scenario のケースは機能別の `*_scenario_test.go` に分離する。  
共有の送信処理は `channel_helper_test.go` に置き、異なる測定分類のケースを同じファイルへ混在させない。  
処理関数を格納するケースの登録表、共通のベンチマークループ、テストのディスパッチャーは設けない。  

## 処理と入口

各 API・実装の処理全体を `mapLibrary` / `mapGo` のような型付き関数に置く。  
機能ごとに `mapLibrary` → `mapGo` → `BenchmarkMap` → `TestMap` の順に配置する。  
`BenchmarkMap` / `TestMap` は比較する実装を明示的な `b.Run` / `t.Run` に分け、それぞれの処理関数を直接呼ぶ。  
Seq の生成・演算子の構築・評価・収集・結果の加算は処理本体に含め、入力データだけを呼び出し元で準備する。  
Cache / Cache2 は Scenario のみで扱い、`scenarioCacheSeqFill` / `scenarioCacheSeqWarm` のように段階を分ける。  
Scenario は `scenarioPipelineSeqLibrary` / `scenarioPipelineSeq2Library` のように、種類と実装ごとに分ける。  

Cache の Fill / Warm はそれぞれ処理・ベンチマーク・テストをまとめ、Reuse は Cached / Uncached を `BenchmarkScenarioCacheSeqReuse` / `TestScenarioCacheSeqReuse` へまとめる。  
Seq2 も同じ構成とし、Pipeline は Seq / Seq2 ごと、Channel は 5 実装を一つの機能としてまとめる。  
各 `b.Run` / `t.Run` に実装だけを囲む中間の `Run` を設けず、`t.Run` は親関数の直下へ個別に記述する。  
ベンチマークの入力条件は表で列挙してよいが、ケース名・順序・測定範囲を保ち、実装ごとに測定ループを記述する。  
ベンチマークの各ケース名には実装と条件を明記し、`b.Loop()` 内から処理本体を直接呼ぶ。  
`class=Unit` / `class=ScenarioPipeline` / `class=ScenarioChannel` / `class=ScenarioCache` で分類し、入力や段数などの比較条件も名前へ残す。  
テストは各処理につき代表入力 1 件で処理本体を直接呼び、簡単な期待値で疎通を確認する。  
公開表のリンクは測定コミットの関数宣言行から自動生成するため、行番号を手動で転記しない。  
配置・入口名の整理だけでは再計測せず、公開済みの測定値・日時・対象コミットと固定リンクを保つ。  
