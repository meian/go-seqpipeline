package seq

import "context"

// FromChannel は受信 channel を、反復時に値を受信する Seq に変換する。
// channel の受信順序と重複を保持し、入力の close または ctx のキャンセルで終了する。
// ctx は非 nil とする。nil の channel はキャンセルまで待機する。
// goroutine を作らず、入力を close しない。yield が false を返すと追加の受信をせず終了する。
// 入力は消費型であり、再評価では残りの値を受信する。複数の受信者には値が分配される。
// キャンセルと受信が競合した場合、受信した値を yield することがある。
// 入力の送信者の停止は呼び出し側で管理する。
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

// ToChannel は Seq の評価を専用 goroutine で開始し、値を反復順に送信する。
// 呼び出すたびにバッファなしの受信専用 channel を新しく作り、順序と重複を保持する。
// goroutine は入力の終了または ctx のキャンセルで出力を close する。nil の入力は空として扱う。
// ctx は非 nil とする。受信を途中でやめる場合は、呼び出し側で ctx をキャンセルする。
// キャンセルは送信待ちを解除し、入力へ停止を伝えるが、競合した送信が完了することもある。
// 入力は次の送信値を得るために、受信側より先に評価されることがある。
// 入力やコールバック内部での待機は中断できないため、それらも同じ ctx などで停止可能にする。
// 再評価や並行評価の可否は入力に従い、呼び出し側で管理する。
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
