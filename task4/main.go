package main

func Subtract(first []string, second []string) []string {
	secondSet := map[string]struct{}{}
	for _, v := range second {
		secondSet[v] = struct{}{}
	}

	var result []string
	for _, v := range first {
		if _, ok := secondSet[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
