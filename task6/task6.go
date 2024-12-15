package task6

import (
	"fmt"
	"math/rand"
	"time"
)

const timeToStop = 10

func Task6() {
	values := make(chan int)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				close(values)
				return
			default:
				values <- rand.Intn(1000 + 1)
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * timeToStop)
		done <- struct{}{}
	}()
	for value := range values {
		fmt.Println(value)
	}
}
