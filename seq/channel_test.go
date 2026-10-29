package seq_test

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFromChannel(t *testing.T) {
	t.Parallel()

	ch := make(chan int, 4)
	for _, v := range []int{3, 1, 1, 2} {
		ch <- v
	}
	src := seq.FromChannel(t.Context(), ch)
	assertEqual(t, len(ch), 4) // 構築時には受信しない。
	for v := range src {
		assertEqual(t, v, 3)
		break
	}
	assertEqual(t, len(ch), 3) // yield(false) 後は受信しない。
	ch <- 4                    // 入力の所有権は呼び出し側にある。
	close(ch)
	assertEqual(t, seq.ToSlice(src), []int{1, 1, 2, 4})
	assertEqual(t, seq.ToSlice(src), []int(nil))
}

func TestFromChannelClosed(t *testing.T) {
	t.Parallel()

	ch := make(chan int)
	close(ch)
	assertEqual(t, seq.ToSlice(seq.FromChannel(t.Context(), ch)), []int(nil))
}

func TestFromChannelCancelReceive(t *testing.T) {
	t.Parallel()

	for _, nilChannel := range []bool{false, true} {
		name := "open"
		if nilChannel {
			name = "nil"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				var ch chan int
				if !nilChannel {
					ch = make(chan int)
				}
				finished := make(chan struct{})
				go func() {
					defer close(finished)
					for range seq.FromChannel(ctx, ch) {
						t.Error("received from an empty channel")
					}
				}()
				synctest.Wait() // 受信待ちになってからキャンセルする。
				select {
				case <-finished:
					t.Fatal("iteration ended before cancellation")
				default:
				}
				cancel()
				<-finished
			})
		})
	}
}

func TestFromChannelCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	ch := make(chan int, 1)
	ch <- 7
	assertEqual(t, seq.ToSlice(seq.FromChannel(ctx, ch)), []int(nil))
	assertEqual(t, len(ch), 1)
}

func TestFromChannelSharedReceivers(t *testing.T) {
	t.Parallel()

	ch := make(chan int, 3)
	for _, v := range []int{1, 2, 3} {
		ch <- v
	}
	close(ch)
	first := seq.FromChannel(t.Context(), ch)
	second := seq.FromChannel(t.Context(), ch)
	for v := range first {
		assertEqual(t, v, 1)
		break
	}
	assertEqual(t, seq.ToSlice(second), []int{2, 3})
	assertEqual(t, seq.ToSlice(first), []int(nil))
}

func TestFromChannelCloseReceive(t *testing.T) {
	t.Parallel()

	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int)
		finished := make(chan struct{})
		go func() {
			defer close(finished)
			for range seq.FromChannel(context.Background(), ch) {
				t.Error("empty channel produced a value")
			}
		}()
		synctest.Wait()
		close(ch)
		<-finished
	})
}

// nil context の panic は、反復前の呼び出し元で捕捉できる。
func TestFromChannelNilContext(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("nil context must panic during FromChannel")
		}
	}()
	_ = seq.FromChannel(nil, make(chan int))
}

func TestToChannel(t *testing.T) {
	t.Parallel()

	src := seq.FromSlice([]int{3, 1, 1, 2})
	for round := 0; round < 2; round++ {
		ch := seq.ToChannel(t.Context(), src)
		assertEqual(t, cap(ch), 0)
		var got []int
		for v := range ch {
			got = append(got, v)
		}
		assertEqual(t, got, []int{3, 1, 1, 2})
	}
}

func TestToChannelEmpty(t *testing.T) {
	t.Parallel()

	for _, src := range []seq.Seq[int]{nil, seq.FromSlice([]int{})} {
		ch := seq.ToChannel(t.Context(), src)
		if ch == nil {
			t.Fatal("output channel must not be nil")
		}
		if _, ok := <-ch; ok {
			t.Fatal("empty input produced a value")
		}
	}
}

func TestToChannelCancelSend(t *testing.T) {
	t.Parallel()

	for _, receiveFirst := range []bool{false, true} {
		name := "without_receiver"
		if receiveFirst {
			name = "after_receiving"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				reads, stops := 0, 0
				finished := make(chan struct{})
				src := seq.Seq[int](func(yield func(int) bool) {
					defer close(finished)
					for _, v := range []int{1, 2, 3} {
						reads++
						if !yield(v) {
							stops++
							return
						}
					}
				})
				ch := seq.ToChannel(ctx, src)
				wantReads := 1
				if receiveFirst {
					assertEqual(t, <-ch, 1)
					wantReads = 2
				}
				synctest.Wait() // 評価開始とバッファなし送信の待機を確認する。
				assertEqual(t, reads, wantReads)
				assertEqual(t, stops, 0)
				cancel()
				synctest.Wait()
				<-finished
				assertEqual(t, stops, 1)
				assertEqual(t, reads, wantReads)
				if _, ok := <-ch; ok {
					t.Fatal("output must close after cancellation")
				}
			})
		})
	}
}

func TestToChannelCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	src := seq.Seq[int](func(func(int) bool) { t.Error("canceled input evaluated") })
	if _, ok := <-seq.ToChannel(ctx, src); ok {
		t.Fatal("canceled output must be closed")
	}
}

// nil context の panic は、goroutine 内ではなく呼び出し元で捕捉できる。
func TestToChannelNilContext(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("nil context must panic during ToChannel")
		}
	}()
	_ = seq.ToChannel(nil, seq.FromSlice([]int{1}))
}
