# 測定環境

測定日時は 2026-10-18T08:05:03+00:00、対象コミットは `d8ec4d818835a3ea6c67006b759ed17cbff2ba76`。  
対象コードと測定コードをこのコミットに固定し、文書・集計結果だけを後続コミットで公開する。  

| 項目 | 値 |
| --- | --- |
| Go | go version go1.27.1 darwin/arm64 |
| OS | macOS-26.6.2-arm64-arm-64bit |
| CPU | Apple M4 |
| GOMAXPROCS | 2 |
| benchtime / count | 100ms / 10 |
| ケース数 | 200 |
| benchstat | mod	golang.org/x/perf	v0.0.0-20261009192801-be2c69fb417e	h1:IB9P+pA897fozv0n9ldn+jTDNWtEMfbomaXXa4dKr74= |
| GOOS | darwin |
| GOARCH | arm64 |
| GOARM64 | v8.0 |
| GOAMD64 | 既定値 |
| GOGC | 既定値 |
| GOMEMLIMIT | 既定値 |
| GOEXPERIMENT | 既定値 |

通常のコンパイラ最適化・インライン化を有効にし、CPU 固定・温度制御・測定順のランダム化は行っていない。  
測定中には別のテストやビルドを実行せず、OS や他プロセスの負荷は制御していない。  
Go のキャッシュと測定の一時ファイルには書き込み可能な一時ディレクトリを使用した。  
信頼区間は benchstat の既定の 95%、同一スナップショット内の比較検定の有意水準は 0.05。  
raw は集計時の一時入力のみで、集計後に削除する。  
