package main

import (
	api "Employee_Management_CLI/Api"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			api.CurrentUser.AddEmployee()
		case 2:
			{
				var num int
				fmt.Println("Give the Id of the employee you want to View")
				fmt.Scan(&num)
				api.CurrentUser.ViewEmployeeDetails(num)
			}
		case 3:
			{

				fmt.Println("Give the Id of the employee")
				val, _ := reader.ReadString('\n')
				num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				api.CurrentUser.UpdateEmployeeDetails(int(num))
			}
		case 4:
			{
				var num int
				fmt.Println("Give the Id of the employee you want to delete")
				fmt.Scan(&num)
				api.CurrentUser.DeleteEmployee(num)
			}
		case 5:
			api.CurrentUser.ListAllEmployees()
		case 6:
			api.CurrentUser.ListSortedEmployees()
		case 7:
			api.CurrentUser.ListEmployeesWithUpcomingBirthday()
		case 8:
			api.CurrentUser.SearchEmployee()
		case 9:
			api.CurrentUser = api.Employee{}
			api.Login()
			return
		default:
			fmt.Println("Please Enter Valid Value")

		}
	}

}
func NonadminMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nNonAdmin Menu")
	fmt.Println("1.View Employee Details")
	fmt.Println("2.Update Your Details")
	fmt.Println("3.Search Employee")
	fmt.Println("4.Logout")
	val, _ := reader.ReadString('\n')
	choice, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)

	switch choice {

	case 1:
		var num int
		fmt.Println("Give the Id of the employee you want to delete")
		fmt.Scan(&num)
		api.CurrentUser.ViewMyDetails(num)

	case 2:

	case 3:
	case 4:
		api.CurrentUser = api.Employee{}
		api.Login()
		return
	}
}
