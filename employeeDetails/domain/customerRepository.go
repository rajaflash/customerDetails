package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
	// "io"
)

type Person struct {
	Id       string `json:"employeeId"`
	Name     string `json:"employeeName"`
	Location string `json:"location"`
}

type MongoConnState struct {
	connState *mongo.Client
}

type IPostCustomerDetail interface {
	PostCustomerDetail(req []interface{}) ([]Person, error)
}
