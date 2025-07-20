package store

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Usecase interface {
	GetAllCatalogList(tipe string) ([]models.ProductClothes, error)
	GetCatalogByID(UNIQUEID string) (*models.ProductClothes, error)
	AddCatalog(catalog models.ProductClothes) (models.ProductClothes, error)
	UpdateCatalog(catalog models.ProductClothes) (models.ProductClothes, error)
	Order(request models.OrderMenuRequest) (models.Order, error)
	GetOrderInfo(request models.GetOrderInfoRequest) (models.Order, error)
 	AdminGetAllOrder() ([]models.Order, error) // ✅ Tidak perlu parameter
}