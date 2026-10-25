# README と利用ガイドの更新

README は導入と最初の実行の入口とし、掲載範囲とレビュー基準は [AGENTS.md](../../AGENTS.md) に従う。  
利用者向けの詳細は [英語ガイド](../guide.md)・[日本語ガイド](../guide.ja.md)、API ごとの契約は GoDoc、実行可能な例は Go Example に置く。  

## CI ガード

`ci/check_readme.py` は日英それぞれの README を最大 100 行・コード例 2 つまでに制限する。  
整理後は両言語とも 73 行のため、暫定値 180 行から、説明の調整余地を残した 100 行に確定した。  
空行・表・コードブロックを含む全行を数え、コード例は `markdown-it-py` で解析し、言語指定の有無によらず引用・リスト内のフェンスやインデント形式のコードブロックも数える。  
単独の `go get github.com/meian/go-seqpipeline` を含む `sh` / `shell` / `bash` ブロックだけを例数から除外する。  
閉じていないフェンス、読み取り失敗、上限超過では失敗し、詳細を利用ガイド・GoDoc・Go Example へ移す案内を出す。  

リポジトリのルートで Python 3.8 以降を使い、文書チェック用の仮想環境に固定バージョンの依存を導入して実行する。  

```sh
python3 -m venv /tmp/seq-readme-venv
. /tmp/seq-readme-venv/bin/activate
python3 -m pip install -r ci/requirements.txt
python3 -m unittest discover -s ci -p 'test_*.py'
python3 ci/check_readme.py
```

ガードは内容の適切さや説明の重複、日英の意味の一致を判定しないため、レビュー基準と併用する。  
上限変更時は既存内容の整理と掲載理由を確認し、スクリプトの定数・本書・AGENTS.md を同期する。  
空行の削除や一行への詰め込みで上限を回避しない。  

## 更新時の確認

- README と利用ガイドの各日英版を同じ変更で同期する
- 移動前の条件・例外が残り、README から詳細へ辿れることを確認する
- 相対リンク・見出し・Example の参照先を確認し、コード例を現在の公開 API で実行する
- README に API 一覧を追加せず、新しい例は Go Example、共通仕様は利用ガイドへ追加する
- 正常・上限ちょうど・超過・例数増加などのガードテストと実ファイルのチェックを実行する
