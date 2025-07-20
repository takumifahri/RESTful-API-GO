package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/takumifahri/RESTful-API-GO/internal/database"
    "github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
    routers "github.com/takumifahri/RESTful-API-GO/internal/delivery/routes"

    strRepo "github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
    orderRepo "github.com/takumifahri/RESTful-API-GO/internal/repository/order"
    strUsecase "github.com/takumifahri/RESTful-API-GO/internal/usecase/store"
)

const (
    dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {
    // 1. Definisikan flags untuk command-line
    migrateFlag := flag.Bool("migrate", false, "Run database migrations")
    freshFlag := flag.Bool("fresh", false, "Drop all tables and run migrations (fresh start)")
    seedFlag := flag.Bool("seed", false, "Seed the database with initial data")
    resetFlag := flag.Bool("reset", false, "Drop all tables (same as fresh without seed)")

    // 2. Parse flags yang diberikan saat program dijalankan
    flag.Parse()

    // 3. Koneksikan ke database (diperlukan untuk semua perintah)
    db := database.ConnectDB(dbAddress)
    if db == nil {
        fmt.Println("FATAL: Failed to connect to database")
        os.Exit(1)
    }

    // 4. Logika untuk menjalankan perintah berdasarkan flag
    
    // Perintah: go run cmd/main.go --fresh --seed (drop + migrate + seed)
    if *freshFlag && *seedFlag {
        fmt.Println("Running fresh migration with seed...")
        database.FreshSeed(db)
        fmt.Println("Fresh migration with seed completed.")
        return
    }

    // Perintah: go run cmd/main.go --fresh (drop + migrate)
    if *freshFlag {
        fmt.Println("Running fresh migration...")
        database.DropAndMigrate(db)
        fmt.Println("Fresh migration completed.")
        return
    }

    // Perintah: go run cmd/main.go --reset (hanya drop tables)
    if *resetFlag {
        fmt.Println("Resetting database (dropping all tables)...")
        database.DropTables(db)
        fmt.Println("Database reset completed.")
        return
    }

    // Perintah: go run cmd/main.go --migrate --seed
    if *migrateFlag && *seedFlag {
        fmt.Println("Running migration with seed...")
        database.Migrate(db)
        database.Seed(db)
        fmt.Println("Migration with seed completed.")
        return
    }

    // Perintah: go run cmd/main.go --migrate
    if *migrateFlag {
        fmt.Println("Running migration...")
        database.Migrate(db)
        fmt.Println("Migration completed.")
        return
    }

    // Perintah: go run cmd/main.go --seed
    if *seedFlag {
        fmt.Println("Running seed...")
        database.Seed(db)
        fmt.Println("Seed completed.")
        return
    }

    // 5. Jika tidak ada flag, jalankan server (perilaku default)
    fmt.Println("No command flags detected, starting server...")
    e := echo.New()
    catalogRepo := strRepo.GetRepository(db)
    orderRepository := orderRepo.GetRepository(db)
    storeUsecase := strUsecase.GetUsecase(catalogRepo, orderRepository)
    handler := rest.NewHandler(storeUsecase)

    routers.LoadRoutes(e, handler)

    fmt.Println("Server starting on port :8081")
    e.Logger.Fatal(e.Start(":8081"))
}