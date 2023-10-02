package api

import (
	"fmt"
)

func (e Employee) ListEmployeesWithUpcomingBirthday() {
	fmt.Println("\nEmployees with Upcoming Birthday")
	//	currentMonth := int(time.Now().Month())
	for _, employee := range Emplist {
		//birthdayMonth := int(time.Date(time.Now().Year(), employee.Birthday.Month(), employee.Birthday.Day(), 0, 0, 0, 0, time.UTC).Month())
		//if birthdayMonth == currentMonth {
		fmt.Println(employee)
	}
}
