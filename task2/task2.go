package task2

import (
	"fmt"
	"math/rand"
)

func getOriginalSlice() []int {
	originalSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(1000)+1)
	}
	return originalSlice
}

func sliceExample(originalSlice []int) []int {
	result := make([]int, 0)
	for _, v := range originalSlice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(originalSlice []int, newValue int) []int {
	newSlice := make([]int, len(originalSlice)+1)

	copy(newSlice, originalSlice)

	newSlice[len(originalSlice)] = newValue

	return newSlice
}

func copySlice(original []int) []int {
	copySlice := make([]int, len(original))

	copy(copySlice, original)

	return copySlice
}

func removeElement(slice []int, index int) []int {
	if index >= len(slice) {
		return slice
	}
	if index == 0 {
		return slice[1:]
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func Task2() {
	originalSlice := getOriginalSlice()
	newSlice := sliceExample(originalSlice)
	newSlice = addElements(newSlice, 1000)
	newSlice = removeElement(newSlice, 5)

	fmt.Println("Original slice:", originalSlice)

	fmt.Println("New slice with evens after adding 1000 and removing the fifth element:", newSlice)
}

func sliceExample2(originalSlice []int) []int {
	newSlice := make([]int, 0, len(originalSlice))
	var doneCount int
	done := make(chan interface{})
	oddElements := make(chan int)
	go func(slice []int) {
		defer func() { done <- struct{}{} }()
		for _, v := range slice {
			if v%2 == 0 {
				oddElements <- v
			}
		}
	}(originalSlice[:len(originalSlice)/2])
	go func(slice []int) {
		defer func() { done <- struct{}{} }()
		for _, v := range slice {
			if v%2 == 0 {
				oddElements <- v
			}
		}
	}(originalSlice[len(originalSlice)/2:])

	for {
		select {
		case odd := <-oddElements:
			newSlice = append(newSlice, odd)

		case <-done:

			doneCount++
			if doneCount == 2 {
				close(oddElements)

				return originalSlice
			}
		}
	}
}
