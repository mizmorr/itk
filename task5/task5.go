package task5

import "sort"

func GetEqualsSort(first, second []int) (found bool, equals []int) {
	equals = make([]int, 0)
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
	if len(equals) != 0 {
		found = true
	}
	return
}
