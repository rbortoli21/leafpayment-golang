package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rbortoli21/leafpayment-golang/models"
	"github.com/rbortoli21/leafpayment-golang/repository"
)

var employeeRepository repository.EmployeeRepository
var dependentRepository repository.DependentRepository
var workloadRepository repository.WorkloadRepository

func init() {
	employeeRepository = repository.NewEmployeeRepository()
	dependentRepository = repository.NewDependentRepository()
	workloadRepository = repository.NewWorkloadRepository()
}

func getEmployeeById(id uint) models.Employee {
	employee, error := employeeRepository.FindByID(id)
	if error != nil {
		panic("Employee not found")
	}

	dependents := dependentRepository.FindByEmployeeID(id)
	workloads := workloadRepository.FindByEmployeeID(id)

	employee.Dependents = &dependents
	employee.Workloads = &workloads

	return *employee
}

func GetLastWorkLoadByEmployeeID(employeeID uint) models.EmployeeWorkload {
	workloads := workloadRepository.FindByEmployeeID(employeeID)

	if len(workloads) == 0 {
		return models.EmployeeWorkload{}
	}

	return workloads[len(workloads)-1]
}

func GetCountFemaleEmployeesWithMoreThan16YearsByEmployerID(employerID uint) int {
	employees := employeeRepository.FindCountFemaleEmployeesByEmployerID(employerID)
	var count int
	for _, employee := range employees {
		if time.Now().Year()-employee.Birthday.Year() >= 16 {
			count++
		}
	}
	return count
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	employees := employeeRepository.FindAll()

	responseBody, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
