package catalog

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	GetAllCatalog(tipe string) ([]models.ProductClothes, error)
}