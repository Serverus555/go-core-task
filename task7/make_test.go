package main

import (
	"slices"
	"testing"
)

const batchSize = 100
const batchesCount = 5

func Test_Ok(t *testing.T) {
	outChs := make([]<-chan int, 0, batchesCount)

	for i := range batchesCount {
		ch := make(chan int)
		outChs = append(outChs, ch)
		go writer(ch, i*batchSize)
	}

	mergedCh := MergeChannels(outChs)
	mergedNums := make([]int, 0, batchesCount*batchSize)

	for val := range mergedCh {
		mergedNums = append(mergedNums, val)
	}

	slices.Sort(mergedNums)
	for i := range batchesCount * batchSize {
		if mergedNums[i] != i {
			t.Error()
		}
	}
}

func writer(ch chan<- int, offset int) {
	for num := range batchSize {
		ch <- offset + num
	}
	close(ch)
}
