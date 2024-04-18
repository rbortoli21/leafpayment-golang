package repository

import (
	"github.com/rbortoli21/leafpayment-golang/models"
	"gorm.io/gorm"
)

type WorkloadRepository interface {
	FindByEmployeeID(employeeID uint) []models.EmployeeWorkload
}

type workloadDatabase struct {
	connection *gorm.DB
}

func NewWorkloadRepository() WorkloadRepository {
	return &workloadDatabase{
		connection: db,
	}
}

func (dbs *workloadDatabase) FindByEmployeeID(employeeID uint) []models.EmployeeWorkload {
	var workloads []models.EmployeeWorkload
	db.Find(&workloads, "employee_id = ?", employeeID)

	return workloads
}
