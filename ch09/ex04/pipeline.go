package pipeline

func pipeline(size int) (chan int, <-chan int) {
	out := make(chan int)
	first := out
	for i := 0; i < size; i++ {
		in := out
		out = make(chan int)
		go func(in, out chan int) {
			for {
				out <- (<-in)
			}
		}(in, out)
	}

	return first, out
}
