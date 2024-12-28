package main

import "fmt"

// send the last value + the last value, starting from 1, quit when arriving at 10

func main() {
	ch := make(chan int, 15)
	quit := make(chan int)

	go func() {
		i := 0
		for i < 10 {
			i++
			ch <- i
		}
		quit <- i
	}()

	for {
		select {
		case r := <-ch:
			fmt.Printf("%d", r)
		case <-quit:
			return
		}
	}
}
