package task6

import (
	"testing"
)

func TestTask(t *testing.T) {
	Task6()
}

func TestLCG(t *testing.T) {
	lcg := newLCG(42, 1103515245, 12345, 1<<31)
	generatedValues := make(map[int64]struct{})
	for i := 0; i < 1000; i++ {
		value := lcg.next()
		if _, ok := generatedValues[value]; ok {
			t.Error("Duplicate value generated")
		}
		generatedValues[value] = struct{}{}
	}
}
