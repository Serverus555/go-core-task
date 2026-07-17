package main

import (
	"context"
	"math/rand"
)

func RandomGenerator(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case out <- rand.Int():
			}
		}
	}()

	return out
}
