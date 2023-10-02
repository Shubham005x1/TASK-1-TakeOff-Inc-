package api

import (
	validations "Employee_Management_CLI/Validations"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func generateId() int {
// 	var randomNumber int
// 	rand.Seed(time.Now().UnixNano())
// 	randomNumber = rand.Intn(50)
// 	return randomNumber

// }
var CountEmp int = 1

func (e Employee) AddEmployee() {
	var employee Employee
	scanner := bufio.NewScanner(os.Stdin)
	CountEmp = CountEmp + 1
	employee.ID = CountEmp

	{
	Name:
		fmt.Print("Enter First Name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name)
		if err != nil {
			fmt.Println(err)
			goto Name
		} else {
			employee.FirstName = name
		}

	}
	{
	LastName:
		fmt.Print("Enter Last Name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name)
		if err != nil {
			fmt.Println(err)
			goto LastName
		} else {
			employee.LastName = name

		}

	}

	{
	Email:
		fmt.Print("Enter Email: ")
		scanner.Scan()
		var str string = scanner.Text()
		err := validations.IsValidEmail(str)
		if err != nil {
			fmt.Println(err)
			goto Email
		} else {
			employee.Email = str
		}
	}

	fmt.Print("Enter Password: ")
	scanner.Scan()
	employee.Password = scanner.Text()

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	employee.PhoneNo = scanner.Text()

	{
	Role:
		fmt.Print("Enter Role(manager/developer/tester/admin): ")
		scanner.Scan()
		role := scanner.Text()
		err := validations.IsValidRole(role)
		if err != nil {
			fmt.Println(err)
			goto Role
		} else {
			employee.Role = role
		}

	}
	{
	Salary:
		fmt.Print("Enter Salary: ")
		scanner.Scan()
		sal := scanner.Text()
		fl, _ := strconv.ParseFloat(sal, 64)
		err := validations.IsValidSalary(fl)
		if err != nil {
			fmt.Println(err)
			goto Salary
		} else {
			employee.Salary = fl
		}

	}

	Emplist = append(Emplist, employee)
	fmt.Println("Employee added successfully!")
	fmt.Println(Emplist)
}
