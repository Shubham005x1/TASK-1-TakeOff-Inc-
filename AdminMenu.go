package main

import (
	api "Employee_Management_CLI/Api"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adminMenu() {
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nAdmin Menu")
		fmt.Println("1. add employee")
		fmt.Println("2. view employee details")
		fmt.Println("3. Update employee details")
		fmt.Println("4. Delete employee")
		fmt.Println("5. list all employees")
		fmt.Println("6. list sorted employees")
		fmt.Println("7. list employees with upcoming birthday")
		fmt.Println("8. search employee")
		fmt.Println("9. logout")

		fmt.Print("Enter your choice: ")

		val, _ := reader.ReadString('\n')
		choice, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)

		switch choice {
		case 1:
			api.CurrentUser.AddEmployee()
		case 2:
			{

				fmt.Println("Give the Id of the employee you want to View")
				val, _ := reader.ReadString('\n')
				n, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				api.CurrentUser.ViewEmployeeDetails(int(n))
			}
		case 3:
			{

				fmt.Println("Give the Id of the employee you want to update")
				val, _ := reader.ReadString('\n')
				num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				api.CurrentUser.UpdateEmployeeDetails(int(num))
			}
		case 4:
			{

				fmt.Println("Give the Id of the employee you want to delete")
				val, _ := reader.ReadString('\n')
				num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				api.CurrentUser.DeleteEmployee(int(num))
			}
		case 5:
			api.CurrentUser.ListAllEmployees()
		case 6:
			api.CurrentUser.ListSortedEmployees()
		case 7:
			api.CurrentUser.ListEmployeesWithUpcomingBirthday() //remaining
		case 8:
			{

				fmt.Println("Give the Id of the employee you want to Search")
				val, _ := reader.ReadString('\n')
				num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
				api.CurrentUser.SearchEmployee(int(num))
			}
		case 9:
			api.CurrentUser = api.Employee{}
			api.Login()
			return
		default:
			fmt.Println("Please enter valid input")

		}
	}

}
