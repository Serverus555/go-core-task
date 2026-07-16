package main

import "sync"

func MergeChannels[T any](channels []<-chan T) <-chan T {
	merged := make(chan T)

	var allClosed sync.WaitGroup

	for _, channel := range channels {
		allClosed.Add(1)
		go forward(&allClosed, channel, merged)
	}

	go func() {
		allClosed.Wait()
		close(merged)
	}()

	return merged
}

func forward[T any](closed *sync.WaitGroup, in <-chan T, out chan<- T) {
	defer closed.Done()
	for val := range in {
		out <- val
	}
}
