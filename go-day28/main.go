package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Global DB Variable
var db *sql.DB

// --------------------
// Database Connection
// --------------------
func connectDB() {

	dsn := "root:root@tcp(127.0.0.1:3306)/golang_db"

	var err error
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Database Connected!")
}

// --------------------
// Create User
// --------------------
func createUser(name, email string) {

	query := "INSERT INTO users(name, email) VALUES(?, ?)"

	result, err := db.Exec(query, name, email)

	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()

	fmt.Println("✅ User Inserted ID:", id)
}

// --------------------
// Read Users
// --------------------
func getUsers() {

	rows, err := db.Query("SELECT id, name, email FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	fmt.Println("\n📌 User List:")

	for rows.Next() {

		var id int
		var name, email string

		err := rows.Scan(&id, &name, &email)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, email)
	}
}

// --------------------
// Update User
// --------------------
func updateUser(id int, name string) {

	query := "UPDATE users SET name=? WHERE id=?"

	_, err := db.Exec(query, name, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ User Updated")
}

// --------------------
// Delete User
// --------------------
func deleteUser(id int) {

	query := "DELETE FROM users WHERE id=?"

	_, err := db.Exec(query, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ User Deleted")
}

// --------------------
// Main Function
// --------------------
func main() {

	connectDB()

	createUser("Nazim", "nazim@gmail.com")
	createUser("Rahim", "rahim@gmail.com")

	getUsers()

	updateUser(1, "Nazim Ahmed")

	deleteUser(2)

	getUsers()

	// Close DB when program ends
	defer db.Close()
}
