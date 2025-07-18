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
	ProductOrderStatusPending    models.ProductORderStatus = "pending"
	ProductOrderStatusProcessing models.ProductORderStatus = "processing"
	ProductOrderStatusCompleted  models.ProductORderStatus = "completed"
	ProductOrderStatusCancelled  models.ProductORderStatus = "cancelled"
	ProductOrderStatusFailed     models.ProductORderStatus = "failed"
)