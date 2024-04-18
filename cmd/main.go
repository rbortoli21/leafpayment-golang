package main

import (
	"log"
	"net/http"

	"github.com/ghenoo/folhadepagamento/api/handler"
	"github.com/ghenoo/folhadepagamento/repository"
)

func main() {
	repository.ConnectWithDataBase()

	http.HandleFunc("/generate-payroll/{id}", handler.GeneratePayrollHandler)
	http.HandleFunc("/find/{id}", handler.GetEmployerById)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
