package repository

import (
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
}
