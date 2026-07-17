package main

import (
	"math/rand/v2"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Ok(t *testing.T) {

	wg := NewMyWaitGroup()
	var counter atomic.Int32
	for range 10 {
		wg.Add()
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.IntN(200) * int(time.Millisecond)))
			counter.Add(1)
		}()
	}

	wg.Wait()
	if counter.Load() != 10 {
		t.Error()
	}
}
