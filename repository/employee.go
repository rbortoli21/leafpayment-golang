package repository

import (
	"github.com/rbortoli21/leafpayment-golang/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindByID(employeeID uint) (*models.Employee, error)
	FindByEmployerID(employerID uint) []models.Employee
	FindCountFemaleEmployeesByEmployerID(employerID uint) []models.Employee
	FindAll() []models.Employee
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

func (dbs *employeeDatabase) FindByEmployerID(employerID uint) []models.Employee {
	var employees []models.Employee
	var db = GetDbConnection()

	result := db.Where("employer_id = ?", employerID).Find(&employees)
	if result.Error != nil {
		return nil
	}
	return employees
}

func (dbs *employeeDatabase) FindCountFemaleEmployeesByEmployerID(employerID uint) []models.Employee {
	var employees []models.Employee
	var db = GetDbConnection()

	db.Where("employer_id = ? AND gender like 'Fem'", employerID).Find(&employees)

	return employees
}

func (dbs *employeeDatabase) FindAll() []models.Employee {
	var employees []models.Employee
	var db = GetDbConnection()

	result := db.Find(&employees)
	if result.Error != nil {
		return nil
	}
	return employees
}
