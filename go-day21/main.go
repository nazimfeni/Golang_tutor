package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

var fileName = "expenses.json"

func loadExpenses() []Expense {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Expense{}
	}

	var expenses []Expense
	json.Unmarshal(data, &expenses)
	return expenses
}

func saveExpenses(expenses []Expense) {
	data, _ := json.MarshalIndent(expenses, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func addExpense() {
	var title string
	var amount float64

	fmt.Print("Enter title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	expenses := loadExpenses()

	newExpense := Expense{
		ID:     len(expenses) + 1,
		Title:  title,
		Amount: amount,
		Date:   time.Now().Format("2006-01-02"),
	}

	expenses = append(expenses, newExpense)
	saveExpenses(expenses)

	fmt.Println("✅ Expense added successfully")
}

func viewExpenses() {
	expenses := loadExpenses()

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	fmt.Println("\n--- Expense List ---")
	for _, e := range expenses {
		fmt.Printf(
			"ID: %d | %s | %.2f | %s\n",
			e.ID, e.Title, e.Amount, e.Date,
		)
	}
}

func totalExpense() {
	expenses := loadExpenses()
	total := 0.0

	for _, e := range expenses {
		total += e.Amount
	}

	fmt.Printf("\nTotal Expense: %.2f\n", total)
}

func main() {
	for {
		fmt.Println("\n===== Expense Tracker =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Total Expense")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Choose option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addExpense()
		case 2:
			viewExpenses()
		case 3:
			totalExpense()
		case 4:
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
