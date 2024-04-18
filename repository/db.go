package repository

import (
	"github.com/rbortoli21/leafpayment-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDbConnection() *gorm.DB {
	return db
}

func ConnectWithDataBase() {
	dsn := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.prneelhggfwsjvolwjdu password=@Folhastech01 dbname=leafpayment port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect with database")
	}

	db.Debug().Logger = db.Debug().Logger.LogMode(3)

	db.AutoMigrate(&models.EmployerConfigurator{})
	db.AutoMigrate(&models.EmployeeWorkload{})
	db.AutoMigrate(&models.Employer{})
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Dependent{})
	db.AutoMigrate(&models.EmployeeWorkload{})

}
