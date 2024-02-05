package app

import (
	"api/employee/employeeDetails/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type CustomerHandler struct {
	Service service.ICustomer
}

func (ch CustomerHandler) PostCustomerHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside handler method")

	response := ch.Service.PostCustomer(req.Body)
	fmt.Println("struct data", response)

	if req.URL.Query().Get("contentType") == "application/json" {
		fmt.Println("contentType is Json")
		res.Header().Add("content-Type", "application/json")
		json.NewEncoder(res).Encode(response)
	} else {
		fmt.Println("contentType is Xml")
		res.Header().Add("content-Type", "application/xml")
		xml.NewEncoder(res).Encode(response)
	}

}
