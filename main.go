package main

import (
	"encoding/json"
	"log"
	"math"
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

	for i := range payroll.EmployeeData {
		payroll.EmployeeData[i].GrossSalary = formatDecimal(payroll.EmployeeData[i].GrossSalary)
		payroll.EmployeeData[i].NetSalary = formatDecimal(payroll.EmployeeData[i].NetSalary)
		payroll.EmployeeData[i].Fgts = formatDecimal(payroll.EmployeeData[i].Fgts)
		payroll.EmployeeData[i].Inss = formatDecimal(payroll.EmployeeData[i].Inss)
		payroll.EmployeeData[i].Irrf = formatDecimal(payroll.EmployeeData[i].Irrf)
		payroll.EmployeeData[i].Vt = formatDecimal(payroll.EmployeeData[i].Vt)
		payroll.EmployeeData[i].Va = formatDecimal(payroll.EmployeeData[i].Va)
		payroll.EmployeeData[i].UnionContribution = formatDecimal(payroll.EmployeeData[i].UnionContribution)
		payroll.EmployeeData[i].TotalDiscounts = formatDecimal(payroll.EmployeeData[i].TotalDiscounts)
		payroll.EmployeeData[i].ExtraHours = formatDecimal(payroll.EmployeeData[i].ExtraHours)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payroll)
}

func formatDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
