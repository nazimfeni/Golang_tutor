package main

import "fmt"

func main() {
	var name string
	var marks int

	fmt.Print("Enter student name: ")
	fmt.Scan(&name)

	fmt.Print("Enter marks: ")
	fmt.Scan(&marks)

	fmt.Println("\n--- Result ---")
	fmt.Println("Name:", name)
	fmt.Println("Marks:", marks)

	switch {
	case marks >= 80 && marks <= 100:
		fmt.Println("Grade: A+")
	case marks >= 70:
		fmt.Println("Grade: A")
	case marks >= 60:
		fmt.Println("Grade: A-")
	case marks >= 50:
		fmt.Println("Grade: B")
	case marks >= 40:
		fmt.Println("Grade: Pass")
	case marks >= 0:
		fmt.Println("Grade: Fail")
	default:
		fmt.Println("Invalid marks")
	}
}
