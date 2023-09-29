package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	PhoneNo   string
	Role      string
	Salary    float64
}

func main() {
	emplist := []Employee{}
	emplist = append(emplist, Employee{ID: 1,
		FirstName: "Shubham",
		LastName:  "Yadav",
		Email:     "shubham005@gmail.com", Password: "abc", PhoneNo: "901535548351", Role: "IT",
		Salary: 454.05})

	fmt.Print(emplist)

}
