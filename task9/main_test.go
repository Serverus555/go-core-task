package main

import (
	"fmt"
	"testing"
)

func Test_Ok(t *testing.T) {
	in := make(chan int8)
	go func() {
		defer close(in)

		for i := range 10 {
			in <- int8(i)
		}
	}()

	out := StartConverter(in)

	var i int
	for val := range out {
		fl := float64(i)
		println(fmt.Sprintf("%f", val))
		if fl*fl != val {
			t.Error()
		}
		i++
	}
}
