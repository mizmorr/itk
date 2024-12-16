package task4

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateSlice(n int) []int {
	resultSlice := make([]int, n)
	for i := 0; i < n; i++ {
		resultSlice[i] = rand.Intn(n * n)
	}
	return resultSlice
}

var (
	first  = generateSlice(10000)
	second = generateSlice(10000)
)

func BenchmarkGetEqualsMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetEqualsMap(first, second)
	}
}

func BenchmarkGetEqualsSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetEqualsSort(first, second)
	}
}

func TestGetEqualsMap(t *testing.T) {
	first := []int{-2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	second := []int{-2, 4, 5, 6, 7, 8, 9, 10, 11, 15, 17, 19}
	expected := []int{-2, 4, 5, 6, 7, 8, 9, 10, 11, 15, 17, 19}
	result := GetEqualsMap(first, second)
	assert.Equal(t, expected, result)
}

func TestGetEqualsSort(t *testing.T) {
	first := []int{-2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	second := []int{-2, 4, 5, 6, 7, 8, 9, 10, 11, 15, 17, 19}
	expected := []int{-2, 4, 5, 6, 7, 8, 9, 10, 11, 15, 17, 19}
	result := GetEqualsSort(first, second)
	assert.Equal(t, expected, result)
}
