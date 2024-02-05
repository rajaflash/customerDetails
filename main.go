package main

import (
	"api/employee/employeeDetails/app"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	// Load variables from .env file
	err := godotenv.Load("F:\\coding\\GolangApp\\employeeDetails\\properties.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	app.Start()
}
