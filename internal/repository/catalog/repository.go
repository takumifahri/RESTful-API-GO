package catalog

import (
	"context"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

type Repository interface {
	GetAllCatalogList(ctx context.Context, tipe string) ([]models.ProductClothes, error)
	GetAllCatalog(ctx context.Context, orderCode string) (models.ProductClothes, error)
	GetCatalogByID(ctx context.Context, UNIQUEID string) (*models.ProductClothes, error)
	CreateCatalog(ctx context.Context, catalog models.ProductClothes) error
	UpdateCatalog(ctx context.Context, uniqueID string, updateData map[string]interface{}) error
}