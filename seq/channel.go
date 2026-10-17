package seq

import "context"

// FromChannel は反復時に channel から受信する Seq を返す。
// 入力を close せず、入力の close・途中終了・ctx のキャンセルで受信を終える。
// 再評価では残りの値を消費し、複数の受信者には値が分配される。
// ctx は非 nil とし、nil の channel はキャンセルまで待機する。
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

// ToChannel は呼び出し時に goroutine で Seq を評価し、バッファなしの channel へ順に送信する。
// 出力は評価終了時に close し、nil の入力は空として扱う。
// ctx は非 nil とし、受信を途中でやめる場合はキャンセルして送信待ちを解除する。
// キャンセルと競合した送信は完了することがある。
// 次の送信値を得るため、入力は受信側より先に評価されることがある。
// 入力内部の待機も解除できるよう、上流にも同じ ctx などで停止手段を用意する。
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
