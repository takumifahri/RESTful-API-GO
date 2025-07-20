# RESTful API - Go Catalog

Proyek ini adalah sebuah RESTful API yang dibangun menggunakan Go dengan framework Echo. API ini berfungsi untuk mengelola katalog produk pakaian.

## Struktur Proyek

Proyek ini menggunakan struktur direktori yang umum digunakan dalam pengembangan aplikasi Go untuk memisahkan berbagai lapisan aplikasi.
-   **`cmd/`**: Direktori ini berisi titik masuk utama aplikasi.
    -   `main.go`: File ini bertanggung jawab untuk menginisialisasi server web Echo, menyambungkan ke database, dan mendaftarkan rute API. Juga menangani command-line flags untuk migration dan seeding.

-   **`internal/`**: Direktori ini berisi semua logika inti aplikasi yang tidak untuk diekspor ke proyek lain.
    -   **`database/`**: Mengelola semua yang berhubungan dengan database.
        -   [`database.go`](internal/database/database.go): Berisi fungsi untuk terhubung ke database PostgreSQL menggunakan GORM.
        -   [`migration.go`](internal/database/migration.go): Berisi fungsi untuk migrasi database (AutoMigrate), drop tables, dan seeding data awal. Mengelola tabel `product_clothes`, `orders`, dan `product_orders`.
    -   **`delivery/rest/`**: Menangani lapisan presentasi melalui REST API.
        -   [`handler.go`](internal/delivery/rest/handler.go): Mendefinisikan struct handler dan dependensinya untuk semua endpoint.
        -   [`catalog_handler.go`](internal/delivery/rest/catalog_handler.go): Berisi implementasi handler untuk endpoint katalog (GET, POST, PATCH dengan partial update).
        -   [`order_handler.go`](internal/delivery/rest/order_handler.go): Berisi implementasi handler untuk endpoint pemesanan produk.
    -   **`delivery/routes/`**: Mengatur routing aplikasi.
        -   [`router.go`](internal/delivery/routes/router.go): Mengatur rute-rute API dan menghubungkannya dengan handler yang sesuai untuk catalog dan order endpoints.
    -   **`models/`**: Mendefinisikan struktur data (model) dan konstanta.
        -   [`catalog.go`](internal/models/catalog.go): Mendefinisikan struct `ProductClothes` dengan GORM tags dan validasi untuk produk pakaian.
        -   [`order.go`](internal/models/order.go): Mendefinisikan struct `Order`, `ProductOrder`, dan request models untuk sistem pemesanan.
        -   [`request.go`](internal/models/request.go): Mendefinisikan struct khusus untuk update requests dengan pointer fields untuk partial updates.
        -   `constant/`: Berisi nilai-nilai konstan yang digunakan di seluruh aplikasi.
            -   [`catalog.go`](internal/models/constant/catalog.go): Mendefinisikan konstanta untuk tipe-tipe pakaian (shirt, pants, outerwear, accessories, shoes).
            -   [`order.go`](internal/models/constant/order.go): Mendefinisikan konstanta untuk status order dan product order (pending, processing, completed, cancelled, failed).
    -   **`repository/`**: Direktori ini berisi logika akses data (Data Access Layer).
        -   `catalog/`: Mengelola query database untuk katalog.
            -   [`repository.go`](internal/repository/catalog/repository.go): Mendefinisikan interface untuk repository katalog dengan CRUD operations.
            -   [`catalog.go`](internal/repository/catalog/catalog.go): Implementasi repository menggunakan GORM untuk operasi database catalog.
        -   `order/`: Mengelola query database untuk pemesanan.
            -   [`repository.go`](internal/repository/order/repository.go): Mendefinisikan interface untuk repository order.
            -   [`order.go`](internal/repository/order/order.go): Implementasi repository menggunakan GORM untuk operasi database order dengan relasi ke ProductOrder.
    -   **`usecase/`**: Direktori ini berisi logika bisnis aplikasi.
        -   `store/`: Mengelola logika bisnis untuk fitur toko/katalog dan pemesanan.
            -   [`usecase.go`](internal/usecase/store/usecase.go): Mendefinisikan interface untuk usecase toko dengan method untuk catalog dan order management.
            -   [`store.go`](internal/usecase/store/store.go): Implementasi logika bisnis untuk catalog CRUD dan order processing dengan UUID generation dan validasi.
    -   **`utils/`**: Direktori untuk utility functions.
        -   [`validator.go`](internal/utils/validator.go): Berisi fungsi validasi untuk struct validation dengan support untuk full validation (CREATE) dan partial validation (UPDATE).

### Order Endpoints
-   **`POST /order`**: Membuat pesanan baru.
    -   Body: JSON dengan format:
    ```json
    {
        "order_product": [
            {
                "product_id": "PRD-unique-id",
                "quantity": 2
            }
        ]
    }
    ```

### Database Management
-   **Migration**: Menggunakan GORM AutoMigrate untuk membuat/update schema
-   **Seeding**: Data awal untuk testing dan development
-   **Fresh Migration**: Drop semua tabel dan buat ulang dengan data baru

### Validation
-   **Struct Validation**: Menggunakan `github.com/go-playground/validator/v10`
-   **Partial Update Validation**: Validasi khusus untuk update operations
-   **Business Logic Validation**: Validasi di usecase layer

### UUID Generation
-   **Unique Product ID**: Setiap produk memiliki unique ID dengan format `PRD-{uuid}`
-   **Unique Order ID**: Setiap order memiliki unique ID dengan format `ORD-{uuid}`

### Error Handling
-   **Structured Error Response**: Response error yang konsisten
-   **Logging**: Error logging untuk debugging
-   **HTTP Status Codes**: Status code yang sesuai untuk setiap scenario


## Cara Menjalankan

1.  Pastikan Anda memiliki Go dan PostgreSQL terinstal.
2.  Ubah string koneksi database di `cmd/main.go` jika diperlukan.
    ```go
    // filepath: cmd/main.go
    // ...existing code...
    const (
        dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
    )
    // ...existing code...
    ```
3.  Jalankan aplikasi dari direktori root.
    ```sh
    go run cmd/main.go
    ```
4.  Server akan berjalan di port `:8081`.

## Endpoint API

-   **`GET /clothes`**: Mengambil semua data produk pakaian.
    -   **Query Param `TypeClothes`** (opsional): Filter produk berdasarkan tipe (misal: `shirt`, `pants`).
    -   Contoh: `http://localhost:8081/clothes?TypeClothes=shirt`

## Tips dan Trik Go

Berikut adalah beberapa konsep dan konvensi penting dalam Go yang digunakan dalam proyek ini:

### 1. Visibilitas: Public (Exported) vs. Private (Unexported)

Go tidak menggunakan kata kunci seperti `public` atau `private`. Visibilitas (apakah sesuatu dapat diakses dari package lain) ditentukan oleh **kapitalisasi nama**:

-   **Huruf Kapital di Awal (PascalCase)**: Berarti **Public** atau **Exported**. Ini membuatnya bisa diakses dari package lain.
    -   Contoh: `rest.LoadRoutes(...)`, `models.ProductClothes`, `repository.GetRepository(...)`.

-   **Huruf Kecil di Awal (camelCase)**: Berarti **Private** atau **Unexported**. Ini membuatnya hanya bisa diakses di dalam package yang sama.
    -   Contoh: `type menuRepo struct` di dalam package `repository`. Struct ini tidak bisa diakses langsung dari package `main`.

Ini berlaku untuk nama `struct`, `interface`, `function`, `method`, dan `variable` di level package.

### 2. Error Handling

Pola umum di Go adalah mengembalikan `error` sebagai nilai terakhir dari sebuah fungsi. Pemanggilan fungsi tersebut harus selalu diikuti dengan pengecekan `if err != nil`.

```go
data, err := someFunction()
if err != nil {
    // Handle error di sini
    return err
}
// Lanjutkan proses jika tidak ada error
```

### 3. Formatting Otomatis dengan `go fmt`

Go memiliki tool bawaan untuk memformat kode secara konsisten. Selalu jalankan perintah ini dari direktori root proyek Anda sebelum melakukan commit untuk menjaga kerapian kode:

```sh
go fmt ./...
```


### 4. Perintah-Perintah Berguna

Berikut adalah cara menggunakan berbagai perintah untuk mengelola aplikasi:

**Menjalankan server (default):**
```sh
go run cmd/main.go
```

**Menjalankan migrasi:**
```sh
go run cmd/main.go -migrate
```

**Reset database (seperti migrate:fresh):**
```sh
go run cmd/main.go --reset
```

**Reset database dan isi data (seperti migrate:fresh --seed):**
```sh
go run cmd/main.go --fresh --seed
```

**Hanya mengisi data (jika tabel kosong):**
```sh
go run cmd/main.go --seed
```
