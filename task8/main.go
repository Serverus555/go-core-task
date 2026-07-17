package main

import (
	"sync/atomic"
)

type MyWaitGroup struct {
	count atomic.Int64
	sema  chan struct{}
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		atomic.Int64{},
		make(chan struct{}, 1),
	}
}

func (w *MyWaitGroup) Add() {
	w.count.Add(1)
}

func (w *MyWaitGroup) Done() {
	c := w.count.Add(-1)
	if c == 0 {
		w.sema <- struct{}{}
	}
}

func (w *MyWaitGroup) Wait() {
	<-w.sema
}
