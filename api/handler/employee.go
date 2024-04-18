package handler

import (
	"github.com/ghenoo/folhadepagamento/models"
	"github.com/ghenoo/folhadepagamento/repository"
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
	/* dependents := dependentRepository.FindByEmployeeID(id)
	workloads := workloadRepository.FindByEmployeeID(id)

	employee.Dependents = &dependents
	employee.Workloads = &workloads
	*/
	return *employee
}
