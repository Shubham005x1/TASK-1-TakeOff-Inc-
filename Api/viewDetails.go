package api

func (e Employee) ViewMyDetails(id int) {
	nonemp := createNewEmployee()
	nonemp.ViewEmployeeDetails(id)
}
