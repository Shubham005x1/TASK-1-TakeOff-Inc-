package main

import "fmt"

func (e *Employee) ViewEmployeeDetails(id int) {
	for _, value := range Emplist {
		if value.ID == id {
			fmt.Printf("[Id=%d, Name=%s, LastName=%s, Email=%s, PhoneNo=%s, Role=%v, Salary=%v]  ", value.ID, value.FirstName,
				value.LastName, value.Email, value.PhoneNo, value.Role, value.Salary)
		}
	}
}
