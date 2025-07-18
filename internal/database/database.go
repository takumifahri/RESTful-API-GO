package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func ConnectDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	Seed(db)
	return db
	

}