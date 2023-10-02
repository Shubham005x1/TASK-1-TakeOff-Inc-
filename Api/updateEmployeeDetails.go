package api

import (
	"bufio"
	"fmt"
	"os"
)

func (e Employee) UpdateEmployeeDetails(id int) {

	var emp Employee
	in := GetIndex(id)

	scanner := bufio.NewScanner(os.Stdin)
	emp.ID = id
	fmt.Print("\nEnter New First Name: ")
	scanner.Scan()
	emp.FirstName = scanner.Text()

	fmt.Print("Enter New Last Name: ")
	scanner.Scan()
	emp.LastName = scanner.Text()

	fmt.Print("Enter New Email: ")
	scanner.Scan()
	emp.Email = scanner.Text()

	fmt.Print("Enter New Phone Number: ")
	scanner.Scan()
	emp.PhoneNo = scanner.Text()

	fmt.Print("Enter New Role: ")
	scanner.Scan()
	emp.Role = scanner.Text()

	fmt.Print("Enter New Salary: ")
	fmt.Scan(&emp.Salary)

	Emplist = append(Emplist[:in], Emplist[in+1:]...)
	Emplist = append(Emplist, emp)
	fmt.Println(Emplist)
	fmt.Println("\nEmployee details updated successfully!")

}
