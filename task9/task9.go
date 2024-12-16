package task9

import (
	"fmt"
	"math"
	"math/rand"
)

const countValues = 50

func Task9() {
	chfirst := make(chan uint8)
	chsecond := make(chan float64)
	go write(chfirst)
	go func() {
		defer close(chsecond)
		for val := range chfirst {
			chsecond <- math.Pow(float64(val), 3)
		}
	}()
	for val := range chsecond {
		fmt.Println(val)
	}
}

func write(ch chan uint8) {
	defer close(ch)
	for i := 0; i < countValues; i++ {
		ch <- uint8(rand.Intn(256))
	}
}
