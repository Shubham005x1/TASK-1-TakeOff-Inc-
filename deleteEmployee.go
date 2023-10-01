package main

import "fmt"

func (e *Employee) DeleteEmployee(id int) {
	index := GetIndex(id)
	Emplist = append(Emplist[:index], Emplist[index+1:]...)
	fmt.Printf("Employee with ID %d is Deleted ", id)

}
