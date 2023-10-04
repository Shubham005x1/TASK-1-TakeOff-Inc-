package main

import (
	api "Employee_Management_CLI/Api"
	nonadmin "Employee_Management_CLI/NonadminApis"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NonadminMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nNonAdmin Menu")
	fmt.Println("1.View my details")
	fmt.Println("2.Update my details")
	fmt.Println("3.Search other employee")
	fmt.Println("4.logout")
	val, _ := reader.ReadString('\n')
	choice, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)

	switch choice {

	case 1:
		var num int = api.CurrentUser.ID
		n := nonadmin.NonAdmin{}
		n.ViewMyDetails(num)

	case 2:
		var num int = api.CurrentUser.ID
		n := nonadmin.NonAdmin{}
		n.UpdateMyDetails(num)

	case 3:
		{
			fmt.Println("Give the Id of the employee you want to Search")
			val, _ := reader.ReadString('\n')
			num, _ := strconv.ParseInt(strings.TrimSpace(val), 36, 64)
			api.CurrentUser.SearchEmployee(int(num))
		}
	case 4:
		api.CurrentUser = api.Employee{}
		api.Login()
		return
	default:
		fmt.Println("Please enter valid input")
	}
}
