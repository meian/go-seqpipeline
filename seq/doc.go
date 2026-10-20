// Package seq は iter を利用したシーケンスと入出力変換を提供する。
// Seq と Seq2 の Pipe / Pipe2 に op パッケージのオペレーターを接続し、
// パイプラインを構築する。
// Pipe / Pipe2 は接続時にオペレーターを呼び出し、入力の評価と停止伝播はその実装に従う。
// 再評価や並行評価の可否は入力とコールバックに従う。
//
// FromSeq / FromSeq2 と ToSeq / ToSeq2 は要素を評価・コピーせず、nil を保持する。
// Seq / Seq2 のゼロ値は nil であり、組み込みオペレーターは空として扱う。
// nil のシーケンスを反復したり、ToSlice / ToMap に渡したりすると panic する。
// ToChannel は nil のシーケンスを空として扱う。
package seq
