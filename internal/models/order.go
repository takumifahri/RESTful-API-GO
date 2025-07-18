package models

type OrderStatus string

type Order struct {
	ID          	uint `gorm:"primaryKey"`
	UNIQUEID		string
	ProductOrder 	[]ProductOrder
	Status 			OrderStatus
}

type ProductORderStatus string

type ProductOrder struct {
	ID          	uint `gorm:"primaryKey"`
	OrderID	 		string
	ProductID   	string
	NamaPakaian 	string
	TotalPrice 		int64
	Quantity   		int64
	Status     		ProductORderStatus
}

type OrderMenuProductRequest struct {
	OrderCode 		string 
	Quantity 		int64
}

type OrderMenuRequest struct {
	OrderProduct 	[]OrderMenuProductRequest
}

type GetOrderInfoRequest struct {
	OrderID 		string
}