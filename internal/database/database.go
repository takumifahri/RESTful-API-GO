package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/plugin/opentelemetry/tracing"
)

func ConnectDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	Seed(db)
	return db
	

}