// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/generate-payroll", generatePayrollHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generatePayrollHandler(w http.ResponseWriter, r *http.Request) {
	var employer Employer
	err := json.NewDecoder(r.Body).Decode(&employer)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	payroll := GeneratePayroll(employer)

	responseBody, err := json.Marshal(payroll)
	if err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
