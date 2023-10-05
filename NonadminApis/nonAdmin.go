package nonadmin

import api "Employee_Management_CLI/Api"

// Employee is embeded into the NonAdmin type from the api package.
type NonAdmin struct {
	api.Employee
}
