package main

import (
	"fmt"
	"time"
)

func cnt(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go cnt(bus)
	for count := range bus {
		fmt.Println("count: ", count)
	}
}
