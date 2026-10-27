# API 別の比較

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  
実装名・時間の列名から、測定した処理のコードを開ける。  

Go はライブラリを使わない同一処理、Library は対応する API を使った処理。  
Compose の Pipe は個別接続、Composed は合成したオペレーターを接続する処理。  
倍率は各表の右側の実装の時間 / 左側の実装の時間で、1 を超えると右側が遅い。  
検定の `~` は差が不明確で、増減率の正値は右側の時間・割り当ての増加を表す。  

## パイプ接続

### Compose / Compose2 の測定条件

以下の Compose / Compose2 の表には、この測定情報を適用する。  
他の API の測定情報は [実行環境](environment.md) を参照。  
入力・処理・構築範囲は [測定条件](../METHOD.md#compose--compose2) を参照。  

| 項目 | 値 |
| --- | --- |
| 測定日時（UTC） | 2026-10-27T13:48:47+00:00 |
| 対象コード（ライブラリ・測定） | [f523f1e84688](https://github.com/meian/go-seqpipeline/commit/f523f1e8468892d07ec296f8236478b455dad029) |
| Go | go version go1.27.3 darwin/arm64 |
| OS / CPU | macOS-27.0.1-arm64-arm-64bit / Apple M4 |
| GOMAXPROCS | 2 |
| 測定時間 / 測定回数 | 500ms / 10 |
| GOGC / GOMEMLIMIT / GOEXPERIMENT | 既定値 / 既定値 / 既定値 |

通常の最適化は有効で、CPU・温度・OS 負荷は固定していない。  

### Compose.Seq：段数=2 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L70) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L51) | 倍率 |
| --- | --- | --- | --- |
| 64 | 270.65 ±0% | 333.95 ±1% | 1.234× |
| 4096 | 17,024 ±1% | 17,390.5 ±0% | 1.022× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L70) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L51) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L70) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L51) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +23.39% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +2.15% | p=0.001 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq：段数=2 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L33) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L14) | 倍率 |
| --- | --- | --- | --- |
| 64 | 270.5 ±0% | 334.6 ±2% | 1.237× |
| 4096 | 17,017.5 ±0% | 17,358.5 ±0% | 1.020× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L33) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L14) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L33) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L14) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +23.70% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +2.00% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq：段数=4 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L152) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L131) | 倍率 |
| --- | --- | --- | --- |
| 64 | 447.35 ±0% | 594.95 ±2% | 1.330× |
| 4096 | 27,938.5 ±0% | 28,119 ±1% | 1.006× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L152) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L131) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L152) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L131) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +32.99% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +0.65% | p=0.004 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq：段数=4 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L111) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L90) | 倍率 |
| --- | --- | --- | --- |
| 64 | 434.35 ±0% | 596.4 ±2% | 1.373× |
| 4096 | 27,155.5 ±0% | 28,154 ±0% | 1.037× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L111) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L90) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L111) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L90) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +37.31% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +3.68% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：段数=2 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L400) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L381) | 倍率 |
| --- | --- | --- | --- |
| 64 | 31.03 ±3% | 336.45 ±2% | 10.843× |
| 4096 | 1,985 ±2% | 16,647.5 ±0% | 8.387× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L400) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L381) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L400) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L381) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +984.27% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +738.66% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：段数=2 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L363) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L344) | 倍率 |
| --- | --- | --- | --- |
| 64 | 31.08 ±3% | 337.8 ±2% | 10.869× |
| 4096 | 1,992.5 ±3% | 16,665 ±0% | 8.364× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L363) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L344) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L363) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L344) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +986.87% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +736.39% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：段数=4 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L482) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L461) | 倍率 |
| --- | --- | --- | --- |
| 64 | 516.9 ±0% | 599.45 ±1% | 1.160× |
| 4096 | 32,282.5 ±0% | 32,242.5 ±0% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L482) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L461) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L482) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L461) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +15.97% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | ~ | p=0.105 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose.Seq2：段数=4 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L441) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L420) | 倍率 |
| --- | --- | --- | --- |
| 64 | 506.1 ±0% | 603.9 ±2% | 1.193× |
| 4096 | 31,685.5 ±0% | 32,287 ±0% | 1.019× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L441) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L420) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L441) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L420) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.32% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.90% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：段数=2 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L730) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L711) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.935 ±2% | 333.85 ±1% | 19.714× |
| 4096 | 1,165 ±1% | 16,728 ±0% | 14.359× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L730) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L711) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L730) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L711) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1871.36% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1335.88% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：段数=2 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L693) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L674) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.07 ±2% | 335.35 ±1% | 19.646× |
| 4096 | 1,168.5 ±1% | 16,697 ±0% | 14.289× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L693) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L674) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L693) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L674) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1864.56% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1328.93% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：段数=4 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L812) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L791) | 倍率 |
| --- | --- | --- | --- |
| 64 | 467.6 ±0% | 594.75 ±1% | 1.272× |
| 4096 | 29,154.5 ±0% | 29,476.5 ±0% | 1.011× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L812) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L791) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L812) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L791) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +27.19% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.10% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq：段数=4 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L771) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L750) | 倍率 |
| --- | --- | --- | --- |
| 64 | 466.95 ±0% | 596.4 ±2% | 1.277× |
| 4096 | 29,070 ±0% | 29,440.5 ±0% | 1.013× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L771) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L750) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L771) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L750) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +27.72% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +1.27% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：段数=2 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1060) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1041) | 倍率 |
| --- | --- | --- | --- |
| 64 | 276.4 ±0% | 335.9 ±1% | 1.215× |
| 4096 | 17,393.5 ±0% | 17,448 ±0% | 1.003× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1060) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1041) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1060) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1041) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +21.53% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +0.31% | p=0.003 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：段数=2 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1023) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1004) | 倍率 |
| --- | --- | --- | --- |
| 64 | 275.85 ±0% | 336.7 ±1% | 1.221× |
| 4096 | 17,380.5 ±0% | 17,464.5 ±0% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1023) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1004) | 272 ±0% | 9 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1023) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1004) | 272 ±0% | 9 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +22.06% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +0.48% | p=0.022 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：段数=4 / 構築済みを利用

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1142) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1121) | 倍率 |
| --- | --- | --- | --- |
| 64 | 417.9 ±1% | 578.75 ±1% | 1.385× |
| 4096 | 25,462.5 ±1% | 27,227.5 ±1% | 1.069× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1142) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1121) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1142) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1121) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +38.49% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +6.93% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Compose2.Seq2：段数=4 / 構築込み

Pipe：`Pipe` / `Pipe2` で同じオペレーター列を個別に接続。  

| N | [Pipe ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1101) | [Composed ns/op](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1080) | 倍率 |
| --- | --- | --- | --- |
| 64 | 417.85 ±2% | 578.05 ±1% | 1.383× |
| 4096 | 25,694.5 ±3% | 27,186 ±0% | 1.058× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1101) | 0 ±0% | 0 ±0% |
| 64 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1080) | 464 ±0% | 15 ±0% |
| 4096 | [Pipe](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1101) | 0 ±0% | 0 ±0% |
| 4096 | [Composed](https://github.com/meian/go-seqpipeline/blob/f523f1e8468892d07ec296f8236478b455dad029/benchmarks/compose_test.go#L1080) | 464 ±0% | 15 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +38.34% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5.80% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
### Seq.Pipe：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.425 ±1% | 17.645 ±0% | 1.013× |
| 4096 | 1,023 ±0% | 1,024 ±0% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.26% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.512 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq.Pipe2：構築・評価

Go：`slices.Values` の値と変換した値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.725 ±0% | 18.55 ±0% | 0.991× |
| 4096 | 1,146.5 ±2% | 1,152 ±0% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.93% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +0.48% | p=0.049 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe：構築・評価

Go：`slices.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L202) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L163) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.88 ±1% | 16.935 ±1% | 1.003× |
| 4096 | 979.35 ±1% | 981.55 ±1% | 1.002× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L202) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L163) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.109 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.183 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Seq2.Pipe2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L278) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L239) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.735 ±0% | 18.6 ±0% | 0.993× |
| 4096 | 1,180.5 ±0% | 1,205 ±0% | 1.021× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L278) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/pipe_test.go#L239) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.72% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +2.08% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## 変換・選別・射影

### Filter：構築・評価

Go：`slices.Values` で値を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.465 ±0% | 28.835 ±2% | 0.862× |
| 4096 | 2,037 ±0% | 1,697 ±0% | 0.833× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -13.84% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.69% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Filter2：構築・評価

Go：`slices.All` で組を選別。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L128) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L89) | 倍率 |
| --- | --- | --- | --- |
| 64 | 33.44 ±0% | 28.75 ±0% | 0.860× |
| 4096 | 2,040 ±0% | 1,697 ±0% | 0.832× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L128) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L89) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L128) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/filter_test.go#L89) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -14.03% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -16.81% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Keys：構築・評価

Go：`maps.All` からキーを取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 307.6 ±10% | 317 ±4% | 1.031× |
| 4096 | 21,928.5 ±2% | 22,057 ±1% | 1.006× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/keys_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.481 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.912 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map：構築・評価

Go：`slices.Values` で値を変換。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.46 ±0% | 17.72 ±0% | 1.015× |
| 4096 | 1,026.5 ±2% | 1,025.5 ±0% | 0.999× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.49% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.592 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Map2：構築・評価

Go：`slices.All` で値を変換し、キーを維持。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 18.74 ±0% | 18.66 ±0% | 0.996× |
| 4096 | 1,180.5 ±0% | 1,131 ±0% | 0.958× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/map_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | -0.43% | p=0.001 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | -4.19% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Values：構築・評価

Go：`maps.All` から値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 318.5 ±6% | 314.05 ±6% | 0.986× |
| 4096 | 22,000.5 ±3% | 21,799 ±1% | 0.991× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/values_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.684 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.143 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

## Channel

### FromChannel：構築・評価

Go：channel を直接受信し、送信元の終了を待つ。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,833 ±1% | 6,945.5 ±1% | 1.191× |
| 4096 | 350,184 ±2% | 418,145 ±2% | 1.194× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L50) | 272 ±0% | 3 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_channel_test.go#L10) | 360 ±0% | 6 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +19.07% | p=0.000 n=10 |
| 64 | B/op | +32.35% | p=0.000 n=10 |
| 64 | allocs/op | +100.00% | p=0.000 n=10 |
| 4096 | sec/op | +19.41% | p=0.000 n=10 |
| 4096 | B/op | +32.35% | p=0.000 n=10 |
| 4096 | allocs/op | +100.00% | p=0.000 n=10 |

</details>

### ToChannel：構築・評価

Go：1 goroutine で送信し、channel を直接受信。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 5,819 ±1% | 7,083.5 ±1% | 1.217× |
| 4096 | 352,309.5 ±1% | 430,760.5 ±1% | 1.223× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L50) | 160 ±0% | 2 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_channel_test.go#L11) | 304 ±0% | 5 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +21.73% | p=0.000 n=10 |
| 64 | B/op | +90.00% | p=0.000 n=10 |
| 64 | allocs/op | +150.00% | p=0.000 n=10 |
| 4096 | sec/op | +22.27% | p=0.000 n=10 |
| 4096 | B/op | +90.00% | p=0.000 n=10 |
| 4096 | allocs/op | +150.00% | p=0.000 n=10 |

</details>

## 入力変換・出力収集

### FromMap：構築・評価

Go：`maps.All` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 316.55 ±8% | 315.9 ±4% | 0.998× |
| 4096 | 22,131.5 ±2% | 22,005 ±1% | 0.994× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_map_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.579 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.218 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromRange.Signed：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.385 ±1% | 101.45 ±1% | 6.192× |
| 4096 | 947.05 ±1% | 4,819.5 ±1% | 5.089× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +519.16% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +408.90% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Signed：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.105 ±1% | 103.4 ±0% | 6.420× |
| 4096 | 936.45 ±1% | 4,856 ±1% | 5.186× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L62) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L9) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +542.04% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +418.55% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Unsigned：昇順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.315 ±1% | 108.2 ±1% | 6.632× |
| 4096 | 942.95 ±1% | 5,036 ±0% | 5.341× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +563.19% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +434.07% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromRange.Unsigned：降順

Go：同じ範囲を整数ループで加算。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.805 ±0% | 106.8 ±0% | 6.757× |
| 4096 | 928.8 ±0% | 5,091.5 ±0% | 5.482× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L173) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_range_test.go#L120) | 88 ±0% | 3 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +575.74% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +448.18% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### FromSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.02 ±1% | 31.73 ±0% | 1.864× |
| 4096 | 986.3 ±1% | 1,914.5 ±0% | 1.941× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +86.43% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +94.11% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.17 ±0% | 17.18 ±0% | 1.001× |
| 4096 | 1,092 ±1% | 1,095 ±1% | 1.003× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.669 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.266 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### FromSlice：構築・評価

Go：`slices.Values` を反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L48) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.925 ±1% | 16.935 ±1% | 1.001× |
| 4096 | 983 ±1% | 983.6 ±1% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L48) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/from_slice_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.810 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.867 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToMap：収集

Go：`maps.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L52) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L13) | 倍率 |
| --- | --- | --- | --- |
| 64 | 2,268.5 ±1% | 2,350 ±1% | 1.036× |
| 4096 | 150,096.5 ±1% | 153,767.5 ±1% | 1.024× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L52) | 4,648 ±0% | 11 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L13) | 4,688 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L52) | 296,120 ±0% | 48 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_map_test.go#L13) | 296,160 ±0% | 50 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +3.59% | p=0.000 n=10 |
| 64 | B/op | +0.86% | p=0.000 n=10 |
| 64 | allocs/op | +18.18% | p=0.000 n=10 |
| 4096 | sec/op | +2.45% | p=0.000 n=10 |
| 4096 | B/op | +0.01% | p=0.000 n=10 |
| 4096 | allocs/op | +4.17% | p=0.000 n=10 |

</details>

### ToSeq：構築・評価

Go：標準 `iter.Seq` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L49) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L10) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.095 ±1% | 17.31 ±5% | 1.013× |
| 4096 | 991.05 ±1% | 996.45 ±1% | 1.005× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L49) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L10) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +1.26% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +0.54% | p=0.009 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSeq2：構築・評価

Go：標準 `iter.Seq2` を `for range`。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L126) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L87) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.215 ±0% | 17.22 ±0% | 1.000× |
| 4096 | 1,093.5 ±1% | 1,095 ±0% | 1.001× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L126) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_seq_test.go#L87) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | ~ | p=0.446 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | ~ | p=0.399 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### ToSlice：収集

Go：`slices.Collect` で収集。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L12) | 倍率 |
| --- | --- | --- | --- |
| 64 | 156.75 ±1% | 261.2 ±1% | 1.666× |
| 4096 | 9,731.5 ±1% | 20,577 ±2% | 2.114× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L51) | 1,016 ±0% | 7 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L12) | 1,080 ±0% | 10 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L51) | 128,248 ±0% | 16 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/to_slice_test.go#L12) | 128,312 ±0% | 19 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +66.63% | p=0.000 n=10 |
| 64 | B/op | +6.30% | p=0.000 n=10 |
| 64 | allocs/op | +42.86% | p=0.000 n=10 |
| 4096 | sec/op | +111.45% | p=0.000 n=10 |
| 4096 | B/op | +0.05% | p=0.000 n=10 |
| 4096 | allocs/op | +18.75% | p=0.000 n=10 |

</details>

## 件数制御・結合

### Skip.Operator：構築・評価

Go：先頭 8 件を除き `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 15.475 ±0% | 30.705 ±0% | 1.984× |
| 4096 | 993 ±0% | 1,957 ±0% | 1.971× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +98.42% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +97.08% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Skip.Operator2：構築・評価

Go：先頭 8 件を除き `slices.All` で元の添字・値を取得。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 16.1 ±1% | 30.705 ±0% | 1.907× |
| 4096 | 1,157.5 ±0% | 1,969 ±0% | 1.701× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/skip_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +90.71% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +70.11% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator：先頭8件

Go：先頭 8 件を `slices.Values` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L50) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.0445 ±3% | 3.5315 ±2% | 1.160× |
| 4096 | 2.9625 ±7% | 3.536 ±1% | 1.194× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L50) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L11) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L50) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L11) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16.00% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +19.36% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Take.Operator2：先頭8件

Go：先頭 8 件を `slices.All` で反復。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L127) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 3.063 ±4% | 3.4605 ±0% | 1.130× |
| 4096 | 3.029 ±6% | 3.42 ±1% | 1.129× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L127) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L88) | 0 ±0% | 0 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L127) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/take_test.go#L88) | 0 ±0% | 0 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +12.98% | p=0.000 n=10 |
| 64 | B/op | ~ | p=1.000 n=10 |
| 64 | allocs/op | ~ | p=1.000 n=10 |
| 4096 | sec/op | +12.91% | p=0.000 n=10 |
| 4096 | B/op | ~ | p=1.000 n=10 |
| 4096 | allocs/op | ~ | p=1.000 n=10 |

</details>

### Zip.AlignBy：構築・評価

Go：左入力を反復し、右の不足分を 0 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L128) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L88) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.92 ±1% | 1,658.5 ±1% | 63.985× |
| 4096 | 1,580 ±0% | 94,144 ±1% | 59.585× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L128) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L88) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L128) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L88) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6298.53% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5858.48% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L51) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L11) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.135 ±0% | 2,904 ±1% | 169.478× |
| 4096 | 1,008 ±0% | 173,100 ±1% | 171.726× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L51) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L11) | 376 ±0% | 11 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L51) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L11) | 376 ±0% | 11 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16847.77% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +17072.62% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L293) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L252) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.9 ±0% | 1,662.5 ±1% | 64.189× |
| 4096 | 1,577.5 ±0% | 93,917.5 ±1% | 59.536× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L293) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L252) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L293) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L252) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6318.92% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5853.57% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue：構築・評価

Go：左入力を反復し、右の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L210) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L170) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.905 ±0% | 1,654.5 ±1% | 63.868× |
| 4096 | 1,577.5 ±0% | 93,885 ±1% | 59.515× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L210) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L170) | 424 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L210) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L170) | 424 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6286.80% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5851.51% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.AlignBy（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 7 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L453) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L412) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.8 ±0% | 1,666 ±1% | 64.574× |
| 4096 | 1,575 ±0% | 94,352 ±1% | 59.906× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L453) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L412) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L453) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L412) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6357.36% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5890.60% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.Operator（キー既定値あり）：構築・評価

Go：添字で同じ位置の値を組にする。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L375) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L335) | 倍率 |
| --- | --- | --- | --- |
| 64 | 17.04 ±1% | 2,882 ±1% | 169.131× |
| 4096 | 1,003.5 ±0% | 172,664 ±1% | 172.062× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L375) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L335) | 440 ±0% | 12 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L375) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L335) | 440 ±0% | 12 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +16813.15% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +17106.18% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultKey（キー既定値あり）：構築・評価

Go：右入力を反復し、左の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L618) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L577) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.75 ±0% | 2,805 ±1% | 108.932× |
| 4096 | 1,575 ±0% | 163,259 ±1% | 103.657× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L618) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L577) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L618) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L577) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +10793.20% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +10265.65% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>

### Zip.WithDefaultValue（キー既定値あり）：構築・評価

Go：左入力を反復し、右の不足分を 9 で補完。  

| N | [Go ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L535) | [Library ns/op](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L495) | 倍率 |
| --- | --- | --- | --- |
| 64 | 25.835 ±0% | 1,676 ±1% | 64.873× |
| 4096 | 1,574 ±0% | 94,472 ±1% | 60.020× |

| N | 実装 | B/op | allocs/op |
| --- | --- | --- | --- |
| 64 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L535) | 0 ±0% | 0 ±0% |
| 64 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L495) | 488 ±0% | 13 ±0% |
| 4096 | [Go](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L535) | 0 ±0% | 0 ±0% |
| 4096 | [Library](https://github.com/meian/go-seqpipeline/blob/6e84c08e445a5977b33f75bb23ff97a48063892c/benchmarks/zip_test.go#L495) | 488 ±0% | 13 ±0% |

<details>
<summary>検定</summary>

| N | 指標 | 増減率 | 検定 |
| --- | --- | --- | --- |
| 64 | sec/op | +6387.32% | p=0.000 n=10 |
| 64 | B/op | ? | p=0.000 n=10 |
| 64 | allocs/op | ? | p=0.000 n=10 |
| 4096 | sec/op | +5902.03% | p=0.000 n=10 |
| 4096 | B/op | ? | p=0.000 n=10 |
| 4096 | allocs/op | ? | p=0.000 n=10 |

</details>
