package api

import (
	validations "Employee_Management_CLI/Validations"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (e Employee) UpdateEmployeeDetails(id int) {

	var emp Employee
	in := GetIndex(id)
	scanner := bufio.NewScanner(os.Stdin)
	{
	Name:
		emp.ID = id
		fmt.Print("\nenter new first name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name)
		if err != nil {
			fmt.Println(err)
			goto Name
		} else {
			emp.FirstName = name
		}
	}
	// emp.FirstName = scanner.Text()
	{
	LastName:
		fmt.Print("enter new last name: ")
		scanner.Scan()
		name := scanner.Text()
		err := validations.IsValidEntry(name)
		if err != nil {
			fmt.Println(err)
			goto LastName
		} else {
			emp.LastName = name

		}

	}
	{
	Email:
		fmt.Print("enter new email: ")
		scanner.Scan()
		var str string = scanner.Text()
		err := validations.IsValidEmail(str)
		if err != nil {
			fmt.Println(err)
			goto Email
		} else {
			emp.Email = str
		}
	}
	{
		fmt.Print("Enter new password: ")
		scanner.Scan()
		emp.Password = scanner.Text()
	}
	{
		fmt.Print("enter new phone number: ")
		scanner.Scan()
		emp.PhoneNo = scanner.Text()
	}
	{
	Role:
		fmt.Print("Enter new role(manager/developer/tester/admin): ")
		scanner.Scan()
		role := scanner.Text()
		err := validations.IsValidRole(role)
		if err != nil {
			fmt.Println(err)
			goto Role
		} else {
			str := strings.ToLower(role)
			emp.Role = str
		}

	}
	{
	Salary:
		fmt.Print("enter new salary: ")
		scanner.Scan()
		sal := scanner.Text()
		fl, _ := strconv.ParseFloat(sal, 64)
		err := validations.IsValidSalary(fl)
		if err != nil {
			fmt.Println(err)
			goto Salary
		} else {
			emp.Salary = fl
		}

	}

	Emplist = append(Emplist[:in], Emplist[in+1:]...)
	Emplist = append(Emplist, emp)
	fmt.Println("\nEmployee details updated successfully!")

}
