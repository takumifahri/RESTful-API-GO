package store

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Usecase interface {
	GetAllCatalogList(tipe string) ([]models.ProductClothes, error)
	Order(request models.OrderMenuRequest) (models.Order, error)
	GetOrderInfo(request models.GetOrderInfoRequest) (models.Order, error)
}