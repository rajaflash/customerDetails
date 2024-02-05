package app

import (
	"api/employee/employeeDetails/domain"
	"api/employee/employeeDetails/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Inside the router")
	// Load variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	router := mux.NewRouter()

	/**constructor approach*/
	repository := domain.Person{}
	service := service.Customer{Domain: repository}
	controller := CustomerHandler{service}

	router.HandleFunc("/customer", controller.Handler)
	// http.HandleFunc("/welcome", Handler)

	log.Fatal(http.ListenAndServe("localhost:1000", router))
}
