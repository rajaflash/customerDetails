package service

import (
	"api/employee/employeeDetails/domain"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Customer struct {
	Domain domain.IPostCustomerDetail
}

type ICustomer interface {
	PostCustomer(req io.ReadCloser) []domain.Person
}

func (c Customer) PostCustomer(req io.ReadCloser) []domain.Person {
	fmt.Println("Inside GetCustomer in service layer")

	body, err := io.ReadAll(req)
	if err != nil {
		fmt.Println("Error while reading request Body", err)
	}

	// string keys, untyped values
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while converting into Json format", err)
	}
	// Accessing the data from the request body
	fmt.Println("Received data:", data)

	//data.([]interface{}) is a type assertion
	dataArr := data["data"].([]interface{})

	var bsonDocuments []interface{}
	for _, element := range dataArr {
		// Ensure element is a map[string]interface{}
		dataMap := element.(map[string]interface{})
		bsonDoc := bson.D{{Key: "employeeId", Value: dataMap["employeeId"]}, {Key: "employeeName", Value: dataMap["employKey: eeName"]}, {Key: "location", Value: dataMap["location"]}}
		bsonDocuments = append(bsonDocuments, bsonDoc)
	}

	repoResponse, _ := c.Domain.PostCustomerDetail(bsonDocuments)
	fmt.Println("final reponse from service layer", repoResponse)
	return repoResponse
}
