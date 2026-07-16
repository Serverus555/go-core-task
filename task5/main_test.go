package main

import (
	"slices"
	"testing"
)

func Test_Ok(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{3, 64}

	result, ok := Intersection(a, b)
	slices.Sort(result)

	if !slices.Equal(expected, result) || !ok {
		t.Error()
	}
}
