package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("Producer-Consumer finished")
}
