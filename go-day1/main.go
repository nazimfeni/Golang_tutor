package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Name: Nazim")
	fmt.Println("Role: IT Head")

	// Print today's date
	today := time.Now().Format("02-01-2006")
	fmt.Println("Date:", today)
}

