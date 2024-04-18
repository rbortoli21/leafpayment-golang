package repository

import (
	"github.com/rbortoli21/leafpayment-golang/models"
	"gorm.io/gorm"
)

type EmployerRepository interface {
	FindByID(employerID uint) *models.Employer
}

type employerDatabase struct {
	connection *gorm.DB
}

func NewEmployerRepository() EmployerRepository {
	var DB = GetDbConnection()
	return &employerDatabase{
		connection: DB,
	}
}

func (dbs *employerDatabase) FindByID(employerID uint) *models.Employer {
	var employer models.Employer
	result := db.First(&employer, employerID)
	if result.Error != nil {
		return nil
	}
	return &employer
}
