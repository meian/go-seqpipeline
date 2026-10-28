# lint の運用

PR の CI とローカルで `.golangci.yml` を使って全 Go パッケージを解析する。  

## ツールと導入

ツールのバージョンを固定し、更新時は CI の入力とローカル導入コマンドを同期する。  

| ツール | バージョン |
| --- | --- |
| reviewdog/action-golangci-lint | [v2.10.2](https://github.com/reviewdog/action-golangci-lint/tree/v2.10.2) |
| golangci-lint | [v2.15.0](https://github.com/golangci/golangci-lint/releases/tag/v2.15.0) |
| reviewdog | [v0.21.2](https://github.com/reviewdog/reviewdog/releases/tag/v0.21.2) |

`go.mod` の要件を満たす Go、Bash、curl、tar を使い、リポジトリのルートで実行する。  
タグを固定した公式インストーラーから配布バイナリを `bin/` へ導入する。  
golangci-lint の各 linter は、この配布バイナリに同梱されたバージョンを使う。  
複雑性の解析には同梱の gocyclo v0.6.0 と gocognit v1.3.0 を使う。  
セキュリティの解析には同梱の gosec v2.29.0 を使う。  
golangci-lint が推奨する[バイナリ導入](https://golangci-lint.run/docs/welcome/install/local/)を使い、ライブラリの `go.mod` にツール依存を追加しない。  

```bash
set -o pipefail
curl -fsSL https://raw.githubusercontent.com/golangci/golangci-lint/v2.15.0/install.sh |
  sh -s -- -b ./bin v2.15.0
curl -fsSL https://raw.githubusercontent.com/reviewdog/reviewdog/v0.21.2/install.sh |
  sh -s -- -b ./bin v0.21.2
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
値だけで修正の要否を決めず、関数リテラル内の処理やテストケースの反復も含む集計結果として確認する。  

| 対象 | 循環的複雑度 | 認知的複雑度 |
| --- | --- | --- |
| 実装コード | 15超 | 20超 |
| `_test.go` | 30超 | 40超 |

パス別の閾値を直接設定できないため、共通閾値で解析し、テスト側の閾値以下の指摘を `exclusions.rules` の数値パターンで除く。  
テスト側の閾値変更時は、このパターンも同期する。  
指摘文の `> 15` / `> 20` は共通閾値を表す。  

重大度は `.golangci.yml` の `severity` で設定する。  
既定を Error とし、種類ごとの上書きには指摘文のチェック ID を使う。  
linter 全体を変更する場合は `linters` を条件とするルールを追加する。  

## PR の CI

`.github/workflows/ci.yml` の `Static checks` は `pull_request` で実行する。  
reviewdog Action の `github-pr-review` reporter と `added` フィルターにより、追加・変更行だけをレビューコメントと失敗判定の対象にする。  

golangci-lint の指摘は Checkstyle 形式で reviewdog へ渡し、linter・位置・Error / Warning を保持する。  
golangci-lint の終了コード 1 は指摘ありとして扱い、reviewdog の `fail_level: error` で差分内の Error だけを失敗とする。  
Warning のみ・指摘なし・差分外の指摘だけの場合は成功し、ツールの起動・解析・結果の読み取りが失敗した場合は失敗する。  

権限はこのジョブに `contents: read` と `pull-requests: write` を付与し、`GITHUB_TOKEN` で投稿する。  
fork PR など書き込み権限がない場合、reviewdog は Actions のログ注釈へフォールバックする。  

## ローカルでの全体確認

差分外も含めて確認する場合は、ルートで次を実行する。  
Checkstyle と表示結果は Git の追跡対象外の `bin/lint/` に保存する。  

```bash
bash <<'SH'
set -euo pipefail
mkdir -p bin/lint
bin/golangci-lint config verify
lint_status=0
bin/golangci-lint run ./... || lint_status=$?
case "$lint_status" in
  0|1) ;;
  *) exit "$lint_status" ;;
esac
bin/reviewdog -f=checkstyle -reporter=local \
  -filter-mode=nofilter -fail-level=error < bin/lint/report.xml |
  tee bin/lint/report.txt
SH
```
