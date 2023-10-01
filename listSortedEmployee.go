package main

import (
	"bufio"
	"fmt"
	"os"
)

func (e *Employee) ListSortedEmployees() {
	var field string
	fmt.Print("\nEnter field to sort by (FirstName/LastName/Email/Role/Salary): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field = scanner.Text()

	var lessFunc func(i, j int) bool

	switch field {
	case "FirstName":
		lessFunc = func(i, j int) bool {
			return Emplist[i].FirstName < Emplist[j].FirstName
		}
	case "LastName":
		lessFunc = func(i, j int) bool {
			return Emplist[i].LastName < Emplist[j].LastName
		}
	case "Email":
		lessFunc = func(i, j int) bool {
			return Emplist[i].Email < Emplist[j].Email
		}
	case "Role":
		lessFunc = func(i, j int) bool {
			return Emplist[i].Role < Emplist[j].Role
		}
	case "Salary":
		lessFunc = func(i, j int) bool {
			return Emplist[i].Salary < Emplist[j].Salary
		}
	default:
		fmt.Println("\nInvalid field for sorting. Please try again.")
		return
	}

	n := len(Emplist)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if lessFunc(j+1, j) { // If Emplist[j+1] < Emplist[j]
				// Swap Emplist[j] and Emplist[j+1]
				Emplist[j], Emplist[j+1] = Emplist[j+1], Emplist[j]
			}
		}
	}
	fmt.Println(Emplist)

}
