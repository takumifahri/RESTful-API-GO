package catalog

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	GetAllCatalogList(tipe string) ([]models.ProductClothes, error)
	GetAllCatalog(orderCode string) (models.ProductClothes, error)

}