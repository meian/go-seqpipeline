# 依存関係の更新

## 管理箇所と更新方法

[Dependabot 設定](../../.github/dependabot.yml)で Actions と CI 用 Python 依存の更新 PR を作成する。  
Go 本体と CI ツールは、毎週の依存更新レビュー時に公式リリースを確認して手動で更新する。  
緊急のセキュリティ更新は週次確認を待たずに対応する。  

| 対象 | 管理箇所 | 更新方法 |
| --- | --- | --- |
| GitHub Actions | `.github/workflows/*.yml` の `uses:` | Dependabot `github-actions` |
| CI 用 Python 依存 | `ci/scripts/requirements.txt` | Dependabot `pip` |
| Go の最低要求バージョン | `go.mod` の `go` 行 | 手動で採用判断 |
| CI の Go バージョン | `ci.yml` の `go-version` | 手動で系列を更新 |
| golangci-lint | `ci.yml` の `golangci_lint_version` | 手動で更新 |
| reviewdog | `ci.yml` の `reviewdog_version` | 手動で更新 |
| ローカル lint 導入手順 | `doc/ai/lint.md` | CI と同じ PR で同期 |

`github-actions` は Action 参照を更新し、`with:` の任意の入力やシェル内の版番号は更新しない。  
SHA 固定の Action は SHA 固定を維持し、同じ行の `# v...` コメントも更新差分で確認する。  
Dependabot には[版番号コメントを更新する実装](https://github.com/dependabot/dependabot-core/blob/main/github_actions/lib/dependabot/github_actions/file_updater/workflow_updater.rb)があるが、リリースタグと SHA の対応もレビューする。  

`gomod` は Go モジュールの依存更新用であり、Go 本体を定期更新する設定としては扱わない。  
現在は外部モジュール依存がないため設定せず、依存を追加するときに `/` を対象とする `gomod` の週次更新を追加する。  
Go toolchain の更新には[上流の機能要望](https://github.com/dependabot/dependabot-core/issues/13520)があり、`go` 行や CI の Go バージョンは手動確認の対象に残す。  
CI ツールの自動更新のために、ライブラリの `go.mod` へツール依存を追加しない。  

## 頻度と PR 数

両エコシステムとも毎週月曜 09:00（Asia/Tokyo）に確認する。  
minor・patch 更新はエコシステムごとにまとめ、major 更新は個別 PR にする。  
バージョン更新 PR は各エコシステムで同時に最大 2 件、合計最大 4 件とする。  
セキュリティ更新はこの週次スケジュールと件数上限の対象外であり、別に確認する。  
設定項目は [Dependabot の公式リファレンス](https://docs.github.com/en/code-security/reference/supply-chain-security/dependabot-options-reference)に従う。  

## Go と CI ツールの手動更新

1. [Go](https://go.dev/dl/)、[golangci-lint](https://github.com/golangci/golangci-lint/releases)、[reviewdog](https://github.com/reviewdog/reviewdog/releases) の公式リリースで更新内容と対応 Go バージョンを確認する
2. CI の Go は `go-version` の系列内の patch を使う指定とし、系列を変える場合に workflow を更新する
3. `go.mod` の最低要求バージョンは CI と別に判断し、変更時は [版番号決定ルール](release.md)に従い、README 日英の要件表記も同期する
4. CI ツールは workflow の固定版と、[lint の運用](lint.md)のバージョン表・インストーラー URL・導入コマンドの引数を同じ PR で更新する
5. lint Action の更新ではバージョン表のリンクも同期し、入力・出力形式・権限の変更を確認する
6. ツール更新後は lint の運用に従って全体を解析し、gosec の抑制範囲と同梱版の説明も確認する

## 更新 PR の CI と確認

`.github/workflows/ci.yml` はパスや作成者で除外せず、`pull_request` で gofmt・build・test・Static checks・README policy を実行する。  
[Dependabot が起動する実行](https://docs.github.com/en/code-security/reference/supply-chain-security/troubleshoot-dependabot/dependabot-on-actions)では通常の Actions secrets を使えず、既定の `GITHUB_TOKEN` も読み取り専用になる。  
既存の lint ジョブは `dependabot[bot]` の場合に `github-pr-annotations` を使い、レビューコメントの書き込みに依存しない。  
他の CI ジョブも公開依存の取得と検証だけを行い、追加の secrets は必要ない。  

1. `main` への設定反映後、GitHub の Insights → Dependency graph → Dependabot で両エコシステムの更新ログとエラーを確認する
2. 更新 PR では、同じ Action を参照する全 workflow、SHA と版番号コメント、関連する文書の同期を確認する
3. 自動マージの対象外を手動でマージするときも、最新 head の CI と保護設定を確認する
4. PR が作成されていない週も、更新不要なのか、件数上限や更新エラーで止まっているのかをログで確認する

## 自動マージ

[Dependabot auto-merge](../../.github/workflows/dependabot-automerge.yml) は、同一リポジトリの Dependabot 作成・`main` 向けの単一コミット PR を `pull_request_target` で判定する。  
`dependabot/fetch-metadata` の署名検証済みメタデータを使い、グループを含む全依存が minor / patch 更新で、実差分と一致する場合だけ対象にする。  
許可する差分は、既存の Action の `uses:` と `ci/scripts/requirements.txt` の固定版の置換に限る。  
SHA 固定では、変更前後の版番号コメントとリリースタグが指す SHA も照合する。  
major・downgrade・prerelease・分類不能な更新と、Go の最低要求バージョン・toolchain・CI の Go バージョンの変更は、グループに含まれる場合も対象外にする。  

`gh pr checks --required --watch --fail-fast` で必須チェックを待ち、公式 Octokit Action で最新 head の結果を取得する。  
Actions の matrix と条件式で任意も含めた全件を確認し、自動マージ workflow 自身を除いて未完了・失敗・キャンセルなどは保留する。  
`push` 側の Static checks の skip だけを許可し、同じ head の `pull_request` 側の CI 実行成功も必須にする。  
`gh pr merge --squash --match-head-commit` で確認した head を指定し、最新 `main` への追随要件や署名・未解決スレッドなどの保護設定を守ってマージする。  

`main` への push と毎週月曜 11:00・21:00（Asia/Tokyo）の定期実行では、`main` に遅れた Dependabot PR へ `@dependabot rebase` を依頼する。  
自動マージ成功後も、同じ workflow 実行で残りの PR を確認する。  
PR の列挙は公式 Octokit Action、各 PR の処理は matrix で行い、rebase 後の新しい head で判定と CI 確認をやり直す。  
実行するコードは `main` のものに限り、PR のコミットは差分確認用のデータとして取得する。  

## 保留と再開

分類不能な差分、チェックの失敗・未完了、競合、rebase 失敗、API 取得失敗、保護条件の未達ではマージしない。  
結果の取りこぼしを避けるため、チェックまたは対象 PR が 100 件を超える場合と、PR 検索結果が不完全な場合も保留する。  
Actions の `Dependabot auto-merge` の実行ログで保留理由を確認する。  
CI 失敗は原因を解消して `gh run rerun RUN_ID --failed` で再実行し、判定を再開するときは Actions の `Re-run jobs` を使う。  
rebase に反応がない場合は Dependabot のログを確認し、必要に応じて `@dependabot rebase` を手動でコメントする。  
対象外の更新は、自動判定の条件を緩めずに手動で確認する。  

Dependabot の更新検出と実際の更新 PR による CI の確認には、設定が既定ブランチに存在する必要がある。  
導入 PR の通常 CI の成功だけで、Dependabot による CI・rebase・自動マージまで確認済みとは扱わない。  
