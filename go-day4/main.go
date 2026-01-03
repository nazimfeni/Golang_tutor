package main

import "fmt"

func main() {
	fmt.Print("Enter three integers: ")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	
	if a>b && a>c{
		fmt.Printf("%d is the largest number",a)
	} else if b>a && b>c{
		fmt.Printf("%d is the largest number",b)
	} else {
		fmt.Printf("%d is the largest number",c)
	}


}

