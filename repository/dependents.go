package repository

import (
	"github.com/ghenoo/folhadepagamento/models"
	"gorm.io/gorm"
)

type DependentRepository interface {
	FindByID(dependentID uint) (*models.Dependent, error)
	FindByEmployeeID(employeeID uint) []models.Dependent
}

type dependentDatabase struct {
	connection *gorm.DB
}

func NewDependentRepository() DependentRepository {

	return &dependentDatabase{
		connection: db,
	}
}

func (dbs *dependentDatabase) FindByID(dependentID uint) (*models.Dependent, error) {
	var dependent models.Dependent
	result := db.First(&dependent, dependentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dependent, nil
}

func (dbs *dependentDatabase) FindByEmployeeID(employeeID uint) []models.Dependent {
	var dependents []models.Dependent
	db.Find(&dependents, "employee_id = ?", employeeID)

	return dependents
}
