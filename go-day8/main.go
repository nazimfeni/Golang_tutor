package main

import "fmt"

func greet() string {
	return "Hello, World!"
}

func add(a int, b int) int {
	return a + b	
}


func main()	{
	fmt.Println(add(4,6))
}	

