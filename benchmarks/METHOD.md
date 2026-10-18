# 測定方法

[今回の環境・コミット・コマンド条件](results/environment.md) と [公開 API の対応表](API.md) を併せて参照。  
全測定値は [Unit 詳細表](results/unit.md) と [Scenario 詳細表](results/scenario.md) に掲載する。  

## コードの配置

測定コードは [変換・収集](conversion_test.go)、[変換・選別・射影](transform_test.go)、[件数制御](limit_test.go)、[Zip](zip_test.go)、[Pipe 接続](pipe_test.go)、[Pipeline](pipeline_test.go)、[Channel](channel_test.go)、[Cache](cache_test.go) に分ける。  
同じ機能の Unit / Scenario と結果・等価性テストを同じファイルに置く。  
[共通ヘルパー](benchmark_helper_test.go) が入力・出力保持・集計と Unit の実行を担当し、測定分類はファイル名ではなく既存のベンチマーク名で識別する。  

## Unit の範囲

名前は `Unit/api=Map/N=64/phase=Evaluate` のように API・入力件数・評価範囲を表す。  
整数の入力を使い、N は 64 / 4096、入力スライス・map・標準 iterator と再評価可能な Seq / Seq2 の準備は測定外とする。  
各ケースの `run` 呼び出し、対象 API による構築・接続、評価、int64 の加算を測定する。  
共通の関数呼び出しと加算の費用も含む指標であり、API の内部処理だけを分離した時間ではない。  
出力チェックサムをグローバル変数へ保存し、ToSlice / ToMap は収集結果自体もグローバル変数へ保持する。  

| phase | 範囲 |
| --- | --- |
| Evaluate | API の呼び出し・構築・全件評価・チェックサム |
| Collect | ToSlice / ToMap の全件収集・割り当て・結果保持・チェックサム |
| Early8 | Take の構築・先頭 8 件までの評価・チェックサム |
| Fill | Cache / Cache2 の呼び出し・全件保存・結果保持。保存後の反復を含めない |
| Warm | 計測前に保存を完了した結果の全件反復・チェックサム |

Map は `(v + 1) & 65535`、Filter は `v % 3 != 0`、対応する Seq2 はキーを維持する。  
Skip は先頭 8 件を除外し、Keys / Values は添字 / 値を抽出する。  
Zip.Operator / ZipKey.Operator は等長入力を短い側に合わせ、他の Zip 設定ケースは片側を N/2 件にして不足分を実際に通す。  
AlignBy は Left / Right、WithDefaultValue は値 7（ZipKey 側は 9）、WithDefaultKey はキー 7（再設定は 9）を使う。  
ZipKey.WithDefaultValue / WithDefaultKey は Longest を使い、設定ごとの費用は接続・評価と合わせた指標として扱う。  

Seq.Pipe は Map、Seq.Pipe2 は Zip、Seq2.Pipe は Values、Seq2.Pipe2 は Map2 と合わせて測る。  
これらの転送メソッドや型変換・設定値だけの極小測定を API の評価性能の代表にはしない。  
FromChannel は容量 0 の入力 channel・送信用 1 goroutine を毎回作り、全受信と送信者の終了待ちを含む。  
ToChannel は API 内で容量 0 の channel・送信用 1 goroutine を毎回作り、全受信して close を観測する。  
どちらも Map / Filter は追加せず、生成・起動・送受信・終了・加算を含む。  

## Scenario の名前

- `ScenarioPipeline/kind=Seq/N=64/depth=2/mode=All/impl=Loop`
- `ScenarioChannel/N=64/buffer=0/impl=Direct`
- `ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=4/impl=Cached`

`impl` を benchstat の列へ投影し、他の軸が同じ行の実装だけを比較する。  
Fill / Warm は独立した指標として示し、構築範囲の異なる Warm / Uncached の比率は計算しない。  

## 比較条件

### ScenarioPipeline

入力は `0..N-1` の `[]int`、N は 64 / 4096。  
`Seq` は値のみ、`Seq2` はスライスの添字と値を扱い、添字を維持する。  
D2 は Map 1 回と Filter、D6 は Map 5 回と Filter で、Map は `(v + 1) & 65535`、Filter は `v % 3 != 0`。  
All は全件評価、Early8 は Filter 後の 8 件で停止し、パイプラインでは追加の Take 1 段を使う。  
したがって D は Take を除くオペレーター数を表す。  

| 実装 | 処理 |
| --- | --- |
| Loop | 通常のループ内で変換・判定を直接記述 |
| Functions | 要素ごとの変換・判定を `mapValue` / `keepValue` に分離 |
| Pipeline | Seq / Seq2 を構築し、Map / Filter / Take を接続して評価 |

全実装で同じ `[]pair{key, value}` に、nil スライスから append で出力を収集する。  
Seq の key は 0 とし、比較対象も同じ形式にそろえる。  
入力生成は計測外、構築・評価・出力割り当ては計測内で、結果はグローバル変数に保持する。  
構築だけの測定は行わない。  
Functions は処理段階ごとにスライスを実体化する方式ではなく、要素ごとの関数呼び出しであり、インライン化される可能性がある。  

### ScenarioChannel

N は 64 / 4096、処理は D2 の Map / Filter と同じで、全件を同じ append 条件で収集する。  
途中終了やキャンセルはこの測定の対象外。  
コンテキストは `context.Background()`、ToChannel の出力容量は API の契約どおり常に 0。  

| ケース | 構成 | 対照 |
| --- | --- | --- |
| Seq | 呼び出し元で Seq を直接反復 | — |
| ToChannel | Seq → ToChannel → 受信。送信用 1 goroutine | Seq |
| Direct | 入力送信用 1 goroutine → channel → 呼び出し元で変換・判定 | — |
| FromChannel | 同じ送信者 → channel → FromChannel → Map / Filter | 同じ容量の Direct |
| RoundTrip | 同じ送信者 → channel → FromChannel → Map / Filter → ToChannel → 受信。計 2 goroutine | 同じ入力容量の Direct |

Direct / FromChannel / RoundTrip は入力 channel の容量 B を 0 / 64 に変える。  
入力スライス生成を除き、channel 生成、goroutine 起動、全送受信、close の観測、送信者の終了待ち、変換・評価・出力収集を計測に含める。  
追加の変換に必要な channel・goroutine 自体も比較するコストに含める。  
出力を読み切り、入力送信者の終了通知を待つため、次の反復に送信待ちの処理を持ち越さない。  

### ScenarioCache

N は 64 / 4096、Seq / Seq2 の Map で各値に `v = (v * 1664525 + 1013904223) & 65535` を W 回適用する。  
W0 は軽い上流、W32 は算術負荷のある上流を表し、Seq2 のキーは元の添字を維持する。  
出力は全件を int64 のチェックサムに加算し、グローバル変数に保存する。  
比較対象には同じ評価・加算条件を使い、下流の出力スライスは作らない。  

| ケース | 測定範囲 |
| --- | --- |
| Fill | 上流構築 + Cache / Cache2 の作成・全件保存。保存結果を保持し、下流の反復は行わない |
| Warm | 計測前に全件保存済み。計測内では保存結果を 1 回反復して加算 |
| R1 / R4 / R16 の Uncached | 上流構築 + キャッシュなしで R 回評価・加算 |
| R1 / R4 / R16 の Cached | 上流構築 + 初回全件保存 + 保存結果を R 回反復・加算 |

Cache / Cache2 は呼び出し時に全件評価するため、Fill が初回保存に相当する。  
R は下流が結果を利用する回数で、Cached の保存時には上流をさらに 1 回評価する。  
Cached / Fill は測定反復ごとに新しいキャッシュを作り、Warm だけが事前作成したキャッシュを使う。  
Warm と Uncached は計測範囲が異なるため比率を計算せず、初回費用を含む比較には同じ R の Cached / Uncached を使う。  

| N | 保持する要素数・組数 | Seq の値データ | Seq2 のキー・値データ |
| ---: | ---: | ---: | ---: |
| 64 | 64 | 512 bytes | 1024 bytes |
| 4096 | 4096 | 32768 bytes | 65536 bytes |

arm64 の int は 8 bytes として計算した論理データ量で、スライスの余剰容量やヘッダー、クロージャーは含まない。  
B/op は途中の拡張時に破棄する領域も含む割り当て総量であり、保持データ量とは異なる。  
初回保存には要素数に比例する保持メモリが必要で、Warm の B/op に保持済みの領域は含まれない。  

## 出力の等価性

`TestPipelineEquivalent` は空入力・1 件・64 件・4096 件について、全条件の出力値・キー・順序をスライス比較する。  
`TestChannelEquivalent` は同じ入力群で、全構成・容量の出力値・順序を確認する。  
`TestCacheEquivalent` は同じ入力群・上流負荷でキャッシュ前後の値・キー・順序を比較し、16 回の再利用でも期待チェックサムが一致することを確認する。  
各機能ファイルの `Test*UnitResults` は空入力・1 件・64 件・4096 件で各 API/phase の期待チェックサムを確認し、Cache の Fill では保存した結果も計測外で確認する。  
これらは時間計測と分離した通常のテストとして実行する。  

## 集計と再作成

[benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) が中央値・中央値の 95% 信頼区間・比較検定を計算する。  
時間の差の検定には既定の有意水準 0.05 を使い、`~` は統計的に差が明確ではないことを表す。  
概要の倍率は同じ集計行の中央値同士の比であり、検定結果も併記する。  
ゼロ割り当てを含む全件の幾何平均は解釈せず、各ケースの値を参照する。  
`publish.py` は benchstat CSV の中央値・信頼区間・検定結果を読み、全件の Markdown 表と概要を生成する。  
Unit は API と測定条件を分けてカテゴリ別に並べ、Scenario は条件が一致する対照と対象の時間を横並びにする。  
横幅を抑えるため Scenario のメモリ割り当ては別表とし、B/op と allocs/op の検定も掲載する。  
比較対象のないケースも単独測定表へ自動的に掲載する。  
新しい API は既知のカテゴリまたは「その他」、新しい Scenario はその名前の見出しへ追加される。  
公開表で未使用のケース・指標、指標の欠落・重複、比較 CSV と全件 CSV の中央値・信頼区間の不一致があれば生成を失敗させる。  
代表例の選択だけは `publish.py` で指定し、定期更新時に見直す。  
統計計算は benchstat に任せ、整形処理では倍率と表示用の単位変換のみを計算する。  

```sh
# 初回のみ。測定ツールをライブラリの go.mod に追加しない
go install golang.org/x/perf/cmd/benchstat@v0.0.0-20261009192801-be2c69fb417e
# 正しさの検証
go test ./... -count=1
python3 -m unittest discover -s benchmarks -p 'test_*.py'
# 現在の結果と概要を全件置き換える
BENCHSTAT="$(go env GOPATH)/bin/benchstat" BENCHTIME=100ms COUNT=10 GOMAXPROCS=2 sh benchmarks/measure.sh
```

Go と benchstat の実行ファイルは `GO_BINARY` / `BENCHSTAT` で指定できる。  
サンドボックスでは事前取得済みの Go ツールチェーンの実行ファイルを `GO_BINARY` に指定し、必要なら `GOTOOLCHAIN=local` で自動ダウンロードを避ける。  
`GOCACHE` を省略すると `/tmp` の一時ディレクトリ内に作り、測定終了時に削除する。  
再利用できる書き込み先がある場合は `GOCACHE` を明示してよいが、raw の出力先は変わらない。  
raw と benchstat CSV・環境記録用の中間ファイルは `/tmp/seq-benchmark.XXXXXX/` にのみ出力する。  
全件の生成と整合性検証に成功してから公開 Markdown を置き換え、成功・失敗・通常の中断時には trap で一時ディレクトリ全体を削除する。  
Git 管理する測定結果は `results/unit.md`、`results/scenario.md`、`results/environment.md` のみとし、raw・TXT・CSV は保存しない。  
SIGKILL や OS 停止では trap が動かないため、その場合は該当する一時ディレクトリを削除してから再実行する。  
集計結果の再生成には再測定が必要で、生データの保管・世代管理は行わない。  
詳細な運用とチェックリストは [運用ガイド](../doc/ai/benchmark.md) を参照。  
