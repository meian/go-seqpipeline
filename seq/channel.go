package seq

import "context"

// FromChannel returns a Seq that receives from ch during iteration.
// Iteration ends when the input channel is closed and drained, ctx is canceled,
// or the consumer breaks out of iteration.
// FromChannel does not close ch, and received values cannot be retrieved again.
// If multiple sequences read from the same channel, they divide values according
// to receive timing.
// A nil ctx panics when FromChannel is called.
// A nil channel waits until cancellation.
// A value racing with cancellation may still be yielded; callers are responsible
// for stopping the sender.
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

// ToChannel evaluates a Seq in a goroutine started when called, sending values
// in order through an unbuffered channel.
// The goroutine closes the output channel when evaluation ends.
// A nil ctx panics when ToChannel is called.
// If the receiver stops early, cancel ctx to release a blocked send.
// A send racing with cancellation may still complete.
// The input may be evaluated ahead of the receiver to obtain the next value.
// To release waits inside the input, provide an upstream stopping mechanism,
// such as the same ctx.
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
