package nonadmin

import api "Employee_Management_CLI/Api"

type NonAdmin struct {
	api.Employee
}

func createNewEmployee() *NonAdmin {
	return &NonAdmin{
		Employee: api.Employee{},
	}
}
