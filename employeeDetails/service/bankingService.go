package service

import (
	"api/employee/employeeDetails/domain"
	"fmt"
)

// type Customer struct {
// 	Id             string
// 	POCName        string
// 	EmployeeDetail []domain.Person
// }

type Customer struct {
	Domain domain.IGetCustomerDetail
}

type ICustomer interface {
	GetCustomer() []domain.Person
}

func (c Customer) GetCustomer() []domain.Person {

	fmt.Println("Inside GetCustomer in service layer")

	repoResponse,_ := c.Domain.GetCustomerDetail()
	// requiredResponse := Customer{Id: "00001", POCName: "Charithra", EmployeeDetail: repoResponse}
	// fmt.Println("final reponse from service layer", requiredResponse)

	fmt.Println("final reponse from service layer", repoResponse)

	return repoResponse
}
