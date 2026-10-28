# ローカルでの lint

[Issue #90](https://github.com/meian/go-seqpipeline/issues/90) の導入前確認として、全 Go パッケージを解析する。  
CI への組み込みと PR へのレビュー投稿は別途行う。  

## ツールと導入

2026-10-29 時点の最新安定版を `ci/scripts/lint.sh` で固定する。  

| ツール | バージョン |
| --- | --- |
| golangci-lint | [v2.15.0](https://github.com/golangci/golangci-lint/releases/tag/v2.15.0) |
| reviewdog | [v0.21.2](https://github.com/reviewdog/reviewdog/releases/tag/v0.21.2) |

`go.mod` の要件を満たす Go、Bash、curl、tar を使い、リポジトリのルートで実行する。  
タグを固定した公式インストーラーから配布バイナリを `bin/lint/` 以下のバージョン別ディレクトリへ導入する。  
golangci-lint の各 linter は、この配布バイナリに同梱されたバージョンを使う。  
複雑性の解析には同梱の gocyclo v0.6.0 と gocognit v1.3.0 を使う。  
セキュリティの解析には同梱の gosec v2.29.0 を使う。  
gosec 単体の最新安定版は [v2.30.0](https://github.com/securego/gosec/releases/tag/v2.30.0) だが、内部の linter は golangci-lint が組み込んだ版に固定される。  
golangci-lint が推奨する[バイナリ導入](https://golangci-lint.run/docs/welcome/install/local/)を使い、ライブラリの `go.mod` にツール依存を追加しない。  

```sh
bash ci/scripts/lint.sh install
bash ci/scripts/lint.sh
```

## 解析対象と重大度

`.golangci.yml` に採用する linter を明示し、`golangci-lint run ./...` でテスト・Example・ベンチマークを含む全 Go パッケージを解析する。  
差分フィルター・指摘件数の上限・同一行での重複抑制を使わず、既存コードの指摘も表示する。  
テストでの意図的なエラー無視を許容するため、`errcheck` の指摘対象から `_test.go` を除外する。  
nil context による panic の検証を許容するため、`staticcheck` の SA1012 も `_test.go` では除外する。  
これらの除外はテスト・Example・ベンチマークに適用する。  

| linter / 指摘 | 重大度 |
| --- | --- |
| govet | Error |
| gosec | Error |
| errcheck | Error |
| ineffassign | Error |
| unused | Error |
| staticcheck の SA 系 | Error |
| staticcheck の S / ST / QF 系 | Warning |
| gocyclo | Warning |
| gocognit | Warning |

[gocyclo](https://github.com/fzipp/gocyclo) で分岐・ループなどによる循環的複雑度、[gocognit](https://github.com/uudashr/gocognit) でネストなども加味する認知的複雑度を確認する。  
初期の確認基準として循環的複雑度 15 超、認知的複雑度 20 超を報告し、テスト・ベンチマークも対象に含める。  
値だけで修正の要否を決めず、関数リテラル内の処理やテストケースの反復も含む集計結果として確認する。  

重大度は `.golangci.yml` の `severity` で設定する。  
既定を Error とし、種類ごとの上書きには指摘文のチェック ID を使う。  
linter 全体を変更する場合は `linters` を条件とするルールを追加する。  

解析結果を `bin/lint/report.xml` に Checkstyle 形式で保存し、linter 名・位置・重大度を reviewdog に渡す。  
表示した指摘は `bin/lint/report.txt` にも保存する。  
reviewdog の `local` reporter と `nofilter` を使うため、GitHub のトークンは不要。  
golangci-lint の終了コード 1 は指摘ありとして扱い、reviewdog の `-fail-level=error` で重大度による成否を判定する。  
Error があれば失敗し、Warning のみ・指摘なしなら成功する。  
ツールの起動・設定検証・解析・結果の読み取りに失敗した場合も、スクリプトは失敗する。  
生成物とツールは既存の `/bin/` の除外対象となる。  

## 初回スキャン結果

2026-10-29 に `origin/main` の `d2e26f4` を基準とし、Go 1.27.3（darwin/arm64）で実行した。  
ルート・`op`・`seq`・`benchmarks` の全 4 パッケージが対象。  

| チェック | 重大度 | 指摘数 |
| --- | --- | --- |
| G115 | Error | 9 |
| ST1023 | Warning | 27 |
| QF1011 | Warning | 27 |
| gocyclo | Warning | 6 |
| gocognit | Warning | 11 |

SA1012 は `seq/channel_test.go` の `TestFromChannelNilContext` と `TestToChannelNilContext` に対する 2 件をテスト除外で抑制した。  
ST1023 と QF1011 は同じ 27 箇所への型明示の省略提案で、API の接続型を明示するテストや Example も含む。  
gosec の G115 は整数変換に対する指摘で、`seq/range.go` に 2 件、`seq/range_test.go` に 2 件、`benchmarks/from_range_test.go` に 5 件ある。  
実装の 2 件は `uint64(step)` と `uint64(-(step + 1))` に対する指摘で、既存のビット列の加算と減算幅の計算に使われている。  
今回はこれらのソース修正や追加の除外を行わず、Error 9 件・Warning 71 件を残した。  
Error 9 件により、ローカル実行の終了コードは 1 となる。  

複雑性の基準を超えた 11 関数の値を以下に示す。  
循環的複雑度の「—」は報告基準以下を表す。  

| 関数 | 循環的複雑度 | 認知的複雑度 |
| --- | --- | --- |
| [zipSequences](../../op/zip.go) | 20 | 46 |
| [TestZipEmpty](../../op/zip_test.go) | — | 21 |
| [TestZipEvaluation](../../op/zip_test.go) | — | 29 |
| [TestToChannelCancelSend](../../seq/channel_test.go) | — | 22 |
| [TestComposeZipRelease](../../seq/compose_integration_test.go) | — | 32 |
| [TestCrossTypePipeline](../../seq/pipeline_test.go) | — | 28 |
| [BenchmarkScenarioCacheSeqReuse](../../benchmarks/cache_scenario_test.go) | 25 | 48 |
| [BenchmarkScenarioCacheSeq2Reuse](../../benchmarks/cache_scenario_test.go) | 25 | 48 |
| [BenchmarkScenarioChannel](../../benchmarks/channel_scenario_test.go) | 17 | 32 |
| [BenchmarkScenarioPipelineSeq](../../benchmarks/pipeline_scenario_test.go) | 25 | 48 |
| [BenchmarkScenarioPipelineSeq2](../../benchmarks/pipeline_scenario_test.go) | 25 | 48 |
