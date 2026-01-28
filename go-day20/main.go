package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	user := User{
		ID:    1,
		Name:  "Nazim",
		Email: "nazim@gmail.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

