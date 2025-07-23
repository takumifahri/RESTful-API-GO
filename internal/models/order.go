package models

type OrderStatus string

type Order struct {
    ID          	uint `gorm:"primaryKey" json:"id"`
    UserUniqueID 	string `json:"user_unique_id"`
    UNIQUEID		string `gorm:"unique;not null;size:255" json:"unique_id"`
    ProductOrder 	[]ProductOrder `json:"product_order"`
    Status 			OrderStatus `json:"status"`
    ReferenceID     string `gorm:"unique;not null;size:255;" json:"reference_id"` // ini mencegah generals problem
}

// Fix typo: ProductORderStatus -> ProductOrderStatus
type ProductOrderStatus string

type ProductOrder struct {
    ID          	uint `gorm:"primaryKey" json:"id"`
    OrderID     	uint `gorm:"not null" json:"order_id"`
    OrderUniqueID	string `json:"order_unique_id"`    // Rename untuk clarity dan reference ke Order.UNIQUEID
    ProductID   	string `json:"product_id"`         // UNIQUEID dari ProductClothes
    NamaPakaian 	string `json:"nama_pakaian"`
    TotalPrice 		int64 `json:"total_price"`
    Quantity   		int64 `json:"quantity"`            // Ubah ke int64 untuk konsistensi
    Status     		ProductOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
    ProductID 		string `json:"product_id"`
    Quantity 		int64 `json:"quantity"`            // Ubah ke int64
}

type OrderMenuRequest struct {
    UserUniqueID    string `json:"-"`
    OrderProduct 	[]OrderMenuProductRequest `json:"order_product"`
    ReferenceID     string `json:"reference_id"` // Tambahkan reference ID untuk order
}

type GetOrderInfoRequest struct {
    UserUniqueID    string `json:"-"`
    OrderID 		string `json:"order_id"`
}

type GetAllOrderRequest struct {
    // Tambahkan field yang diperlukan untuk mendapatkan semua order
    ProductID       string `json:"product_id"` // Misalnya, jika ingin filter berdasarkan ProductID
    ReferenceID     string `json:"reference_id"` // Tambahkan reference ID untuk order
    Quantity        int64 `json:"quantity"` // Tambahkan quantity jika diperlukan
    TotalPrice      int64 `json:"total_price"` // Tambahkan total price jika diperlukan
    OrderStatus     OrderStatus `json:"order_status"` // Tambahkan status order jika diperlukan
}