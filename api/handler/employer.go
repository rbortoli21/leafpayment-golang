package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rbortoli21/leafpayment-golang/models"
	"github.com/rbortoli21/leafpayment-golang/repository"
)

var employerRepository repository.EmployerRepository
var employerConfiguratorRepository repository.EmployerConfiguratorRepository

func init() {
	employerRepository = repository.NewEmployerRepository()
	employerConfiguratorRepository = repository.NewEmployerConfiguratorRepository()
}

func getEmployerById(id uint) models.Employer {
	employer := employerRepository.FindByID(id)
	configurator := employerConfiguratorRepository.FindByEmployerID(id)

	employer.Configurator = configurator

	return *employer
}

func GetEmployerById(w http.ResponseWriter, r *http.Request) {
	employeeIDStr := r.URL.Path[len("/find/"):]

	employerId, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	payroll := getEmployerById((uint(employerId)))

	responseBody, err := json.Marshal(payroll)
	if err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
