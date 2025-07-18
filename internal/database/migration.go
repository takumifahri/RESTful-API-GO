package database

import (
    "fmt"

    "github.com/google/uuid"
    "github.com/takumifahri/RESTful-API-GO/internal/models"
    "github.com/takumifahri/RESTful-API-GO/internal/models/constant"
    "gorm.io/gorm"
)
var allModels = []interface{}{
    &models.ProductClothes{},
    &models.Order{},
    &models.ProductOrder{},
}
// Migrate hanya menjalankan AutoMigrate
func Migrate(db *gorm.DB) {
    fmt.Println("Running database migrations...")
    // Menambahkan nama tabel yang dimigrasi
    for _, model := range allModels {
        tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model)[1:])
        fmt.Printf("-> Dropping table: %s\n", tableName)
    }
    err := db.AutoMigrate(allModels...)
    if err != nil {
        fmt.Println("Failed to migrate table:", err)
        return
    }
    fmt.Println("Migrations completed successfully.")
}

// DropTables menghapus tabel
func DropTables(db *gorm.DB) {
    fmt.Println("Dropping tables...")
    // Menambahkan nama tabel yang dihapus
    for _, model := range allModels {
        tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model)[1:])
        fmt.Printf("-> Dropping table: %s\n", tableName)
    }
    err := db.AutoMigrate(allModels...)
    if err != nil {
        fmt.Println("Failed to drop table:", err)
        return
    }
    fmt.Println("Tables dropped successfully.")
}

// Seed mengisi database dengan data awal
func Seed(db *gorm.DB) {
    // Cek dulu apakah data sudah ada
    var count int64
    db.Model(&models.ProductClothes{}).Count(&count)
    if count > 0 {
        // Menambahkan nama tabel di pesan skip
        fmt.Println("Table 'product_clothes' already seeded. Skipping.")
        return
    }

    // Menambahkan nama tabel yang di-seed
    fmt.Println("-> Seeding table: product_clothes")
    ClothesSeed := []models.ProductClothes{
        {

            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "T-Shirt Cotton",
            Price:       200000,
            Deskripsi:   "T-Shirt nyaman untuk sehari-hamodels",
            Stock:       10,
            TypeClothes: constant.SHIRT,
        },
        {
            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "Jeans Slim Fit",
            Price:       450000,
            Deskripsi:   "Celana jeans dengan potongan slim fit",
            Stock:       5,
            TypeClothes: constant.PANTS,
        },
        {
            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "Jacket Bomber",
            Price:       750000,
            Deskripsi:   "Jaket bomber stylish untuk cuaca dingin",
            Stock:       3,
            TypeClothes: constant.OUTERWEAR,
        },
        {
            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "Leather Watch",
            Price:       300000,
            Deskripsi:   "Jam tangan kulit premium",
            Stock:       8,
            TypeClothes: constant.ACCESSORIES,
        },
        {
            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "Sneakers Canvas",
            Price:       500000,
            Deskripsi:   "Sepatu sneakers canvas untuk gaya kasual",
            Stock:       12,
            TypeClothes: constant.SHOES,
        },
    }

    if err := db.Create(&ClothesSeed).Error; err != nil {
        fmt.Println("Failed to seed database:", err)
        return
    }
    fmt.Println("Database seeded successfully.")
}