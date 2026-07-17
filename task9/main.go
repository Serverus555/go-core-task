package main

func StartConverter(in <-chan int8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)

		for val := range in {
			converted := float64(val)
			out <- converted * converted
		}
	}()

	return out
}
