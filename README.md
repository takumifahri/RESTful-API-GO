# RESTful API - Go Catalog

Proyek ini adalah sebuah RESTful API yang dibangun menggunakan Go dengan framework Echo. API ini berfungsi untuk mengelola katalog produk pakaian.

## Struktur Proyek

Proyek ini menggunakan struktur direktori yang umum digunakan dalam pengembangan aplikasi Go untuk memisahkan berbagai lapisan aplikasi.

-   **`cmd/`**: Direktori ini berisi titik masuk utama aplikasi.
    -   `main.go`: File ini bertanggung jawab untuk menginisialisasi server web Echo, menyambungkan ke database, dan mendaftarkan rute API.

-   **`internal/`**: Direktori ini berisi semua logika inti aplikasi yang tidak untuk diekspor ke proyek lain.
    -   **`database/`**: Mengelola semua yang berhubungan dengan database.
        -   [`database.go`](internal/database/database.go): Berisi fungsi untuk terhubung ke database PostgreSQL.
        -   [`seed.go`](internal/database/seed.go): Berisi data awal (seed) untuk produk pakaian yang akan dimasukkan ke dalam database saat aplikasi pertama kali dijalankan.
    -   **`delivery/rest/`**: Menangani lapisan presentasi melalui REST API.
        -   [`handler.go`](internal/delivery/rest/handler.go): Mendefinisikan struct handler dan dependensinya.
        -   [`catalog_handler.go`](internal/delivery/rest/catalog_handler.go): Berisi implementasi handler untuk endpoint katalog.
        -   [`router.go`](internal/delivery/rest/router.go): Mengatur rute-rute API dan menghubungkannya dengan handler yang sesuai.
    -   **`models/`**: Mendefinisikan struktur data (model) dan konstanta.
        -   [`catalog.go`](internal/models/catalog.go): Mendefinisikan struct `ProductClothes` yang merepresentasikan data produk.
        -   `constant/`: Berisi nilai-nilai konstan yang digunakan di seluruh aplikasi.
            -   [`catalog.go`](internal/models/constant/catalog.go): Mendefinisikan konstanta untuk tipe-tipe pakaian.
    -   **`repository/`**: Direktori ini berisi logika akses data (Data Access Layer).
        -   `catalog/`: Mengelola query database untuk katalog.
            -   [`repository.go`](internal/repository/catalog/repository.go): Mendefinisikan interface untuk repository katalog.
            -   [`catalog.go`](internal/repository/catalog/catalog.go): Implementasi dari interface repository katalog.
    -   **`usecase/`**: Direktori ini berisi logika bisnis aplikasi.
        -   `store/`: Mengelola logika bisnis untuk fitur toko/katalog.
            -   [`usecase.go`](internal/usecase/store/usecase.go): Mendefinisikan interface untuk usecase toko.
            -   [`store.go`](internal/usecase/store/store.go): Implementasi dari interface usecase toko.

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