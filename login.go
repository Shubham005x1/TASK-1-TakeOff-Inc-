package main

import (
	"bufio"
	"fmt"
	"os"
)

var Emplist []Employee
var CurrentUser Employee

func InsertData() {
	Emplist = append(Emplist, Employee{ID: 1,
		FirstName: "Shubham",
		LastName:  "Yadav",
		Email:     "123", Password: "123", PhoneNo: "901535548351", Role: "admin",
		Salary: 454.05})
}

func Login() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEmployee Management System")
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	for _, employee := range Emplist {
		if employee.Email == username && employee.Password == password {
			CurrentUser = employee
			fmt.Printf("\nWelcome, %s!\n", CurrentUser.FirstName)

			return
		}
	}

	fmt.Println("\nInvalid username or password. Please try again.")
}
