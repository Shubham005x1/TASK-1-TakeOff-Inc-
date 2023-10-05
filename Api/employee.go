package api

import "time"

// Admin interface defines the methods that an admin can perform.
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

// NonAdmin interface defines the methods that a non-admin can perform.
type NonAdmin interface {
	ViewMyDetails(id int)
	UpdateMyDetails(id int)
	SearchEmployee(id int)
}
