package domain


type Person struct {
	Id       string `json:"employeeId"`
	Name     string `json:"employeeName"`
	Location string `json:"location"`
}

type MongoConnState struct {
	connState *mongo.Client
}

type IGetCustomerDetail interface {
	GetCustomerDetail() ([]Person, error)
}


