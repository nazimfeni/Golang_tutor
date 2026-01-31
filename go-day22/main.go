package main

import (
	"fmt"
	"time"
)

func printNumbers(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go printNumbers(1, 5)
	go printNumbers(6, 10)

	// wait for goroutines to finish
	time.Sleep(3 * time.Second)
	fmt.Println("Done printing numbers")
}



