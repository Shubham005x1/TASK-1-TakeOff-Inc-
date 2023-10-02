package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (e Employee) SearchEmployee() {
	var query string
	fmt.Print("\nEnter search query: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	query = scanner.Text()

	for _, employee := range Emplist {
		if strings.Contains(strings.ToLower(employee.FirstName), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.LastName), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.Email), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.Role), strings.ToLower(query)) {
			fmt.Println(employee)
		}
	}

}
