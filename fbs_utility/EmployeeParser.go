package fbs_utility

import (
	"FlatBuffer-Go/Models"
	"FlatBuffer-Go/fbs_schema"
	"github.com/google/flatbuffers/go"
)

func SerializeEmployee(emp Models.Employee) []byte {
	fbsBuilder := flatbuffers.NewBuilder(0)

	fbs_firstname := fbsBuilder.CreateString(emp.FirstName)
	fbs_lastname := fbsBuilder.CreateString(emp.LastName)
	fbs_phoneno := fbsBuilder.CreateString(emp.PhoneNo)

	fbs_schema.EmployeeStart(fbsBuilder)
	fbs_schema.EmployeeAddFistName(fbsBuilder, fbs_firstname)
	fbs_schema.EmployeeAddEno(fbsBuilder,int32(emp.Eno))
	fbs_schema.EmployeeAddLastName(fbsBuilder,fbs_lastname)
	fbs_schema.EmployeeAddPhonoNo(fbsBuilder,fbs_phoneno)
	fbs_schema.EmployeeAddSalary(fbsBuilder,int32(emp.Salary))

	emp_position := fbs_schema.EmployeeEnd(fbsBuilder)
	fbsBuilder.Finish(emp_position)

	return fbsBuilder.Bytes[fbsBuilder.Head():]
}

func DeSerializeEmployee(incommingData []byte) Models.Employee {
	fbs_Employee := fbs_schema.GetRootAsEmployee(incommingData,0)
	newEmployee := *new(Models.Employee)

	newEmployee.Eno = int(fbs_Employee.Eno())
	newEmployee.FirstName = string(fbs_Employee.FistName())
	newEmployee.LastName = string(fbs_Employee.LastName())
	newEmployee.PhoneNo = string(fbs_Employee.PhonoNo())
	newEmployee.Salary = int(fbs_Employee.Salary())
	return newEmployee
}
