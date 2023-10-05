package api

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (e Employee) ListSortedEmployees() {
	var field int
	fmt.Print("\nEnter field to sort by (1: ID, 2: FirstName, 3: LastName, 4: Email, 5: Role, 6: Salary): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fieldStr := scanner.Text()

	// Convert the user input to an integer
	field, err := strconv.Atoi(fieldStr)
	if err != nil || field < 1 || field > 6 {
		fmt.Println("\nInvalid input for sorting field. Please enter a number between 1 and 6.")
		return
	}

	var lessFunc func(i, j int) bool
	switch field {
	case 1:
		lessFunc = func(i, j int) bool {
			return Emplist[i].ID < Emplist[j].ID
		}
	case 2:
		lessFunc = func(i, j int) bool {
			return Emplist[i].FirstName < Emplist[j].FirstName
		}
	case 3:
		lessFunc = func(i, j int) bool {
			return Emplist[i].LastName < Emplist[j].LastName
		}
	case 4:
		lessFunc = func(i, j int) bool {
			return Emplist[i].Email < Emplist[j].Email
		}
	case 5:
		lessFunc = func(i, j int) bool {
			return Emplist[i].Role < Emplist[j].Role
		}
	case 6:
		lessFunc = func(i, j int) bool {
			return Emplist[i].Salary < Emplist[j].Salary
		}
	default:
		fmt.Println("\nInvalid field for sorting. Please try again.")
		return
	}

	// Perform bubble sort based on the chosen comparison function
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
