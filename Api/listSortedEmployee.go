package api

import (
	"bufio"
	"fmt"
	"os"
)

func (e Employee) ListSortedEmployees() {
	var field string
	fmt.Print("\nEnter field to sort by (ID/FirstName/LastName/Email/Role/Salary): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field = scanner.Text()

	var lessFunc func(i, j int) bool

	switch field {
	case "ID":
		lessFunc = func(i, j int) bool {
			return Emplist[i].ID < Emplist[j].ID
		}
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
	for _, value := range Emplist {

		fmt.Printf("[Id=%d, Name=%s, LastName=%s, Email=%s, Birthday=%s, PhoneNo=%s, Role=%v, Salary=%v] \n ", value.ID, value.FirstName,
			value.LastName, value.Email, value.Birthday.Format("01-01-2006"), value.PhoneNo, value.Role, value.Salary)
	}

}
