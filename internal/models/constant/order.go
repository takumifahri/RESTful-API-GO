package constant

import "github.com/takumifahri/RESTful-API-GO/internal/models"

const (
	OrderStatusPending    	models.OrderStatus 		= "pending"
	OrderStatusProcessing  	models.OrderStatus		= "processing"
	OrderStatusCompleted   	models.OrderStatus     	= "completed"
	OrderStatusCancelled    models.OrderStatus    	= "cancelled"
	OrderStatusFailed       models.OrderStatus    	= "failed"
)

const (
	ProductOrderStatusPending    models.ProductOrderStatus = "pending"
	ProductOrderStatusProcessing models.ProductOrderStatus = "processing"
	ProductOrderStatusCompleted  models.ProductOrderStatus = "completed"
	ProductOrderStatusCancelled  models.ProductOrderStatus = "cancelled"
	ProductOrderStatusFailed     models.ProductOrderStatus = "failed"
)