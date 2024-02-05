package domain

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "io"
	"log"
	"os"
)

func (m MongoConnState) PostCustomerDetail(req []interface{}) ([]Person, error) {
	fmt.Println("Database")
	collection := m.connState.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	fmt.Println("ping check", collection)
	res, err := collection.InsertMany(context.Background(), req)
	if err != nil {
		fmt.Println("Error while inserting into mongoDB", err)
	}
	fmt.Println("Response from posting to database", res)
	return []Person{{"1", "Raj", "Kumbakonam"}, {"1", "Arun", "cuddalore"}}, nil
}

func GetMongoDbConnection() MongoConnState {
	uri := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("MongoDb Connection error", err)
	}

	return MongoConnState{client}
}
