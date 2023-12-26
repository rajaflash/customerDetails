package domain

import "fmt"

type Person struct {
	Id      string `json:"employeeId"`
	Name    string `json:"employeeName"`
	Location string `json:"location"`
}

type IGetCustomerDetail interface {
	GetCustomerDetail() ([]Person, error)
}

func (p Person) GetCustomerDetail() ([]Person, error) {
	stubValue := []Person{{"1", "Raj", "Kumbakonam"}, {"1", "Arun", "cuddalore"}}
	fmt.Println("The stubbed value in repository", stubValue)
	return stubValue, nil
}
