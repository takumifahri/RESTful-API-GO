# RESTful API - Go Catalog

Proyek ini adalah sebuah RESTful API yang dibangun menggunakan Go dengan framework Echo. API ini berfungsi untuk mengelola katalog produk pakaian.

## Struktur Proyek

Proyek ini menggunakan struktur direktori yang umum digunakan dalam pengembangan aplikasi Go untuk memisahkan berbagai lapisan aplikasi.
-   **`cmd/`**: Direktori ini berisi titik masuk utama aplikasi.
    -   `main.go`: File ini bertanggung jawab untuk menginisialisasi server web Echo, menyambungkan ke database, dan mendaftarkan rute API. Juga menangani command-line flags untuk migration dan seeding.

-   -   **`internal/`**: Direktori ini berisi semua logika inti aplikasi yang tidak untuk diekspor ke proyek lain.
    -   **`database/`**: Mengelola semua yang berhubungan dengan database.
        -   [`database.go`](internal/database/database.go): Berisi fungsi untuk terhubung ke database PostgreSQL menggunakan GORM.
        -   [`migration.go`](internal/database/migration.go): Berisi fungsi untuk migrasi database (AutoMigrate), drop tables, dan seeding data awal. Mengelola tabel `product_clothes`, `orders`, `product_orders`, dan `users`. Menyediakan fungsi `DropTables()`, `Migrate()`, `Seed()`, `FreshSeed()`, dan `DropAndMigrate()`.

    -   **`delivery/rest/`**: Menangani lapisan presentasi melalui REST API.
        -   [`handler.go`](internal/delivery/rest/handler.go): Mendefinisikan struct handler utama dengan dependensi `storeUsecase` dan `AuthUsecase` untuk semua endpoint.
        -   [`catalog_handler.go`](internal/delivery/rest/catalog_handler.go): Berisi implementasi handler untuk endpoint katalog (GET, POST, PATCH dengan partial update). Termasuk validasi menggunakan `utils.ValidateStruct()`.
        -   [`order_handler.go`](internal/delivery/rest/order_handler.go): Berisi implementasi handler untuk endpoint pemesanan produk (`Order`, `GetOrderInfo`, `AdminGetAllOrder`).
        -   **`user/`**: Sub-direktori untuk handler autentikasi.
            -   [`auh_handler.go`](internal/delivery/rest/user/auh_handler.go): Berisi implementasi handler untuk autentikasi user (`RegisterUser`) dengan embedded struct dari `rest.Handler`.

    -   **`delivery/routes/`**: Mengatur routing aplikasi.
        -   [`router.go`](internal/delivery/routes/router.go): Mengatur rute-rute API dan menghubungkannya dengan handler yang sesuai. Meliputi endpoint catalog, order, dan autentikasi (`/auth/register`).

    -   **`middlewares/`**: Mengelola middleware aplikasi.
        -   [`middlewares.go`](internal/middlewares/middlewares.go): Mendefinisikan middleware umum dan `RoleChecker` untuk authorization.
        -   [`cors.go`](internal/middlewares/cors.go): Implementasi CORS middleware menggunakan Echo CORS dengan konfigurasi `AllowOrigins: ["*"]`.

    -   **`models/`**: Mendefinisikan struktur data (model) dan konstanta.
        -   [`catalog.go`](internal/models/catalog.go): Mendefinisikan struct `ProductClothes` dengan GORM tags dan validasi untuk produk pakaian.
        -   [`order.go`](internal/models/order.go): Mendefinisikan struct `Order`, `ProductOrder`, dan request models (`OrderMenuRequest`, `GetOrderInfoRequest`, `GetAllOrderRequest`) untuk sistem pemesanan dengan relasi database.
        -   [`user.go`](internal/models/user.go): Mendefinisikan struct `User` untuk autentikasi dengan field `Hash` untuk menyimpan password yang di-encrypt, dan `RegisterRequest` untuk registrasi user.
        -   **`constant/`**: Berisi nilai-nilai konstan yang digunakan di seluruh aplikasi.
            -   [`catalog.go`](internal/models/constant/catalog.go): Konstanta untuk tipe pakaian (shirt, pants, outerwear, accessories, shoes).
            -   [`order.go`](internal/models/constant/order.go): Konstanta untuk status order dan product order (pending, processing, completed, cancelled, failed).
            -   [`user.go`](internal/models/constant/user.go): Konstanta untuk roles user (admin, user, cashier, manager).

    -   **`repository/`**: Direktori ini berisi logika akses data (Data Access Layer).
        -   **`catalog/`**: Mengelola query database untuk katalog.
            -   [`repository.go`](internal/repository/catalog/repository.go): Interface untuk repository katalog dengan CRUD operations.
            -   [`catalog.go`](internal/repository/catalog/catalog.go): Implementasi repository menggunakan GORM untuk operasi database catalog.
        -   **`order/`**: Mengelola query database untuk pemesanan.
            -   [`repository.go`](internal/repository/order/repository.go): Interface untuk repository order.
            -   [`order.go`](internal/repository/order/order.go): Implementasi repository menggunakan GORM untuk operasi database order dengan relasi ke ProductOrder menggunakan `Preload()`.
        -   **`users/`**: Mengelola query database untuk user management.
            -   **`auth/`**: Repository khusus untuk autentikasi.
                -   [`repository.go`](internal/repository/users/auth/repository.go): Interface untuk repository autentikasi (`RegisterUser`, `CheckRegistered`, `GenerateUserHash`).
                -   [`auth.go`](internal/repository/users/auth/auth.go): Implementasi repository autentikasi dengan AES-GCM cipher dan konfigurasi Argon2.
                -   [`hash.go`](internal/repository/users/auth/hash.go): Implementasi password hashing menggunakan Argon2ID dengan enkripsi AES-GCM tambahan.

    -   **`usecase/`**: Direktori ini berisi logika bisnis aplikasi.
        -   **`store/`**: Mengelola logika bisnis untuk fitur toko/katalog dan pemesanan.
            -   [`usecase.go`](internal/usecase/store/usecase.go): Interface untuk usecase toko dengan method untuk catalog dan order management.
            -   [`store.go`](internal/usecase/store/store.go): Implementasi logika bisnis untuk catalog CRUD dan order processing dengan UUID generation, validasi, dan business rules.
        -   **`auth/`**: Mengelola logika bisnis untuk autentikasi.
            -   [`usecase.go`](internal/usecase/auth/usecase.go): Interface untuk usecase autentikasi.
            -   [`implements.go`](internal/usecase/auth/implements.go): Implementasi logika bisnis autentikasi dengan validasi user registration, password hashing, dan duplicate checking.
        -   **`admin/`**: Reserved untuk logika bisnis admin (dalam pengembangan).
            -   [`usecase.go`](internal/usecase/admin/usecase.go): Placeholder untuk admin usecase.

    -   **`utils/`**: Direktori untuk utility functions.
        -   [`validator.go`](internal/utils/validator.go): Berisi fungsi validasi untuk struct validation dengan support untuk full validation (CREATE) dan partial validation (UPDATE). Menggunakan `github.com/go-playground/validator/v10`.
        -   [`JWT.go`](internal/utils/JWT.go): Placeholder untuk implementasi JWT token generation dan validation (dalam pengembangan).

## 🗄️ **Struktur Database**

Aplikasi menggunakan **PostgreSQL** dengan **GORM** sebagai ORM dan mengelola tabel-tabel berikut:

### **Tabel `product_clothes`**
- Primary Key: `id` (auto-increment)
- Unique Key: `unique_id` (format: PRD-{uuid})
- Fields: `nama_pakaian`, `price`, `deskripsi`, `stock`, `type_clothes`

### **Tabel `orders`**
- Primary Key: `id` (auto-increment) 
- Unique Key: `unique_id` (format: ORD-{uuid})
- Unique Key: `reference_id` (untuk mencegah duplicate orders)
- Fields: `status`
- **Relasi**: Has Many `product_orders`

### **Tabel `product_orders`**
- Primary Key: `id` (auto-increment)
- Foreign Key: `order_unique_id` → `orders.unique_id`
- Fields: `product_id`, `nama_pakaian`, `quantity`, `total_price`, `status`

### **Tabel `users`**
- Primary Key: `id` (auto-increment)
- Unique Key: `unique_id` (format: USR-{uuid})
- Unique Fields: `name`, `email`, `phone`
- Fields: `hash` (encrypted password), `address`, `roles`

## 🔐 **Security Features**

### **Password Security**
- **Argon2ID** hashing algorithm (winner of Password Hashing Competition)
- **AES-GCM** encryption untuk additional layer security
- **Salt generation** menggunakan crypto/rand
- **No plaintext password storage** - hanya hash yang disimpan

### **Database Security**
- **UUID untuk semua identifier** (mencegah enumeration attacks)
- **Unique constraints** pada field sensitif
- **Reference ID validation** untuk mencegah duplicate orders
- **GORM query protection** terhadap SQL injection

### **API Security**
- **Struct validation** menggunakan validator/v10
- **Error handling** yang tidak membocorkan informasi sensitif
- **CORS configuration** untuk cross-origin requests
- **Type safety** dengan Go's strong typing

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

# 1. Jalankan server biasa
go run cmd/main.go

# 2. Migration saja (create/update tables)
go run cmd/main.go --migrate

# 3. Migration + seed
go run cmd/main.go --migrate --seed

# 4. Fresh migration (drop + create tables) - seperti migrate:fresh
go run cmd/main.go --fresh

# 5. Fresh migration + seed (drop + create + seed) - seperti migrate:fresh --seed
go run cmd/main.go --fresh --seed

# 6. Reset database (drop semua tables)
go run cmd/main.go --reset

# 7. Seed saja (jika tabel sudah ada)
go run cmd/main.go --seed


### Struktur JWT bisa di bagi 3
**1. Header**
    - Berisi informasi tentang algoritma enkripsi yang digunakan (misal: HS256, RS256)
    - Tipe token (biasanya "JWT")
    - Di-encode dalam Base64URL

**2. Payload**
    - Berisi data/claims tentang user atau informasi lainnya
    - Claims bisa berupa: user ID, role, expiration time, dll
    - Di-encode dalam Base64URL
    - Contoh claims: `sub` (subject), `exp` (expiration), `iat` (issued at)
    - Ada 3 jenis claims:
        - **Registered Claims**: Claims yang sudah terdefinisi dalam standar JWT
            - `iss` (issuer): Penerbit token
            - `sub` (subject): Subjek token (biasanya user ID)
            - `aud` (audience): Target penerima token
            - `exp` (expiration time): Waktu kedaluwarsa token
            - `nbf` (not before): Token tidak valid sebelum waktu ini
            - `iat` (issued at): Waktu token dibuat
            - `jti` (JWT ID): ID unik untuk token
        - **Public Claims**: Claims yang didefinisikan secara terbuka dan terdaftar
            - Biasanya menggunakan namespace URI untuk menghindari konflik
            - Contoh: `https://example.com/jwt_claims/role`
        - **Private Claims**: Claims kustom yang dibuat untuk keperluan aplikasi spesifik
            - Tidak terdaftar secara publik
            - Contoh: `user_id`, `permissions`, `department`

**3. Signature**
    - Digunakan untuk memverifikasi bahwa token tidak diubah
    - Dibuat dengan menggabungkan header + payload + secret key
    - Menggunakan algoritma yang disebutkan di header
    - Format: `HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)`

Format lengkap JWT: `header.payload.signature`


### Argon2 Password Hashing

**Argon2** adalah algoritma hashing password yang memenangkan kompetisi Password Hashing Competition (PHC) pada tahun 2015. Argon2 dirancang untuk menjadi resistance terhadap serangan GPU dan ASIC, serta memberikan keamanan yang sangat tinggi untuk penyimpanan password.

#### Keunggulan Argon2

1. **Memory-Hard Function**: Membutuhkan memori yang besar, membuat serangan brute force menjadi mahal
2. **Time-Cost Parameter**: Dapat mengatur waktu komputasi yang diperlukan
3. **Parallelism**: Dapat memanfaatkan multiple cores untuk meningkatkan keamanan
4. **Side-Channel Resistance**: Tahan terhadap serangan side-channel

#### Varian Argon2

**1. Argon2d**
- Data-dependent memory access
- Lebih cepat dan menggunakan memori secara maksimal
- Rentan terhadap side-channel attacks
- Cocok untuk aplikasi tanpa ancaman side-channel

**2. Argon2i**
- Data-independent memory access
- Tahan terhadap side-channel attacks
- Sedikit lebih lambat dari Argon2d
- Cocok untuk hashing password

**3. Argon2id** (Recommended)
- Hybrid dari Argon2d dan Argon2i
- Menggunakan Argon2i untuk pass pertama, Argon2d untuk pass selanjutnya
- Memberikan keseimbangan terbaik antara keamanan dan performa
- **Direkomendasikan untuk kebanyakan use case**

#### Parameter Argon2

```go
// Contoh konfigurasi Argon2id
type Argon2Params struct {
    Memory      uint32 // Memori dalam KB (misal: 64*1024 = 64MB)
    Iterations  uint32 // Jumlah iterasi (misal: 1-3)
    Parallelism uint8  // Tingkat paralelisme (misal: 1-4)
    SaltLength  uint32 // Panjang salt dalam bytes (misal: 16)
    KeyLength   uint32 // Panjang hash output (misal: 32)
}
```

**Parameter Guidelines:**
- **Memory**: 64MB untuk aplikasi web, 1GB+ untuk high-security
- **Iterations**: 1-3 iterasi (lebih tinggi = lebih lambat)
- **Parallelism**: Sesuai dengan jumlah CPU cores
- **Salt Length**: Minimal 16 bytes, recommended 32 bytes
- **Key Length**: 32 bytes (256-bit) untuk keamanan optimal

#### Implementasi di Go

```go
package main

import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "errors"
    "fmt"
    "strings"
    
    "golang.org/x/crypto/argon2"
)

type Argon2Params struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

// Recommended parameters untuk production
var DefaultParams = &Argon2Params{
    Memory:      64 * 1024, // 64 MB
    Iterations:  3,
    Parallelism: 2,
    SaltLength:  16,
    KeyLength:   32,
}

// Generate hash dari password
func GenerateFromPassword(password string, params *Argon2Params) (string, error) {
    salt, err := generateRandomBytes(params.SaltLength)
    if err != nil {
        return "", err
    }
    
    hash := argon2.IDKey([]byte(password), salt, params.Iterations, 
                        params.Memory, params.Parallelism, params.KeyLength)
    
    // Format: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
    encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
    encodedHash := base64.RawStdEncoding.EncodeToString(hash)
    
    return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
        params.Memory, params.Iterations, params.Parallelism,
        encodedSalt, encodedHash), nil
}

// Verify password dengan hash
func CompareHashAndPassword(password, hashedPassword string) (bool, error) {
    params, salt, hash, err := decodeHash(hashedPassword)
    if err != nil {
        return false, err
    }
    
    otherHash := argon2.IDKey([]byte(password), salt, params.Iterations,
                             params.Memory, params.Parallelism, params.KeyLength)
    
    // Gunakan subtle.ConstantTimeCompare untuk mencegah timing attacks
    return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    return b, err
}

func decodeHash(hashedPassword string) (*Argon2Params, []byte, []byte, error) {
    parts := strings.Split(hashedPassword, "$")
    if len(parts) != 6 {
        return nil, nil, nil, errors.New("invalid hash format")
    }
    
    var version int
    _, err := fmt.Sscanf(parts[2], "v=%d", &version)
    if err != nil || version != argon2.Version {
        return nil, nil, nil, errors.New("invalid version")
    }
    
    params := &Argon2Params{}
    _, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d",
        &params.Memory, &params.Iterations, &params.Parallelism)
    if err != nil {
        return nil, nil, nil, err
    }
    
    salt, err := base64.RawStdEncoding.DecodeString(parts[4])
    if err != nil {
        return nil, nil, nil, err
    }
    params.SaltLength = uint32(len(salt))
    
    hash, err := base64.RawStdEncoding.DecodeString(parts[5])
    if err != nil {
        return nil, nil, nil, err
    }
    params.KeyLength = uint32(len(hash))
    
    return params, salt, hash, nil
}
```

#### Best Practices

1. **Parameter Tuning**: Test performa di environment production untuk menentukan parameter optimal
2. **Salt Generation**: Gunakan cryptographically secure random number generator
3. **Timing Attack Protection**: Gunakan `subtle.ConstantTimeCompare` untuk verifikasi
4. **Memory Considerations**: Sesuaikan parameter memory dengan kapasitas server
5. **Upgrade Path**: Simpan versi dan parameter dalam hash untuk kemudahan upgrade

#### Perbandingan dengan Algoritma Lain

| Algoritma | Keamanan | Performa | Memory Usage | Recommendation |
|-----------|----------|----------|--------------|----------------|
| MD5       | ❌ Sangat Rendah | ⚡ Sangat Cepat | 💾 Rendah | ❌ Jangan gunakan |
| SHA-256   | ⚠️ Rendah | ⚡ Cepat | 💾 Rendah | ❌ Tidak untuk password |
| bcrypt    | ✅ Baik | 🐌 Lambat | 💾 Rendah | ✅ Acceptable |
| scrypt    | ✅ Baik | 🐌 Lambat | 💾 Tinggi | ✅ Baik |
| **Argon2id** | **✅ Sangat Baik** | **🐌 Lambat** | **💾 Tinggi** | **✅ Terbaik** |

#### Contoh Penggunaan dalam Aplikasi

```go
// Saat registrasi user
func RegisterUser(email, password string) error {
    hashedPassword, err := GenerateFromPassword(password, DefaultParams)
    if err != nil {
        return err
    }
    
    // Simpan ke database
    user := &User{
        Email:    email,
        Password: hashedPassword,
    }
    return db.Create(user).Error
}

// Saat login user
func LoginUser(email, password string) (*User, error) {
    var user User
    err := db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    
    match, err := CompareHashAndPassword(password, user.Password)
    if err != nil || !match {
        return nil, errors.New("invalid credentials")
    }
    
    return &user, nil
}
```

**Argon2** memberikan proteksi terbaik untuk password storage dan harus menjadi pilihan utama untuk aplikasi modern yang mengutamakan keamanan.



### Perbedaan argon2.IDKey vs argon2.Key

#### argon2.IDKey (Argon2id)
```go
func IDKey(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
```

**Karakteristik:**
- Menggunakan varian **Argon2id** (hybrid)
- **Recommended untuk password hashing**
- Menggabungkan keunggulan Argon2d dan Argon2i
- Pass pertama menggunakan Argon2i (data-independent)
- Pass kedua dan seterusnya menggunakan Argon2d (data-dependent)
- **Tahan terhadap side-channel attacks**
- **Optimal untuk aplikasi web dan mobile**

#### argon2.Key (Argon2d)
```go
func Key(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
```

**Karakteristik:**
- Menggunakan varian **Argon2d** (data-dependent)
- Akses memori bergantung pada data input
- **Lebih cepat** dibanding Argon2i dan Argon2id
- **Rentan terhadap side-channel attacks**
- Cocok untuk key derivation dalam environment yang aman
- **Tidak direkomendasikan untuk password hashing**

#### Perbandingan Detail

| Aspek | argon2.IDKey (Argon2id) | argon2.Key (Argon2d) |
|-------|------------------------|---------------------|
| **Keamanan** | ✅ Sangat Tinggi | ⚠️ Tinggi (tapi rentan side-channel) |
| **Performa** | 🐌 Sedang | ⚡ Lebih Cepat |
| **Side-Channel Resistance** | ✅ Ya | ❌ Tidak |
| **Use Case** | Password Hashing | Key Derivation |
| **Recommendation** | ✅ **Pilihan Utama** | ⚠️ Hanya untuk KDF |

#### Contoh Implementasi

```go
package main

import (
    "fmt"
    "golang.org/x/crypto/argon2"
)

func main() {
    password := []byte("mySecretPassword")
    salt := []byte("randomSalt123456") // 16 bytes
    
    // Parameter yang sama untuk perbandingan
    time := uint32(3)
    memory := uint32(64 * 1024) // 64 MB
    threads := uint8(2)
    keyLen := uint32(32)
    
    // Argon2id - RECOMMENDED untuk password
    hashID := argon2.IDKey(password, salt, time, memory, threads, keyLen)
    fmt.Printf("Argon2id: %x\n", hashID)
    
    // Argon2d - untuk key derivation saja
    hashD := argon2.Key(password, salt, time, memory, threads, keyLen)
    fmt.Printf("Argon2d:  %x\n", hashD)
    
    // Output akan berbeda karena algoritma yang berbeda
}
```

#### Kapan Menggunakan Masing-Masing

**Gunakan argon2.IDKey ketika:**
- ✅ Hashing password untuk autentikasi
- ✅ Aplikasi web/mobile dengan ancaman side-channel
- ✅ Compliance dengan standar keamanan modern
- ✅ Penyimpanan credential yang persistent

**Gunakan argon2.Key ketika:**
- ⚠️ Key derivation dalam environment terkontrol
- ⚠️ Performa sangat kritis dan tidak ada ancaman side-channel
- ⚠️ Backward compatibility dengan sistem legacy
- ❌ **JANGAN untuk password hashing**

#### Implementasi Production-Ready

```go
// BENAR: Menggunakan Argon2id untuk password
func HashPassword(password string) (string, error) {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    // Gunakan argon2.IDKey (Argon2id)
    hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 2, 32)
    
    encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
    encodedHash := base64.RawStdEncoding.EncodeToString(hash)
    
    return fmt.Sprintf("$argon2id$v=19$m=65536,t=3,p=2$%s$%s",
        encodedSalt, encodedHash), nil
}

// SALAH: Jangan gunakan argon2.Key untuk password
func WrongHashPassword(password string) []byte {
    salt := []byte("fixedSalt") // Bad practice
    // Argon2d tidak aman untuk password hashing
    return argon2.Key([]byte(password), salt, 3, 64*1024, 2, 32)
}
```

**Kesimpulan:** Selalu gunakan `argon2.IDKey` untuk password hashing. `argon2.Key` hanya untuk key derivation dalam konteks yang sangat spesifik dan aman.

---

## Logging di Go: Menggunakan Logrus

### Apa itu Logrus?

[Logrus](https://github.com/sirupsen/logrus) adalah library logging populer di Go yang menyediakan fitur logging yang lebih kaya dibandingkan `log` atau `fmt`. Logrus mendukung level log (info, warn, error, debug), format output (JSON, text), hooks, dan struktur log yang lebih baik.

### Kapan Harus Menggunakan Logrus?

Gunakan Logrus (atau library logging lain seperti Zap, Zerolog) ketika:
- Aplikasi sudah mulai kompleks (bukan sekadar CLI sederhana)
- Membutuhkan log terstruktur (misal: JSON untuk log aggregator/monitoring)
- Ingin membedakan level log (debug, info, warn, error, fatal)
- Perlu menulis log ke file, syslog, atau layanan eksternal
- Ingin menambah metadata (field) pada log (misal: user_id, request_id)

### Kelebihan Logrus dibanding Debugging dengan `fmt`

| Fitur                | `fmt.Println` / `log.Println` | Logrus (dan sejenisnya)      |
|----------------------|-------------------------------|------------------------------|
| Level log            | ❌ Tidak ada                  | ✅ Ada (Debug, Info, Warn, Error, Fatal, Panic) |
| Format output        | ❌ Plain text                 | ✅ Text/JSON/custom           |
| Field/metadata       | ❌ Tidak ada                  | ✅ Bisa tambah field (structured logging) |
| Output ke banyak tujuan | ❌ Sulit                   | ✅ Mudah (file, syslog, hook) |
| Filtering log        | ❌ Tidak bisa                 | ✅ Bisa filter per level      |
| Integrasi monitoring | ❌ Manual                     | ✅ Mudah diintegrasi          |
| Stack trace/error    | ❌ Manual                     | ✅ Built-in untuk error/fatal |

### Contoh Penggunaan Logrus

```go
import (
    log "github.com/sirupsen/logrus"
)

func main() {
    // Set format ke JSON (opsional)
    log.SetFormatter(&log.JSONFormatter{})

    // Set level log (misal: hanya tampilkan info ke atas)
    log.SetLevel(log.InfoLevel)

    log.Info("Server started")
    log.WithFields(log.Fields{
        "user_id": "USR-123",
        "action": "login",
    }).Warn("Suspicious login detected")

    log.Error("Database connection failed")
}
```

### Best Practice

- Gunakan log level sesuai kebutuhan (`Debug` untuk development, `Info` untuk event penting, `Warn` untuk potensi masalah, `Error` untuk error)
- Tambahkan field/metadata untuk memudahkan tracing (misal: request_id, user_id)
- Gunakan format JSON untuk aplikasi production (mudah di-parse oleh log aggregator)
- Jangan gunakan `fmt.Println` untuk logging di aplikasi production

---

**Kesimpulan:**  
Logrus (atau library logging lain) sangat direkomendasikan untuk aplikasi Go production karena memberikan fleksibilitas, kemudahan debugging, dan integrasi dengan sistem monitoring/log aggregator. Gunakan `fmt.Println` hanya untuk debugging sangat sederhana atau script sekali pakai.


## Menangani Panic dengan Middleware Log + Recover

Pada aplikasi Go berbasis web (seperti Echo), **panic** yang tidak tertangani dapat menyebabkan server crash dan menghentikan semua request. Untuk menjaga **konsistensi error** dan memastikan server tetap berjalan, gunakan middleware **Recover** yang juga melakukan logging error.

### Apa itu Panic?

- Panic terjadi saat aplikasi menemukan error fatal yang tidak bisa ditangani (misal: index out of range, nil pointer).
- Jika panic tidak di-recover, aplikasi akan crash.

### Solusi: Middleware Recover

Framework seperti Echo menyediakan middleware `Recover()` yang secara otomatis menangkap panic, mencegah crash, dan mengembalikan response error yang konsisten ke client.

#### Contoh Penggunaan di Echo

```go
import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    log "github.com/sirupsen/logrus"
)

func main() {
    e := echo.New()

    // Middleware Recover + Logging
    e.Use(middleware.Recover())
    e.Use(middleware.Logger())

    // Atau custom recover dengan logrus
    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            defer func() {
                if r := recover(); r != nil {
                    log.WithField("panic", r).Error("Recovered from panic")
                    c.JSON(500, map[string]string{"error": "internal server error"})
                }
            }()
            return next(c)
        }
    })

    // ... routes ...
    e.Logger.Fatal(e.Start(":8081"))
}
```

### Manfaat

- **Server tidak crash** meski terjadi panic.
- **Error tetap ter-log** (misal dengan Logrus) untuk keperluan debugging.
- **Response error konsisten** ke client (tidak bocor stack trace atau info sensitif).
- **Stabilitas aplikasi** lebih terjaga.

### Best Practice

- Selalu aktifkan middleware `Recover` di aplikasi production.
- Gabungkan dengan logging structured (Logrus/Zap) agar root cause mudah ditelusuri.
- Jangan gunakan panic untuk flow control biasa, hanya untuk error fatal.

---

**Kesimpulan:**  
Gunakan middleware log + recover untuk menangani panic secara terpusat, menjaga aplikasi tetap berjalan, dan memastikan error tercatat dengan baik.

## Cara Generate Mock Otomatis dengan mockgen

Selain membuat mock secara manual, Anda juga bisa menggunakan tool [mockgen](https://github.com/golang/mock) dari package `golang/mock` untuk menghasilkan kode mock secara otomatis dari interface Go.

### 1. Install mockgen

```sh
go install github.com/golang/mock/mockgen@latest
```

Pastikan `$GOPATH/bin` sudah ada di PATH Anda.

### 2. Generate Mock

Misal Anda punya interface di file `internal/repository/catalog/repository.go`:

```go
type CatalogRepository interface {
    GetAll() ([]ProductClothes, error)
    GetByID(id string) (*ProductClothes, error)
    Create(product *ProductClothes) error
}
```

Jalankan perintah berikut di terminal untuk generate mock-nya:

```sh
mockgen -source=internal/repository/catalog/repository.go -destination=internal/repository/catalog/mock_repository.go -package=catalog
```

- `-source`: file interface asli
- `-destination`: file output mock
- `-package`: nama package untuk file mock

### 3. Contoh Penggunaan di Test

Import mock yang sudah di-generate, lalu gunakan di unit test Anda:

```go
import (
    "testing"
    "github.com/golang/mock/gomock"
    "yourmodule/internal/repository/catalog"
)

func TestGetAllCatalog(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := catalog.NewMockCatalogRepository(ctrl)
    mockRepo.EXPECT().GetAll().Return([]catalog.ProductClothes{
        {UniqueID: "PRD-1", NamaPakaian: "Kaos"},
    }, nil)

    usecase := NewCatalogUsecase(mockRepo)
    result, err := usecase.GetAll()
    require.NoError(t, err)
    require.Len(t, result, 1)
}
```

### Tips

- Gunakan mockgen untuk interface yang sering berubah atau kompleks.
- mockgen bisa diintegrasikan ke Makefile atau script build/test Anda.

---

# Unit Test dan Mocking untuk RESTful API Go Catalog

File ini berisi contoh dan best practice untuk penulisan unit test di proyek Go Catalog, termasuk penggunaan **gomock** untuk mocking dependency.

## Struktur Test

- Semua unit test sebaiknya ditempatkan di file dengan suffix `_test.go` (misal: `catalog_handler_test.go`, `order_usecase_test.go`).
- Untuk interface, gunakan **gomock** untuk generate mock otomatis.
- Test file ini hanya contoh, implementasi sesuaikan dengan struktur dan interface di proyek Anda.

---

## 1. Install Dependency Test

Install package testing dan mock:
```sh
go get github.com/golang/mock/gomock
go get github.com/stretchr/testify
```

## 2. Generate Mock Interface

Gunakan mockgen untuk membuat mock dari interface repository atau usecase:
```sh
mockgen -source=internal/repository/catalog/repository.go -destination=internal/repository/catalog/mock_repository.go -package=catalog
```

## 3. Contoh Unit Test Handler

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/labstack/echo/v4"
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/require"
    "yourmodule/internal/repository/catalog"
)

func TestGetAllClothesHandler(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := catalog.NewMockCatalogRepository(ctrl)
    mockRepo.EXPECT().GetAll().Return([]catalog.ProductClothes{
        {UniqueID: "PRD-1", NamaPakaian: "Kaos"},
    }, nil)

    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/clothes", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    handler := NewCatalogHandler(mockRepo)
    err := handler.GetAllClothes(c)

    require.NoError(t, err)
    require.Equal(t, http.StatusOK, rec.Code)
    // Tambahkan assertion sesuai response
}
```

## 4. Best Practice Testing

- Mock dependency (repository, usecase) agar test terisolasi.
- Gunakan `require` atau `assert` dari testify untuk assertion.
- Test semua skenario: sukses, error, validasi gagal.
- Jalankan test dengan `go test ./...`.

---

**Kesimpulan:**  
Unit test dengan mocking memastikan logika aplikasi teruji tanpa tergantung database atau dependency eksternal. Gunakan mockgen dan testify untuk test yang robust dan mudah dipelihara.

## Behavior Driven Development (BDD) di Go

### Apa itu BDD?

**BDD (Behavior Driven Development)** adalah pendekatan pengembangan perangkat lunak yang berfokus pada perilaku aplikasi dari sudut pandang pengguna. BDD menulis spesifikasi dalam bentuk skenario yang mudah dipahami, biasanya menggunakan format Given-When-Then, sehingga komunikasi antara developer, QA, dan stakeholder menjadi lebih jelas.

### Seperti Apa BDD?

- Spesifikasi ditulis sebagai **skenario**:  
    - **Given** (dengan kondisi awal tertentu)
    - **When** (ketika aksi dilakukan)
    - **Then** (maka hasil yang diharapkan)
- Contoh skenario BDD:
    ```
    Given user belum terdaftar
    When user melakukan registrasi dengan data valid
    Then user berhasil terdaftar dan mendapatkan unique ID
    ```

### Implementasi BDD di Go

Di Go, BDD dapat diimplementasikan menggunakan framework seperti [Ginkgo](https://github.com/onsi/ginkgo) dan [Gomega](https://github.com/onsi/gomega).

#### 1. Install Ginkgo & Gomega

```sh
go get github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
```

#### 2. Contoh Test BDD dengan Ginkgo

```go
package catalog_test

import (
        . "github.com/onsi/ginkgo/v2"
        . "github.com/onsi/gomega"
)

var _ = Describe("Catalog API", func() {
        Context("Ketika user mengambil semua produk", func() {
                It("harus mengembalikan list produk", func() {
                        products := GetAllProducts()
                        Expect(products).NotTo(BeEmpty())
                })
        })

        Context("Ketika user menambah produk baru", func() {
                It("harus berhasil menambah produk", func() {
                        err := AddProduct(Product{Nama: "Kaos", Price: 100000})
                        Expect(err).To(BeNil())
                })
        })
})
```

#### 3. Jalankan Test BDD

```sh
ginkgo -v
```

### Kegunaan BDD

- **Meningkatkan komunikasi** antara tim teknis dan non-teknis
- **Dokumentasi hidup**: test adalah spesifikasi perilaku aplikasi
- **Membantu validasi fitur**: setiap fitur diuji dari sisi perilaku
- **Memudahkan refactoring**: test BDD memastikan perilaku tetap konsisten

---

**Kesimpulan:**  
BDD di Go (dengan Ginkgo/Gomega) membantu menulis test yang mudah dipahami, menjaga aplikasi tetap sesuai dengan kebutuhan bisnis, dan meningkatkan kualitas serta kolaborasi tim.