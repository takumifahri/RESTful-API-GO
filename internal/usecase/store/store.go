package store

import (
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
)

type storeUsecase struct {
	menuRepo catalog.Repository
}

func GetUsecase(menuyRepo catalog.Repository) Usecase {
	return &storeUsecase{
		menuRepo: menuyRepo,
	}
}

func (s *storeUsecase) GetAllCatalog(tipe string) ([]models.ProductClothes, error) {
	return s.menuRepo.GetAllCatalog(tipe)
}