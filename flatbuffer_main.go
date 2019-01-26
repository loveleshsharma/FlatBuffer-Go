package main

import (
	"FlatBuffer-Go/Models"
	"FlatBuffer-Go/fbs_utility"
	"fmt"
)

func main() {

	employee := Models.GetNewEmployee(1,"Lovelesh","Sharma","7401094877",10000)
	flatBytes := fbs_utility.SerializeEmployee(employee)

	newEmp := fbs_utility.DeSerializeEmployee(flatBytes)
	fmt.Println(newEmp)
}
