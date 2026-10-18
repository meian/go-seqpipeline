# 性能概要

このレポートは測定時点の性能スナップショットである。  
[測定環境・対象コミット](results/environment.md)、[測定方法と範囲](METHOD.md)、[今回の API 対応表](API.md) を参照。  
全 Unit / Scenario の時間・割り当て量と信頼区間は [詳細集計](results/all.txt)（[CSV](results/all.csv)）に分離している。  

## Scenario の代表値

N4096 の代表例を示し、全件・途中終了、キャッシュの初回費用を含む利用回数を区別する。  
時間は benchstat の中央値、倍率は同じ行の対象 / 対照であり、過去の測定とは比較しない。  
倍率が異なっていても検定が `~` の場合は差を断定しない。  

| 条件 | 対照 → 対象 | 対照 ns/op | 対象 ns/op | 倍率 | 同一測定内の検定 |
| --- | --- | ---: | ---: | ---: | --- |
| Pipeline/kind=Seq/N=4096/depth=2/mode=All | Loop → Pipeline | 13349.00 | 25078.00 | 1.879 | +87.86% (p=0.000 n=10) |
| Pipeline/kind=Seq/N=4096/depth=6/mode=All | Functions → Pipeline | 18257.50 | 52655.00 | 2.884 | +188.40% (p=0.000 n=10) |
| Pipeline/kind=Seq2/N=4096/depth=6/mode=All | Loop → Pipeline | 18249.50 | 58021.50 | 3.179 | +217.93% (p=0.000 n=10) |
| Pipeline/kind=Seq/N=4096/depth=6/mode=Early8 | Loop → Pipeline | 73.62 | 369.10 | 5.013 | +401.32% (p=0.000 n=10) |
| Channel/N=4096/buffer=0 | Direct → FromChannel | 374654.50 | 469948.00 | 1.254 | +25.44% (p=0.000 n=10) |
| Channel/N=4096/buffer=64 | Direct → FromChannel | 110629.50 | 183729.50 | 1.661 | +66.08% (p=0.000 n=10) |
| Channel/N=4096/buffer=0 | Seq → ToChannel | 14604.00 | 356143.50 | 24.387 | +2338.67% (p=0.000 n=10) |
| Channel/N=4096/buffer=64 | Direct → RoundTrip | 110629.50 | 668010.50 | 6.038 | +503.83% (p=0.000 n=10) |
| Cache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=1 | Uncached → Cached | 11198.50 | 23290.00 | 2.080 | +107.97% (p=0.000 n=10) |
| Cache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=4 | Uncached → Cached | 303532.50 | 98969.00 | 0.326 | -67.39% (p=0.000 n=10) |
| Cache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=16 | Uncached → Cached | 1234016.50 | 171404.50 | 0.139 | -86.11% (p=0.000 n=10) |

## 読み方

軽い変換では、接続する段数が多いほど関数・反復・構築の費用が表れやすい。  
Early8 は必要な出力がそろうと上流も止まるため、入力件数が多くても全件は処理しない。  
Channel は送受信・goroutine 起動・終了処理を含み、入力容量で対照の絶対時間も変わる。  
Cache は初回保存の時間と要素数に比例する保持メモリを必要とし、上流負荷と利用回数によって再評価を省く効果が変わる。  
Warm の測定値は保存費用を含まないため、利用全体の判断には Reuse の同じ回数の Uncached / Cached を比較する。  

これらは整数と軽い算術を使った条件の結果で、外部 I/O や別の型・出力方式へ一般化しない。  
短い測定・OS・GC・温度等の影響があり、細かい順位や統計的に不明確な差は結論にしない。  
詳細な実装間比較は [Loop](results/pipeline-loop.txt)、[Functions](results/pipeline-functions.txt)、[FromChannel](results/channel-from.txt)、[往復変換](results/channel-roundtrip.txt)、[ToChannel](results/channel-to.txt)、[Cache 再利用](results/cache-reuse.txt) を参照。  

更新は専用の定期更新 Issue で現在の結果を置き換える。  
通常の機能変更 PR やリリースで追加測定を義務付けず、手順は [運用ガイド](../doc/ai/benchmark.md) にまとめる。  
