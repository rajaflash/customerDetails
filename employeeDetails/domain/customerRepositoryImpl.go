package domain

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//calling mongoDb for testing purpose
func (p Person) GetCustomerDetail() ([]Person, error) {
	stubValue := []Person{{"1", "Raj", "Kumbakonam"}, {"1", "Arun", "cuddalore"}}
	fmt.Println("The stubbed value in repository", stubValue)

	connState := GetMongoDbConnection()
	testMongoDb := MongoConnState.PostCustomer(connState)
	fmt.Println("connection tested", testMongoDb)

	return stubValue, nil
}

func (m MongoConnState) PostCustomer() string {
	collection := m.connState.Database(os.Getenv("golangDb")).Collection(os.Getenv("customer"))
	fmt.Println("ping check", collection)
	return "hey!!!"
}

func GetMongoDbConnection() MongoConnState {
	uri := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("MongoDb Connection error", err)
	}

	return MongoConnState{client}
}