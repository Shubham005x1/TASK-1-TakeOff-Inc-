package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CountEmp int = 1

func main() {
	InsertData()
	if CurrentUser.ID == 0 {
		Login()
	}
	switch CurrentUser.Role {
	case "admin":
		adminMenu()
	case "manager", "developer", "tester":
	}
}

func adminMenu() {
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nAdmin Menu")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employee Details")
		fmt.Println("3. Update Employee Details")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. List All Employees")
		fmt.Println("6. List Sorted Employees")
		fmt.Println("7. List Employees with Upcoming Birthday")
		fmt.Println("8. Search Employee")
		fmt.Println("9. Logout")

		fmt.Print("Enter your choice: ")

		val, _ := reader.ReadString('\n')
		choice, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)

		switch choice {
		case 1:
			CurrentUser.AddEmployee()
		case 2:
			{
				var num int
				fmt.Println("Give the Id of the employee you want to View")
				fmt.Scan(&num)
				CurrentUser.ViewEmployeeDetails(num)
			}
		case 3:
			{

				fmt.Println("Give the Id of the employee")
				val, _ := reader.ReadString('\n')
				num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				CurrentUser.UpdateEmployeeDetails(int(num))
			}
		case 4:
			{
				var num int
				fmt.Println("Give the Id of the employee you want to delete")
				fmt.Scan(&num)
				CurrentUser.DeleteEmployee(num)
			}
		case 5:
			CurrentUser.ListAllEmployees()
		case 6:
			CurrentUser.ListSortedEmployees()
		case 7:
			CurrentUser.ListEmployeesWithUpcomingBirthday()
		case 8:
			CurrentUser.SearchEmployee()
		case 9:
			CurrentUser = Employee{}
			return
		default:
			fmt.Println("Please Enter Valid Value")

		}
	}

}
