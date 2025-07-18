package main

import (
    "flag" // Import paket flag
    "fmt"  // Import paket fmt
    "os"   // Import paket os

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

    // 2. Parse flags yang diberikan saat program dijalankan
    flag.Parse()

    // 3. Koneksikan ke database (diperlukan untuk semua perintah)
    db := database.ConnectDB(dbAddress)
    if db == nil {
        fmt.Println("FATAL: Failed to connect to database")
        os.Exit(1) // Keluar dari program jika koneksi gagal
    }

    // 4. Logika untuk menjalankan perintah berdasarkan flag
    // Perintah: go run cmd/main.go -fresh -seed
    if *freshFlag {
        database.DropTables(db)
        database.Migrate(db)
        if *seedFlag {
            database.Seed(db)
        }
        fmt.Println("Fresh migration completed.")
        return // Hentikan program setelah selesai
    }

    // Perintah: go run cmd/main.go -migrate
    if *migrateFlag {
        database.Migrate(db)
        if *seedFlag {
            database.Seed(db)
        }
        fmt.Println("Migration completed.")
        return // Hentikan program setelah selesai
    }

    // Perintah: go run cmd/main.go -seed
    if *seedFlag {
        database.Seed(db)
        return // Hentikan program setelah selesai
    }

    // 5. Jika tidak ada flag, jalankan server (perilaku default)
    fmt.Println("No command flags detected, starting server...")
    e := echo.New()
    catalogRepo := strRepo.GetRepository(db)
    orderRepository := orderRepo.GetRepository(db)
    storeUsecase := strUsecase.GetUsecase(catalogRepo, orderRepository)
    handler := rest.NewHandler(storeUsecase) // Menggunakan NewHandler sesuai kode Anda

    routers.LoadRoutes(e, handler)

    // Logger untuk port nya
    e.Logger.Fatal(e.Start(":8081"))
}