package task7

import (
	"fmt"
	"sync"
)

const (
	N          = 7
	ChannelLen = 4
)

func writeFromChToCh(sender <-chan int, receiver chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range sender {
		receiver <- v
	}
}

func generateSender() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < ChannelLen; i++ {
			ch <- i
		}
	}()
	return ch
}

func Task7() {
	stream := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		ch := generateSender()
		go writeFromChToCh(ch, stream, &wg)
	}

	go func() {
		wg.Wait()
		close(stream)
	}()
	for v := range stream {
		fmt.Println(v)
	}
}
