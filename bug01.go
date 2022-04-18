package main

import "fmt"

// previously, "ch <- "Hello world!"" was blocking, and we
// solve it by moving it out to a separate goroutine

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
