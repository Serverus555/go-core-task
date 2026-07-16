package main

func Intersection(first []int, second []int) ([]int, bool) {
	var small, large []int
	if len(first) < len(second) {
		small, large = first, second
	} else {
		small, large = second, first
	}

	smallSet := map[int]struct{}{}
	for _, v := range small {
		smallSet[v] = struct{}{}
	}

	var result []int
	for _, v := range large {
		if _, ok := smallSet[v]; ok {
			result = append(result, v)
			delete(smallSet, v)
		}
	}
	return result, len(result) > 0
}
