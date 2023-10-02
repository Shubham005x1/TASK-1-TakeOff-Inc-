package api

type EmployeeManager interface {
	AddEmployee()
	ViewEmployeeDetails(id int)
	UpdateEmployeeDetails(id int)
	DeleteEmployee(id int)
	ListAllEmployees()
	ListSortedEmployees()
	ListEmployeesWithUpcomingBirthday()
	SearchEmployee()
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
}

type AdminManager interface {
	EmployeeManager
	ListEmployeesSortedByField(string)
}

type NonAdminManager interface {
	ViewMyDetails(id int)
	//UpdateMyDetails(id int, newDetails Employee)
	//SearchEmployee(query string) Employee
}
