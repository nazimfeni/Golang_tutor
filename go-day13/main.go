package main

import "fmt"

type Employee struct {
    ID         int
    Name       string
    Department string
    Salary     float64
    Active     bool
}

func main() {

    employees := []Employee{
        {1, "Nazim", "IT", 75000, true},
        {2, "Rahim", "HR", 50000, true},
        {3, "Karim", "Finance", 65000, false},
    }

    fmt.Println("Active Employees:")
    for _, emp := range employees {
        if emp.Active {
            fmt.Printf("ID:%d Name:%s Dept:%s Salary:%.2f\n",
                emp.ID, emp.Name, emp.Department, emp.Salary)
        }
    }
}
