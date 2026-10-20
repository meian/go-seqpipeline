package seq

import "context"

// FromChannel は反復時に channel から受信する seq.Seq を返す。
// 元の channel を close して残りの値を読み終えるか、ctx をキャンセルするか、break することで反復を終了する。
// 入力 channel は、外部で明示的に close しない限り close されない。
// 再評価では先頭に戻らず、後続の値を受信する。
// 同じ channel から複数の seq.Seq を作ると、受信タイミングに応じて値が分配される。
// ctx に nil を渡すと呼び出し時に panic する。
// nil の channel はキャンセルまで待機する。
// キャンセルと競合した値を返すことがある。送信者の停止は呼び出し側で管理する。
func FromChannel[T any](ctx context.Context, ch <-chan T) Seq[T] {
	done := ctx.Done()
	return func(yield func(T) bool) {
		for {
			if ctx.Err() != nil {
				return
			}
			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok || !yield(v) {
					return
				}
			}
		}
	}
}

// ToChannel は呼び出し時に goroutine で seq.Seq を評価し、バッファなしの channel へ順に送信する。
// 出力は評価終了時に close し、nil の入力は空として扱う。
// ctx に nil を渡すと呼び出し時に panic する。
// 受信を途中でやめる場合は ctx をキャンセルして送信待ちを解除する。
// キャンセルと競合した送信は完了することがある。
// 次の送信値を得るため、入力は受信側より先に評価されることがある。
// 入力内部の待機を解除するためには、上流にも同じ ctx などで停止手段を用意する必要がある。
func ToChannel[T any](ctx context.Context, s Seq[T]) <-chan T {
	done := ctx.Done()
	ch := make(chan T)
	go func() {
		defer close(ch)
		if s == nil || ctx.Err() != nil {
			return
		}
		for v := range s {
			if ctx.Err() != nil {
				return
			}
			select {
			case <-done:
				return
			case ch <- v:
			}
		}
	}()
	return ch
}
