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
    &models.User{},
}

// Migrate hanya menjalankan AutoMigrate (create/update tables)
func Migrate(db *gorm.DB) {
    fmt.Println("Running database migrations...")
    
    for _, model := range allModels {
        tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model)[1:])
        fmt.Printf("-> Migrating table: %s\n", tableName)
    }
    
    err := db.AutoMigrate(allModels...)
    if err != nil {
        fmt.Println("Failed to migrate table:", err)
        return
    }
    fmt.Println("Migrations completed successfully.")
}

// DropTables benar-benar DROP semua tabel (seperti migrate:fresh)
func DropTables(db *gorm.DB) {
    fmt.Println("Dropping all tables...")
    
    // Drop dalam urutan terbalik untuk menghindari foreign key constraint
    modelsReversed := make([]interface{}, len(allModels))
    for i, model := range allModels {
        modelsReversed[len(allModels)-1-i] = model
    }
    
    for _, model := range modelsReversed {
        tableName := db.NamingStrategy.TableName(fmt.Sprintf("%T", model)[1:])
        fmt.Printf("-> Dropping table: %s\n", tableName)
        
        // Benar-benar DROP table, bukan AutoMigrate
        err := db.Migrator().DropTable(model)
        if err != nil {
            fmt.Printf("Warning: Failed to drop table %s: %v\n", tableName, err)
        }
    }
    
    fmt.Println("All tables dropped successfully.")
}

// DropAndMigrate = DropTables + Migrate (untuk fresh migration)
func DropAndMigrate(db *gorm.DB) {
    DropTables(db)
    Migrate(db)
}

// Seed mengisi database dengan data awal
func Seed(db *gorm.DB) {
    fmt.Println("Seeding database...")
    
    // Cek dulu apakah data sudah ada
    var count int64
    db.Model(&models.ProductClothes{}).Count(&count)
    if count > 0 {
        fmt.Println("-> Table 'product_clothes' already has data. Skipping seed.")
        return
    }

    fmt.Println("-> Seeding table: product_clothes")
    clothesSeed := []models.ProductClothes{
        {
            UNIQUEID:    fmt.Sprintf("PRD-%s", uuid.New().String()),
            NamaPakaian: "T-Shirt Cotton",
            Price:       200000,
            Deskripsi:   "T-Shirt nyaman untuk sehari-hari",
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

    if err := db.Create(&clothesSeed).Error; err != nil {
        fmt.Println("Failed to seed database:", err)
        return
    }
    fmt.Println("Database seeded successfully.")
}

// FreshSeed = DropTables + Migrate + Seed (complete fresh start)
func FreshSeed(db *gorm.DB) {
    DropTables(db)
    Migrate(db)
    Seed(db)
}