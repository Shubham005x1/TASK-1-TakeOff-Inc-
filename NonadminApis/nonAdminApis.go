package nonadmin

// ViewMyDetails allows a non-admin user to view their own details.
func (e NonAdmin) ViewMyDetails(id int) {

	e.ViewEmployeeDetails(id)
}

// UpdateMyDetails allows a non-admin user to update their own details.
func (e NonAdmin) UpdateMyDetails(id int) {

	e.UpdateEmployeeDetails(id)
}

// NonAdminSearchEmployee allows a non-admin user to search for other employees.
func (e NonAdmin) NonAdminSearchEmployee(id int) {

	e.SearchEmployee(id)
}
