package main

import "fmt"

func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(fmt.Sprintf("%v", sliceExample(originalSlice)))
	println(fmt.Sprintf("%v", addElement(originalSlice, 11)))
	println(fmt.Sprintf("%v", copySlice(originalSlice)))
	println(fmt.Sprintf("%v", removeElement(originalSlice, 1)))
}

func sliceExample(slice []int) []int {
	result := make([]int, 0)
	for _, val := range slice {
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	return result
}

func addElement(origin []int, val int) []int {
	result := make([]int, len(origin)+1)
	copy(result, origin)
	result[len(origin)] = val
	return result
}

func copySlice(origin []int) []int {
	result := make([]int, len(origin))
	copy(result, origin)
	return result
}

func removeElement(origin []int, index int) []int {
	result := make([]int, len(origin)-1)
	copy(result, origin[:index])
	copy(result[index:], origin[index+1:])
	return result
}
