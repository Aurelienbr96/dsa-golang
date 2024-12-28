package main

import (
	"fmt"
	"sync"
)

func main() {
	myslice := make([]int, 15)

	var wg sync.WaitGroup
	wg.Add(len(myslice))

	for i := range myslice {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d", i)
		}(i)
	}
	wg.Wait()
}
