package store

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Usecase interface {
	GetAllCatalog(tipe string) ([]models.ProductClothes, error)
}