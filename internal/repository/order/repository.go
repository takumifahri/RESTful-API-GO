package order


import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrder(order models.Order) (models.Order, error)
	GetInfoOrder(orderID string) (models.Order, error)
}

