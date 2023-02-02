package enumerable

func Range(start, count int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := start; i <= count; i++ {
			ch <- i
		}
	}()

	return ch
}
