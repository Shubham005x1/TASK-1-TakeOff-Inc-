package api

import (
	validations "Employee_Management_CLI/Validations"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var CountEmp int = 1

// AddEmployee allows an admin to add a new employee.
func (e Employee) AddEmployee() {
	
	var employee Employee
	scanner := bufio.NewScanner(os.Stdin)
	// Increment the employee count for assigning a unique ID
	CountEmp++
	employee.ID = CountEmp

	{
	Name:
		fmt.Print("Enter first name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name) // Validation for First Name
		if err != nil {
			fmt.Println(err)
			goto Name
		} else {
			employee.FirstName = name
		}

	}
	{
	LastName:
		fmt.Print("Enter last name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name) // Validation for Last Name
		if err != nil {
			fmt.Println(err)
			goto LastName
		} else {
			employee.LastName = name

		}

	}

	{
	Email:
		fmt.Print("Enter email: ")
		scanner.Scan()
		var str string = scanner.Text()
		err := validations.IsValidEmail(str) // Validation for Email
		if err != nil {
			fmt.Println(err)
			goto Email
		} else {
			employee.Email = str
		}
	}
	{
	Password:
		fmt.Print("Enter password: ")
		scanner.Scan()
		pass := scanner.Text()
		if pass == "" {
			err := errors.New("=>Password cannot be empty please enter valid password")
			fmt.Println(err)
			goto Password
		} else {
			employee.Password = pass
		}

	}
	{
	Birthday:
		fmt.Print("Enter Birthday (YYYY-MM-DD): ")
		scanner.Scan()
		birthdayInput := scanner.Text()
		birthday, err := time.Parse("2006-01-02", birthdayInput)
		if err != nil {
			err := errors.New("=>Please enter valid value for the Birthday")
			fmt.Println(err)
			goto Birthday
		} else {
			currentdate := time.Now()
			if birthday.After(currentdate) {
				err := errors.New("=>Please enter a birthday that is not in the future")
				fmt.Println(err)
				goto Birthday
			} else {
				employee.Birthday = birthday
			}

		}

	}

	{
	Number:
		fmt.Print("Enter Phone Number: ")
		scanner.Scan()
		val := scanner.Text()
		err := validations.IsNumberValid(val) // Validation for Number
		if err != nil {
			fmt.Println(err)
			goto Number
		} else {
			employee.PhoneNo = val
		}

	}
	{
	Role:
		fmt.Print("Enter Role(manager/developer/tester/admin): ")
		scanner.Scan()
		role := scanner.Text()
		err := validations.IsValidRole(role) // Validation for Role
		if err != nil {
			fmt.Println(err)
			goto Role
		} else {
			str := strings.ToLower(role)
			employee.Role = str
		}

	}
	{
	Salary:
		fmt.Print("Enter Salary: ")
		scanner.Scan()
		sal := scanner.Text()
		fl, _ := strconv.ParseFloat(sal, 64)
		err := validations.IsValidSalary(fl) // Validation for Salary
		if err != nil {
			fmt.Println(err)
			goto Salary
		} else {
			employee.Salary = fl
		}

	}
	// Append the new employee to the list
	Emplist = append(Emplist, employee)
	fmt.Println("Employee added successfully!")
	for _, value := range Emplist {

		fmt.Printf("[Id=%d, Name=%s, LastName=%s, Email=%s, Birthday=%s, PhoneNo=%s, Role=%v, Salary=%v] \n ", value.ID, value.FirstName,
			value.LastName, value.Email, value.Birthday.Format("01-01-2006"), value.PhoneNo, value.Role, value.Salary)
	}
}
