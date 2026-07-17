package main

import (
	"slices"
	"testing"
)

func Test_Example(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targetSlice := []int{2, 4, 6, 8, 10}
	result := sliceExample(originalSlice)
	if !slices.Equal(result, targetSlice) {
		t.Error()
	}
}

func Test_AddOkNoSideEffect(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := addElement(originalSlice, 11)
	if len(result) != 11 || result[10] != 11 || slices.Equal(originalSlice, result) {
		t.Error()
	}
}

func Test_CopyOkNoSideEffect(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targetSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := copySlice(originalSlice)

	if !slices.Equal(result, targetSlice) {
		t.Error()
	}
	for i, _ := range originalSlice {
		originalSlice[i] = -1
	}
	if !slices.Equal(result, targetSlice) {
		t.Error()
	}
}

func Test_RemoveOkNoSideEffect(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targetSlice := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}

	result := removeElement(originalSlice, 1)

	if !slices.Equal(result, targetSlice) {
		t.Error()
	}
	for i, _ := range originalSlice {
		originalSlice[i] = -1
	}
	if !slices.Equal(result, targetSlice) {
		t.Error()
	}
}
