package Models

import "fmt"

type Employee struct {
	Eno int
	FirstName string
	LastName string
	PhoneNo string
	Salary int
}

func NewEmployee() Employee {
	return *new(Employee)
}

func GetNewEmployee(eno int, firstName string,lastName string,phoneNo string, salary int) Employee {
	newEmp := NewEmployee()
	newEmp.Eno = eno
	newEmp.FirstName = firstName
	newEmp.LastName = lastName
	newEmp.PhoneNo = phoneNo
	newEmp.Salary = salary
	return newEmp
}

func (emp *Employee) GetPhoneNo() string {
 	return emp.PhoneNo
}

func (emp *Employee) GetSalary() int {
	return emp.Salary
}

func (emp *Employee) String() string {
	return fmt.Sprintf("Eno: %d, FirstName: %s, LastName: %s, PhoneNo: %s, Salary: %d",emp.Eno,emp.FirstName,emp.LastName,emp.PhoneNo,emp.Salary)
}