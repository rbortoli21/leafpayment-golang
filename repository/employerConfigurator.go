package repository

import (
	"github.com/ghenoo/folhadepagamento/models"
	"gorm.io/gorm"
)

type EmployerConfiguratorRepository interface {
	FindByEmployerID(employerID uint) *models.EmployerConfigurator
}

type employerConfiguratorDatabase struct {
	connection *gorm.DB
}

func NewEmployerConfiguratorRepository() EmployerConfiguratorRepository {
	return &employerConfiguratorDatabase{
		connection: db,
	}
}

func (dbs *employerConfiguratorDatabase) FindByEmployerID(employerID uint) *models.EmployerConfigurator {
	var employerConfigurator models.EmployerConfigurator

	result := db.First(&employerConfigurator, employerID)
	if result.Error != nil {
		return nil
	}
	return &employerConfigurator
}
