package app

import (
	"api/employee/employeeDetails/domain"
	"api/employee/employeeDetails/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Inside the router")

	router := mux.NewRouter()

	/**constructor approach*/
	repository := domain.GetMongoDbConnection()
	service := service.Customer{Domain: repository}
	controller := CustomerHandler{service}

	router.HandleFunc("/Postcustomer", controller.PostCustomerHandler).Methods("POST")
	// http.HandleFunc("/welcome", Handler)

	log.Fatal(http.ListenAndServe("localhost:1000", router))
}
