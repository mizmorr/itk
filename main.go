package main

import (
	"fmt"

	"javacode/task1"
	"javacode/task2"
	"javacode/task3"
	"javacode/task4"
	"javacode/task5"
	"javacode/task7"
	"javacode/task8"
	"javacode/task9"
)

func main() {
	number := 1
	fmt.Printf("running task %v..\n", number)
	task1.Task1()
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	task2.Task2()
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	customMap := task3.NewStringIntMap()
	customMap.Add("apple", 5)
	fmt.Println("Custom map:", customMap)
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	first := []int{-2, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	second := []int{-2, 4, 5, 6, 7, 8, 9}
	expected := []int{-2, 4, 5, 6, 7, 8, 9}
	result := task4.GetEqualsSort(first, second)
	fmt.Printf("Equals array:%v, expected: %v\n", result, expected)
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	expectedFoundValue := true
	found, foundValue := task5.GetEqualsSort(first, second)
	fmt.Printf("Found: %v, expectedFound: %v, foundValue: %v\n", found, expectedFoundValue, foundValue)
	number++
	fmt.Println("----------------------------------------------------------------")

	// Generating a lot values
	//	fmt.Printf("running tasks %v..\n", number)
	//	task6.Task6()
	number++
	//	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	task7.Task7()
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	s := task8.NewSemaphore()
	s.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer s.Done()
			fmt.Printf("Goroutine %v is started\n", i)
		}()
	}
	s.Wait()
	fmt.Println("All goroutines finished")
	number++
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("running task %v..\n", number)
	task9.Task9()
	fmt.Println("----------------------------------------------------------------")
}
