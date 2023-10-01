package main

import (
	"bufio"
	"fmt"
	"os"
)

// func generateId() int {
// 	var randomNumber int
// 	rand.Seed(time.Now().UnixNano())
// 	randomNumber = rand.Intn(50)
// 	return randomNumber

// }

func (e *Employee) AddEmployee() {
	var employee Employee
	scanner := bufio.NewScanner(os.Stdin)
	CountEmp = CountEmp + 1
	employee.ID = CountEmp

	fmt.Print("Enter First Name: ")
	scanner.Scan()
	employee.FirstName = scanner.Text()

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	employee.LastName = scanner.Text()

	fmt.Print("Enter Email: ")
	scanner.Scan()
	employee.Email = scanner.Text()

	fmt.Print("Enter Password: ")
	scanner.Scan()
	employee.Password = scanner.Text()

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	employee.PhoneNo = scanner.Text()

	fmt.Print("Enter Role: ")
	scanner.Scan()
	employee.Role = scanner.Text()

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	Emplist = append(Emplist, employee)
	fmt.Println("Employee added successfully!")
	fmt.Println(Emplist)
}
