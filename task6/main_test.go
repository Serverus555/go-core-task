package main

import (
	"context"
	"testing"
	"time"
)

func Test_Ok(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rg := RandomGenerator(ctx)

	_, ok := <-rg
	if !ok {
		t.Error()
	}
	cancel()
	time.Sleep(100 * time.Millisecond)
	_, ok = <-rg
	if ok {
		t.Error()
	}
}
