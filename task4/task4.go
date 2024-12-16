package task4

import (
	"sort"
)

func GetEqualsMap(first []int, second []int) (result []int) {
	mapForFirst := make(map[int]interface{})

	for _, num := range first {
		mapForFirst[num] = struct{}{}
	}

	for _, num := range second {
		if _, exists := mapForFirst[num]; exists {
			result = append(result, num)
			delete(mapForFirst, num)
		}
	}
	return result
}

func GetEqualsSort(first, second []int) (equals []int) {
	sort.Ints(first)
	sort.Ints(second)

	minLen := len(first)

	if len(second) < minLen {
		minLen = len(second)
	}

	for i, j := 0, 0; i < len(first) && j < len(second); {
		if first[i] == second[j] {
			equals = append(equals, second[j])
			i, j = i+1, j+1
		} else if first[i] < second[j] {
			i++
		} else {
			j++
		}
	}
	return
}
