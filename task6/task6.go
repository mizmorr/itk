package task6

import (
	"fmt"
	"math/rand"
	"time"
)

const timeToStop = 10

type lcg struct {
	seed       int64
	multiplier int64
	increment  int64
	modulus    int64
}

func newLCG(seed, multiplier, increment, modulus int64) *lcg {
	return &lcg{
		seed:       seed,
		multiplier: multiplier,
		increment:  increment,
		modulus:    modulus,
	}
}

func (lcg *lcg) next() int64 {
	lcg.seed = (lcg.multiplier*lcg.seed + lcg.increment) % lcg.modulus
	var additionalMultiplier int64 = 1
	if rand.Intn(2) == 0 {
		additionalMultiplier = -1
	}
	return lcg.seed * additionalMultiplier
}

func posix() *lcg {
	return newLCG(
		42,
		1103515245,
		12345,
		1<<31,
	)
}

func Task6() {
	values := make(chan int64)
	done := make(chan struct{})
	generator := posix()
	go func() {
		for {
			select {
			case <-done:
				close(values)
				return
			default:
				values <- generator.next()
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
