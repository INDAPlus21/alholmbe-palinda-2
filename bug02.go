package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// The problem was that the main routine finished before 11 got printed
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		wg.Add(1) // add one to the waitgroup in every iteration
		ch <- i
	}
	wg.Wait() // waits until we have done wg.Add and wg.Done on every iteration
	close(ch)
}

func Print(ch <-chan int) {
	for n := range ch {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
		wg.Done() // remove 1 from the waitgroup in every iteration
	}
}
