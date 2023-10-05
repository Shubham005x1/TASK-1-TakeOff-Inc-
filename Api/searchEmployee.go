package api

import (
	"fmt"
)

func (e Employee) SearchEmployee(id int) {
	flag := false
	for _, value := range Emplist {
		if value.ID == id {
			fmt.Printf("[Name=%s, LastName=%s, Email=%s, PhoneNo=%s, Role=%v]  ", value.FirstName,
				value.LastName, value.Email, value.PhoneNo, value.Role)
			flag = true
		}
	}
	if !flag {
		fmt.Println("=>Sorry no employee was found!! ")
	}
}
