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

// func generateId() int {
// 	var randomNumber int
// 	rand.Seed(time.Now().UnixNano())
// 	randomNumber = rand.Intn(50)
// 	return randomNumber

// }
var CountEmp int = len(Emplist)

func (e Employee) AddEmployee() {
	var employee Employee
	scanner := bufio.NewScanner(os.Stdin)
	CountEmp = CountEmp + 1
	employee.ID = CountEmp

	{
	Name:
		fmt.Print("Enter first name: ")
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
		fmt.Print("Enter last name: ")
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
		fmt.Print("Enter email: ")
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
			employee.Birthday = birthday
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
	Number:
		fmt.Print("Enter Phone Number: ")
		scanner.Scan()
		val := scanner.Text()
		err := validations.IsNumberValid(val)
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
		err := validations.IsValidRole(role)
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
	for _, value := range Emplist {

		fmt.Printf("[Id=%d, Name=%s, LastName=%s, Email=%s, Birthday=%s, PhoneNo=%s, Role=%v, Salary=%v] \n ", value.ID, value.FirstName,
			value.LastName, value.Email, value.Birthday.Format("01-01-2006"), value.PhoneNo, value.Role, value.Salary)
	}
}
