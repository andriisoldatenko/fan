package main

import "fmt"

func main() {

	ch1 := make(chan int)
	go func() {
		ch2 := make(chan int)
		go func() {
			ch3 := make(chan int)
			go func() {
				ch3 <- 1
			}()
			for {
				ch2 <- <-ch3
				ch2 <- 2
			}
		}()
		for {
			ch1 <- <-ch2
			ch1 <- 3
		}
	}()
	fmt.Println(<-ch1)
}
