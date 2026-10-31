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
PR の head ブランチへの追加 push は、Dependabot の rebase 後も含めて `pull_request` の `synchronize` で CI を実行する。  
`push` の対象は `main` に限定し、同じ PR の head に push 起因の CI を重複させない。  
`main` 以外で PR の head でもないブランチへの push だけでは、この CI は実行しない。  
[Dependabot が起動する実行](https://docs.github.com/en/code-security/reference/supply-chain-security/troubleshoot-dependabot/dependabot-on-actions)では通常の Actions secrets を使えず、既定の `GITHUB_TOKEN` も読み取り専用になる。  
既存の lint ジョブは `dependabot[bot]` の場合に `github-pr-annotations` を使い、レビューコメントの書き込みに依存しない。  
他の CI ジョブも公開依存の取得と検証だけを行い、追加の secrets は必要ない。  

1. `main` への設定反映後、GitHub の Insights → Dependency graph → Dependabot で両エコシステムの更新ログとエラーを確認する
2. 更新 PR では、同じ Action を参照する全 workflow、SHA と版番号コメント、関連する文書の同期を確認する
3. 自動マージの対象外を手動でマージするときも、最新 head の CI と保護設定を確認する
4. PR が作成されていない週も、更新不要なのか、件数上限や更新エラーで止まっているのかをログで確認する

## 自動マージ

[Dependabot auto-merge](../../.github/workflows/dependabot-automerge.yml) は、同一リポジトリの Dependabot 作成・`main` 向けの単一コミット PR を `pull_request_target` で判定する。  
PR の作成・更新イベントの実行者も `dependabot[bot]` に限定し、人が head を差し替えた実行は自動判定しない。  
`dependabot/fetch-metadata` でコミットの author と署名の検証結果を確認し、グループを含む全依存の更新種別・ecosystem・対象ディレクトリを matrix で確認する。  
minor / patch と分類された更新だけを対象とし、major・対象外・分類不能の依存が一つでもあればマージしない。  
実差分は公開 Action で確認し、既存 workflow の `uses:` 行と `ci/scripts/requirements.txt` の数値による固定版の置換だけを許可する。  
Go の最低要求バージョン・toolchain・CI の Go バージョンなど、依存参照以外の変更はグループに含まれる場合も対象外にする。  
独自スクリプトや workflow の `run:`・`script:` は追加しない。  

`wechuli/allcheckspassed` で head を固定して全チェックと commit status を待ち、未完了・失敗・キャンセル・skip・neutral は保留する。  
同じ head の `pull_request` CI の成功も確認し、REST API に検証した head の SHA を指定して squash merge する。  
マージ直前にも PR の head・マージ先・draft 状態・全チェックの集約結果・マージ可否を再確認する。  
最新 `main` への追随、必須チェック、署名、未解決スレッドなどは既存の ruleset で強制し、管理者 bypass は使わない。  

`main` への push と毎週月曜 11:00・21:00（Asia/Tokyo）の定期実行では、`main` に遅れた Dependabot PR へ `@dependabot rebase` を依頼する。  
対象判定を通過した PR は、マージが見送られた場合や API に拒否された場合も、同じ workflow 実行で遅れた PR を再確認する。  
PR の列挙は公式 Octokit Action、各 PR への rebase 依頼は matrix で行う。  
Dependabot が自身の権限で rebase と head ブランチへの push を行い、新しい head を対象に CI と自動マージ判定をやり直す。  
複数 PR は並行して判定するため、一つのマージで他の head が遅れると、残りの PR は再度 rebase と CI を行う。  
待機中の main 更新はマージ直前の再確認と strict な ruleset で検出し、古い main に基づく成功結果ではマージしない。  
PR のコードは checkout・実行せず、GitHub API からメタデータと差分だけを取得する。  

### rebase 依頼用の設定

リポジトリの Actions secret `DEPENDABOT_REBASE_TOKEN` に、リポジトリへの push 権限を持つユーザーの fine-grained PAT を登録する。  
対象リポジトリだけに限定し、Pull requests の Read and write を許可する。  
[Dependabot は GitHub App からのコマンドを受け付けない](https://github.com/dependabot/dependabot-core/issues/9147)ため、`GITHUB_TOKEN` では代用しない。  
この PAT は rebase 依頼のコメント投稿だけに使い、マージには通常の `GITHUB_TOKEN` を使う。  
rebase ジョブ自身はブランチを書き換えないため、`GITHUB_TOKEN` の `contents` 権限は `read` とする。  
Secret がない場合は rebase 依頼が失敗して保留となるため、マージ前に登録する。  

## 保留と再開

分類不能な差分、チェックの失敗・未完了、競合、rebase 失敗、API 取得失敗、保護条件の未達ではマージしない。  
手動で PR を再開した場合も自動マージは保留し、Dependabot による head 更新で再判定する。  
差分がない場合、変更ファイルまたは対象 PR が 100 件を超える場合、PR 検索結果が不完全な場合も保留する。  
Actions の `Dependabot auto-merge` の実行ログで保留理由を確認する。  
CI 失敗は原因を解消して `gh run rerun RUN_ID --failed` で再実行し、判定を再開するときは Actions の `Re-run jobs` を使う。  
rebase に反応がない場合は Dependabot のログを確認し、必要に応じて `@dependabot rebase` を手動でコメントする。  
対象外の更新は、自動判定の条件を緩めずに手動で確認する。  

Dependabot の更新検出と実際の更新 PR による CI の確認には、設定が既定ブランチに存在する必要がある。  
導入 PR の通常 CI の成功だけで、Dependabot による CI・rebase・自動マージまで確認済みとは扱わない。  
