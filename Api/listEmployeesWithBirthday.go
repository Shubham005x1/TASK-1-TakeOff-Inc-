package api

import (
	"fmt"
	"time"
)

func (e Employee) ListEmployeesWithUpcomingBirthday() {
	fmt.Println("\nEmployees with Upcoming Birthday")

	currentMonth := int(time.Now().Month())
	// Flag to check if any employees are found with upcoming birthdays
	found := false
	for _, employee := range Emplist {
		birthdayMonth := int(employee.Birthday.Month())
		// Check if the employee's birthday month matches the current month
		if birthdayMonth == currentMonth {
			fmt.Printf("%s %s\n", employee.FirstName, employee.LastName)
			fmt.Printf("Birthday: %s\n", employee.Birthday.Format("2006-01-02"))
			found = true
		}
	}
	if !found {
		fmt.Println("No employees found with upcoming birthdays this month.")
	}
}
