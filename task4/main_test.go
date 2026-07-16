package main

import (
	"slices"
	"testing"
)

func Test_Ok(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	slices.Sort(expected)

	result := Subtract(slice1, slice2)
	slices.Sort(result)

	if !slices.Equal(expected, result) {
		t.Error()
	}
}
