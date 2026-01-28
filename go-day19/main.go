package main

import (
	"fmt"
	"os"
	"time"
)

func writeLog(message string) {
	file, err := os.OpenFile(
		"app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	log := fmt.Sprintf(
		"%s - %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
	)

	file.WriteString(log)
}

func readLogs() {
	data, err := os.ReadFile("app.log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\n--- Logs ---")
	fmt.Println(string(data))
}

func main() {
	writeLog("Application started")
	writeLog("User logged in")
	writeLog("Application closed")

	readLogs()
}

