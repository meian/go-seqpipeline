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

## 更新 PR の CI

`.github/workflows/ci.yml` はパスや作成者で除外せず、`pull_request` で gofmt・build・test・Static checks・README policy を実行する。  
[Dependabot が起動する実行](https://docs.github.com/en/code-security/reference/supply-chain-security/troubleshoot-dependabot/dependabot-on-actions)では通常の Actions secrets を使えず、既定の `GITHUB_TOKEN` も読み取り専用になる。  
既存の lint ジョブは `dependabot[bot]` の場合に `github-pr-annotations` を使い、レビューコメントの書き込みに依存しない。  
他の CI ジョブも公開依存の取得と検証だけを行い、追加の secrets は必要ない。  

1. `main` への設定反映後、GitHub の Insights → Dependency graph → Dependabot で両エコシステムの更新ログとエラーを確認する
2. 更新 PR では、同じ Action を参照する全 workflow、SHA と版番号コメント、関連する文書の同期を確認する
3. 自動マージの対象外を手動でマージするときも、最新 head の CI と保護設定を確認する
4. PR が作成されていない週も、更新不要なのか、件数上限や更新エラーで止まっているのかをログで確認する

Dependabot の更新検出と実際の更新 PR による CI の確認には、設定が既定ブランチに存在する必要がある。  
導入 PR の通常 CI の成功だけで、Dependabot による CI・rebase・自動マージまで確認済みとは扱わない。  

## 自動マージの対象

[Dependabot auto-merge](../../.github/workflows/dependabot-automerge.yml) が、同一リポジトリの Dependabot 作成・`main` 向け PR を判定する。  
署名を検証できる Dependabot コミットの `updated-dependencies` メタデータと実差分を照合し、全依存が minor / patch 更新の場合だけ対象にする。  
タイトルやラベルだけでは判定しない。  

許可する差分は、既存の Action の `uses:` 参照と `ci/scripts/requirements.txt` の既存の固定版の置換に限る。  
SHA 固定の Action は、変更前後の版番号コメントと、そのタグが指す SHA も確認する。  
グループ PR は、全依存のメタデータと差分が一致する必要がある。  
major・downgrade・prerelease・分類不能な更新や、許可した行以外の変更は手動確認に回す。  
`go.mod` の `go` / `toolchain`、CI の `go-version` などの Go バージョン更新は、patch でもグループに含まれていても対象外になる。  

## main への追随とマージ条件

PR が `main` に遅れている場合は `@dependabot rebase` を依頼し、新しい head の CI が完了するまで保留する。  
head と `main` の組をコメントのマーカーに記録し、同じ状態への依頼を重複させない。  
競合・rebase 失敗・API 取得失敗も保留し、古い head の成功結果を使わない。  

必須の gofmt・build・test・Static checks・README policy と、`main` に適用中の branch rules から取得した追加必須チェックの成功を確認する。  
必須チェックの未作成・実行中・失敗・キャンセルに加え、任意チェックの失敗や実行中も保留する。  
`push` 側の Static checks の skip だけは、同じ head の `pull_request` 側が成功していれば許可する。  
自動マージ処理自身は待機対象から外す。  

マージ直前に head と `main` を再取得し、[非同期マージ API](https://docs.github.com/en/rest/pulls/pulls?apiVersion=2026-03-10#merge-a-pull-request-asynchronously) に head の `sha`、`merge_method: squash`、`merge_action: direct_merge` を指定する。  
`GITHUB_TOKEN` で `bypass_rules: false` を明示し、strict な必須チェック、署名、未解決スレッドなどの既存の保護設定を守る。  
応答が `pending` なら結果を 1 回確認し、未完了の場合は UUID と完了未確認をサマリに記録する。  
`merged` の応答だけをマージ完了として扱う。  
GitHub の auto-merge 設定は有効化せず、条件がそろった実行で直接マージする。  

## 自動化の権限と再実行

自動化は `CI` の完了、`main` への push、15 分ごとの定期実行、手動実行で開いている対象 PR を再判定する。  
`main` のコードだけを checkout し、PR のコードや依存は実行しない。  
`GITHUB_TOKEN` に `contents: write`・`pull-requests: write`・`checks: read`・`statuses: read`・`actions: read` を与え、追加の secret は使わない。  
権限とイベントの制約は [GitHub 公式の Dependabot 自動化](https://docs.github.com/en/code-security/tutorials/secure-your-dependencies/automate-dependabot-with-actions)も参照する。  

Actions の `Dependabot auto-merge` の実行サマリで、PR ごとの判定と保留理由を確認する。  
手動実行は `dry-run: true` が既定で、rebase の依頼やマージをせずに判定する。  
`pr-number` を省略すると全件、指定するとその PR だけを再判定する。  

```sh
# 全件の判定だけ確認する
gh workflow run 'Dependabot auto-merge' --ref main -f dry-run=true

# 指定した PR の再判定と、条件を満たす場合の rebase 依頼・マージを行う
gh workflow run 'Dependabot auto-merge' --ref main -f pr-number=123 -f dry-run=false

# サマリに残った UUID で非同期マージの結果を確認する
gh api repos/meian/go-seqpipeline/pulls/PR_NUMBER/merge-async/UUID
```

CI 失敗はログを確認して原因を解消し、必要に応じて `gh run rerun RUN_ID --failed` で再実行する。  
rebase 依頼に反応がない場合は、PR の状態と Dependabot のログを確認し、必要に応じて `@dependabot rebase` を手動でコメントする。  
分類不能な差分や保護設定による保留は、内容を手動で確認し、自動判定の条件を緩めずに対応する。  
