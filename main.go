package main

import (
	api "Employee_Management_CLI/Api"
)

func main() {
	api.InsertData()
	for {
		if api.CurrentUser.ID == 0 {
			api.Login()
		}
		switch api.CurrentUser.Role {
		case "admin":
			adminMenu()
		case "manager", "developer", "tester":
			NonadminMenu()
		}
	}
}
