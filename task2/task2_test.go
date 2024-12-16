package task2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceExample(t *testing.T) {
	originalSlice := getOriginalSlice()
	newSlice := sliceExample(originalSlice)
	if len(newSlice) == 0 {
		t.Error("Expected non-empty slice, got empty slice")
	}
}

func BenchmarkSliceExample(b *testing.B) {
	originalSlice := getOriginalSlice()
	for n := 0; n < b.N; n++ {
		sliceExample(originalSlice)
	}
}

func BenchmarkSliceExample2(b *testing.B) {
	originalSlice := getOriginalSlice()
	for n := 0; n < b.N; n++ {
		sliceExample2(originalSlice)
	}
}

func TestAddElement(t *testing.T) {
	originalSlice := getOriginalSlice()
	newSlice := addElements(originalSlice, 1000)

	assert.Equal(t, newSlice[len(newSlice)-1], 1000)
	assert.NotEqual(t, len(newSlice), len(originalSlice))

	if &newSlice[0] == &originalSlice[0] {
		t.Error("Expected new slice to be a copy of the original slice")
	}
}

func TestRemoveElement(t *testing.T) {
	originalSlice := getOriginalSlice()
	elemFifth := originalSlice[5]

	newSlice := removeElement(originalSlice, 5)

	assert.Equal(t, len(newSlice), len(originalSlice)-1)
	assert.NotEqual(t, elemFifth, newSlice[5])
	assert.Equal(t, &originalSlice[0], &newSlice[0])
}

func TestRemoveElementLast(t *testing.T) {
	originalSlice := getOriginalSlice()
	newSlice := removeElement(originalSlice, len(originalSlice)-1)

	assert.Equal(t, len(newSlice), len(originalSlice)-1)
	assert.NotEqual(t, originalSlice[len(originalSlice)-1], newSlice[len(newSlice)-1])
	assert.Equal(t, &originalSlice[0], &newSlice[0])
}

func TestRemoveElementFirst(t *testing.T) {
	originalSlice := getOriginalSlice()
	newSlice := removeElement(originalSlice, 0)

	assert.Equal(t, len(newSlice), len(originalSlice)-1)
	assert.NotEqual(t, originalSlice[0], newSlice[0])
	assert.Equal(t, &originalSlice[1], &newSlice[0])
}

func BenchmarkAddElement(b *testing.B) {
	originalSlice := getOriginalSlice()
	for n := 0; n < b.N; n++ {
		addElements(originalSlice, 1000)
	}
}

func TestTask2(t *testing.T) {
	Task2()
}
