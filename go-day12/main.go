package main

import "fmt"

func main() {
	// User database (username : password)
	users := map[string]string{
		"admin": "1234",
		"nazim": "pass@123",
		"john":  "abcd",
	}

	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	// Check user existence
	storedPassword, exists := users[username]

	if !exists {
		fmt.Println("❌ User does not exist")
		return
	}

	// Check password
	if storedPassword == password {
		fmt.Println("✅ Login successful")
	} else {
		fmt.Println("❌ Invalid password")
	}
}

