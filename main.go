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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payroll)
}
