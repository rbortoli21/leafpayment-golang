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
	http.HandleFunc("/employees", handler.GetAllEmployees)

	handler := corsMiddleware(http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
