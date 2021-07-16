package main

import (
	"fmt"
	"sync"
	"time"

	"./calculator"
)

func main() {
	appStart := time.Now()
	cache := calculator.New()

	// Define numbers to calculate
	fibonumbers := []int{50, 34, 70, 10, 40, 34}

	var wg sync.WaitGroup
	for _, n := range fibonumbers {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value := cache.Get(index)

			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)
	}

	wg.Wait()
	fmt.Printf("Everything has been done in %s\n", time.Since(appStart))
}
