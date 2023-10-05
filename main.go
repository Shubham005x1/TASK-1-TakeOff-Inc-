package main

import (
	api "Employee_Management_CLI/Api"
)

func main() {
	api.InsertData() // Load the initial employee data for the first login

	for {
		if api.CurrentUser.ID == 0 {
			api.Login() // If no user is logged in, login will be called
		}
		switch api.CurrentUser.Role {
		case "admin":
			adminMenu() // If the user is an admin, admin menu will get called
		case "manager", "developer", "tester":
			NonadminMenu() // If the user is a manager, developer, or tester, non-admin menu will get called
		}
	}
}
