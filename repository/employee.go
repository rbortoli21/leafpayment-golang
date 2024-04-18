package repository

import (
	"github.com/rbortoli21/leafpayment-golang/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindByID(employeeID uint) (*models.Employee, error)
}

type employeeDatabase struct {
	connection *gorm.DB
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeDatabase{
		connection: db,
	}
}

func (dbs *employeeDatabase) FindByID(employeeID uint) (*models.Employee, error) {
	var employee models.Employee
	var db = GetDbConnection()

	result := db.First(&employee, employeeID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}
