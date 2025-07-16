package database

import (
	"fmt"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	ClothesSeed := []models.ProductClothes{
		{
			NamaPakaian: "T-Shirt Cotton",
			Price:       200000,
			Deskripsi:   "T-Shirt nyaman untuk sehari-hamodels",
			TypeClothes: constant.SHIRT,
		},
		{
			NamaPakaian: "Jeans Slim Fit",
			Price:       450000,
			Deskripsi:   "Celana jeans dengan potongan slim fit",
			TypeClothes: constant.PANTS,
		},
		{
			NamaPakaian: "Jacket Bomber",
			Price:       750000,
			Deskripsi:   "Jaket bomber stylish untuk cuaca dingin",
			TypeClothes: constant.OUTERWEAR,
		},
		{
			NamaPakaian: "Leather Watch",
			Price:       300000,
			Deskripsi:   "Jam tangan kulit premium",
			TypeClothes: constant.ACCESSORIES,
		},
		{
			NamaPakaian: "Sneakers Canvas",
			Price:       500000,
			Deskripsi:   "Sepatu sneakers canvas untuk gaya kasual",
			TypeClothes: constant.SHOES,
		},
	}

	fmt.Println("Seeding database with initial data...", ClothesSeed)

	// db, err := gorm.Open(postgres.Open(dbAddress))
	// if err != nil {
	// 	panic(err)
	// }
	// Drop tables untuk recreate dengan tipe yang benar
	db.Migrator().DropTable( &models.ProductClothes{})

	// Auto migrate dengan struct yang sudah diupdate
	db.AutoMigrate( &models.ProductClothes{})

	// Insert seed data
	if err := db.First(&models.ProductClothes{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&ClothesSeed)
		fmt.Println("Database seeded successfully!")
	}

}