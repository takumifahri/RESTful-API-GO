# Generics dalam Pemrograman

## Menambahkan Expired Time pada Caching (Contoh dengan Library bawaan Go)

Jika ingin menambahkan expired time pada cache tanpa library eksternal, Anda bisa membuat cache sederhana menggunakan map dan menyimpan waktu kedaluwarsa (expired) untuk setiap item. Berikut contoh implementasinya:

```go
import (
    "sync"
    "time"
)

type cacheItem[T any] struct {
    value      T
    expiration int64 // waktu kedaluwarsa dalam Unix timestamp (detik)
}

type Cache[T any] struct {
    data map[string]cacheItem[T]
    mu   sync.RWMutex
}

func NewCache[T any]() *Cache[T] {
    return &Cache[T]{data: make(map[string]cacheItem[T])}
}

func (c *Cache[T]) Set(key string, value T, duration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = cacheItem[T]{
        value:      value,
        expiration: time.Now().Add(duration).Unix(),
    }
}

func (c *Cache[T]) Get(key string) (T, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    item, found := c.data[key]
    if !found || time.Now().Unix() > item.expiration {
        var zero T
        return zero, false
    }
    return item.value, true
}
```

Dengan cara ini, setiap data yang disimpan di cache akan memiliki waktu expired sesuai
    if a > b {
        return a
    }
    return b
}
```

### 3. Repository Pattern pada Database

Pada aplikasi POS (Point of Sale), Anda bisa membuat repository generic untuk operasi CRUD pada berbagai entitas (produk, pelanggan, transaksi).

```go
type Repository[T any] interface {
    Create(item T) error
    GetByID(id int) (T, error)
    Update(item T) error
    Delete(id int) error
}
```

## Kesimpulan

Generics sangat berguna untuk membuat kode yang lebih efisien, reusable, dan aman. Fitur ini sangat membantu dalam pengembangan aplikasi skala besar seperti sistem POS, sistem inventory, dan lain-lain.

# Caching dalam Pemrograman

## Apa itu Caching?

Caching adalah teknik penyimpanan data sementara di memori agar data tersebut dapat diakses lebih cepat saat dibutuhkan kembali, tanpa harus melakukan proses komputasi atau pengambilan data yang sama berulang kali.


## Note : 
Caching biasanya digunakan dengan random access memory.
## Kegunaan Caching

- **Meningkatkan Performa:** Mengurangi waktu akses data yang sering digunakan.
- **Mengurangi Beban Sistem:** Meminimalisir akses ke sumber data utama (misal: database, API).
- **Efisiensi Resource:** Menghemat bandwidth dan resource komputasi.

## Kapan Harus Menggunakan Caching?

Gunakan caching ketika:
- Data sering diakses berulang kali dalam waktu singkat.
- Proses pengambilan data dari sumber utama memakan waktu atau resource besar.
- Konsistensi data tidak harus selalu real-time.

## Contoh Real Case di Dunia Kerja

### 1. Caching pada API Response

Misal, aplikasi POS sering mengambil daftar produk dari API. Dengan caching, daftar produk bisa disimpan sementara di memori.

```go
type Cache[T any] struct {
    data map[string]T
}

func NewCache[T any]() *Cache[T] {
    return &Cache[T]{data: make(map[string]T)}
}

func (c *Cache[T]) Get(key string) (T, bool) {
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache[T]) Set(key string, value T) {
    c.data[key] = value
}
```

### 2. Caching Query Database

Pada aplikasi dengan query database yang berat, hasil query bisa disimpan di cache untuk mengurangi beban database.

### 3. Caching pada Layer Service

Service yang sering memproses data yang sama dapat menggunakan cache untuk mempercepat respon ke user.

## Kesimpulan

Caching sangat penting untuk meningkatkan performa aplikasi, terutama pada sistem yang membutuhkan akses data cepat dan efisien seperti POS, e-commerce, dan aplikasi berbasis web lainnya.
## Cara Menambahkan Expired Time pada Caching dengan Library Go

Jika Anda ingin menggunakan expired time pada caching dengan library eksternal di Go, Anda bisa memakai library seperti [`patrickmn/go-cache`](https://github.com/patrickmn/go-cache). Library ini sudah menyediakan fitur expired time (TTL) dan pembersihan otomatis item yang sudah kedaluwarsa.

### Contoh Penggunaan `go-cache`:

```go
import (
    "github.com/patrickmn/go-cache"
    "time"
)

func main() {
    // Membuat cache dengan default expiration 5 menit, dan pembersihan setiap 10 menit
    c := cache.New(5*time.Minute, 10*time.Minute)

    // Menyimpan data dengan expired time 1 menit
    c.Set("key", "value", 1*time.Minute)

    // Mengambil data dari cache
    value, found := c.Get("key")
    if found {
        // gunakan value
    } else {
        // data sudah expired atau tidak ditemukan
    }
}
```

Dengan library ini, Anda tidak perlu mengelola expired time secara manual karena sudah di-handle otomatis oleh library.