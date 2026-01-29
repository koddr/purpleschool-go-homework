package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	intCh := make(chan int, 10)
	intSlice := make([]int, 10)

	wg.Go(func() {
		for range len(intSlice) {
			intCh <- rand.Intn(100)
		}
	})

	wg.Go(func() {
		for i := range len(intSlice) {
			intSlice[i] = int(math.Pow(float64(<-intCh), 2))
		}
	})

	wg.Wait()

	fmt.Println(intSlice)
}
