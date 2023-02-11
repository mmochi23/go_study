package main

import (
	"fmt"
	"time"
)

func receive(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func nreceive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch1 := make(chan int)

	//go receive(ch1)
	//ch1 <- 3
	//ch1 <- 1
	close(ch1)

	ch2 := make(chan int, 20)
	go nreceive("1st goroutine", ch2)
	go nreceive("2st goroutine", ch2)
	go nreceive("3st goroutine", ch2)
	i := 0
	for i < 10 {
		ch2 <- i
		i++
	}
	close(ch2)
	time.Sleep((3 * time.Second))
}
