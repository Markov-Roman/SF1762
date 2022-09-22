package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 50; i++ {
			ch2 <- i + 50
		}
		close(ch2)
	}()
	for {
		select {
		case <-ch1:
			fmt.Println("In 1 chan")
		case <-ch2:
			fmt.Println("In 2 chan")
		default:
			fmt.Println(time.Now())
		}
	}
}
