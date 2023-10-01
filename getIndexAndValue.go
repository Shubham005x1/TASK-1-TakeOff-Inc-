package main

func GetVal(id int) (e Employee) {
	for _, value := range Emplist {
		if value.ID == id {
			return value
		}

	}
	return Employee{}
}
func GetIndex(id int) int {

	for i, value := range Emplist {
		if value.ID == id {
			return i
		}

	}
	return -1
}
