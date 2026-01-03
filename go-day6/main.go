package main
import "fmt"
func main() {

//Grade evaluation using switch-case
//     fmt.Println("Hello, World!")
//     mark:= 85
//     switch {
//     case mark >= 90:
// 	  fmt.Println("Grade: A")
//     case mark >= 80:
// 	  fmt.Println("Grade: B")
//     case mark >= 70:
// 	  fmt.Println("Grade: C")
//     case mark >= 60:
// 	  fmt.Println("Grade: D")
//     default:
// 	  fmt.Println("Grade: F")
// }
// Simple calculator using switch-case
var a, b int
	var choice int

	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	switch choice {
	case 1:
		fmt.Println("Result:", a+b)
	case 2:
		fmt.Println("Result:", a-b)
	case 3:
		fmt.Println("Result:", a*b)
	case 4:
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Cannot divide by zero")
		}
	default:
		fmt.Println("Invalid choice")
	}

}