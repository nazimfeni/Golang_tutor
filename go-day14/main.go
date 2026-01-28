package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Phone string
	Email string
}

var contacts []Contact
var nextID = 1
var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		showMenu()

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			addContact()
		case 2:
			viewContacts()
		case 3:
			searchContact()
		case 4:
			deleteContact()
		case 5:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\n--- Contact Management System ---")
	fmt.Println("1. Add Contact")
	fmt.Println("2. View Contacts")
	fmt.Println("3. Search Contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Exit")
	fmt.Print("Enter choice: ")
}

func addContact() {
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter Phone: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')

	contact := Contact{
		ID:    nextID,
		Name:  strings.TrimSpace(name),
		Phone: strings.TrimSpace(phone),
		Email: strings.TrimSpace(email),
	}

	contacts = append(contacts, contact)
	nextID++

	fmt.Println("Contact added successfully!")
}

func viewContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, c := range contacts {
		fmt.Printf("ID:%d Name:%s Phone:%s Email:%s\n",
			c.ID, c.Name, c.Phone, c.Email)
	}
}

func searchContact() {
	fmt.Print("Enter name to search: ")
	searchName, _ := reader.ReadString('\n')
	searchName = strings.TrimSpace(searchName)

	for _, c := range contacts {
		if strings.EqualFold(c.Name, searchName) {
			fmt.Printf("Found → ID:%d Name:%s Phone:%s Email:%s\n",
				c.ID, c.Name, c.Phone, c.Email)
			return
		}
	}

	fmt.Println("Contact not found.")
}

func deleteContact() {
	fmt.Print("Enter Contact ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	for i, c := range contacts {
		if c.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Println("Contact deleted successfully!")
			return
		}
	}

	fmt.Println("Contact ID not found.")
}
