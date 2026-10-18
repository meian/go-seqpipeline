package benchmarks

func producer(values []int, capacity int) (<-chan int, <-chan struct{}) {
	ch := make(chan int, capacity)
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch, done
}
