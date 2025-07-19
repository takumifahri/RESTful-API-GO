package models
type OrderStatus string

type Order struct {
    ID          	uint `gorm:"primaryKey" json:"id"`
    UNIQUEID		string `gorm:"unique;not null;size:255" json:"unique_id"`
    ProductOrder 	[]ProductOrder `json:"product_order"`
    Status 			OrderStatus `json:"status"`
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
    OrderProduct 	[]OrderMenuProductRequest `json:"order_product"`
}

type GetOrderInfoRequest struct {
    OrderID 		string `json:"order_id"`
}