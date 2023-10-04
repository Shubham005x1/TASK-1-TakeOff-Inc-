package nonadmin

func (e NonAdmin) ViewMyDetails(id int) {
	nonemp := createNewEmployee()
	nonemp.ViewEmployeeDetails(id)
}

func (e NonAdmin) UpdateMyDetails(id int) {
	nonemp := createNewEmployee()
	nonemp.UpdateEmployeeDetails(id)
}

func (e NonAdmin) SearchEmployee(id int) {
	nonemp := createNewEmployee()
	nonemp.SearchEmployee(id)
}
