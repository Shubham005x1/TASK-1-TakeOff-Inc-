package api

import "time"

type Admin interface {
	AddEmployee()
	ViewEmployeeDetails(id int)
	UpdateEmployeeDetails(id int)
	DeleteEmployee(id int)
	ListAllEmployees()
	ListSortedEmployees()
	ListEmployeesWithUpcomingBirthday()
	SearchEmployee(id int)
}
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	PhoneNo   string
	Role      string
	Salary    float64
	Birthday  time.Time
}
type NonAdmins interface {
	ViewMyDetails(id int)
	UpdateMyDetails(id int)
	SearchEmployee(id int)
}
