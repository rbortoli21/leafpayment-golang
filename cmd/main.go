package main

import (
	"log"
	"net/http"

	"github.com/rbortoli21/leafpayment-golang/api/handler"
	"github.com/rbortoli21/leafpayment-golang/repository"
)

func main() {
	repository.ConnectWithDataBase()

	http.HandleFunc("/generate-payroll/{id}", handler.GeneratePayrollHandler)
	http.HandleFunc("/find/{id}", handler.GetEmployerById)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
