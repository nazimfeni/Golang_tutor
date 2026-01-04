package main

import "fmt"

func main() {
	todos := []string{}

	todos = append(todos, "Learn Go basics")
	todos = append(todos, "Practice slices")
	todos = append(todos, "Build a mini project")

	fmt.Println("📝 To-Do List:")

	for i, task := range todos {
		fmt.Printf("%d. %s\n", i+1, task)
	}

	fmt.Println("\nTotal tasks:", len(todos))
	fmt.Println("Capacity:", cap(todos))
}



