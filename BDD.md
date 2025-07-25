
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