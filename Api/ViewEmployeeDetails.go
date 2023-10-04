package api

import "fmt"

func (e Employee) ViewEmployeeDetails(id int) {
	for _, value := range Emplist {
		if value.ID == id {
			fmt.Printf("[Id=%d, Name=%s, LastName=%s, Email=%s, Birthday=%s, PhoneNo=%s, Role=%v, Salary=%v] \n ", value.ID, value.FirstName,
				value.LastName, value.Email, value.Birthday.Format("01-01-2006"), value.PhoneNo, value.Role, value.Salary)
		}
	}
}
