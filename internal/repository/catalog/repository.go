package catalog

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	GetAllCatalogList(tipe string) ([]models.ProductClothes, error)
	GetAllCatalog(orderCode string) (models.ProductClothes, error)
	GetCatalogByID(UNIQUEID string) (*models.ProductClothes, error)
	CreateCatalog(catalog models.ProductClothes) error
}