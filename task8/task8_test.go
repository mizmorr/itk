package task8

import (
	"fmt"
	"testing"
)

func TestSema(t *testing.T) {
	s := NewSemaphore()
	s.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer s.Done()
			fmt.Printf("Goroutine %v is started\n", i)
		}()
	}

	s.Wait()
	t.Log("All goroutines finished")
}
