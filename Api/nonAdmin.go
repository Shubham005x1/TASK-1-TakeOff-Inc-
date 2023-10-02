package api

type NonAdmin struct {
	*Employee
}

func createNewEmployee() *NonAdmin {
	return &NonAdmin{
		Employee: &Employee{},
	}
}
