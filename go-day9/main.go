// package main
// import "fmt"
// // Function that returns multiple values: sum and product of two integers

// func calculate(a, b int) (int, int) {
//     return a + b, a * b
// }
// // Function that uses named return value to calculate area of a rectangle
// func rectangle(length, width int) (area int) {
//     area = length * width
//     return
// }

// func main() {
//     sum, mul := calculate(4, 5)
//     fmt.Println("Sum:", sum)
//     fmt.Println("Multiply:", mul)
// 	area := rectangle(4, 5)
// 	fmt.Println("Area of rectangle:", area)
// }

// error handling with multiple return values
package main

import (
	"errors"
	"fmt"
)

func getNumber(flag bool) (int, error) {
    if !flag {
        return 0, errors.New("invalid condition")
    }
    return 42, nil
}

func main() {
    num, err := getNumber(false)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Number:", num)
}