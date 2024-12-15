package task5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEqualsFound(t *testing.T) {
	first := []int{1, 2, 3, 4, 5}
	second := []int{5, 4, 3, 2, 1}

	found, equals := GetEqualsSort(first, second)

	assert.Equal(t, found, true)
	assert.Equal(t, equals, []int{1, 2, 3, 4, 5})
}

func TestGetEqualsNotFound(t *testing.T) {
	first := []int{1, 2, 3, 4, 5}
	second := []int{6, 7, 8, 9, 10}

	found, equals := GetEqualsSort(first, second)

	assert.Equal(t, found, false)
	assert.Equal(t, equals, []int{})
}
