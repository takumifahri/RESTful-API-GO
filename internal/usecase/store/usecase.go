package store

import (
	"context"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
)


type Usecase interface {
	GetAllCatalogList(ctx context.Context, tipe string) ([]models.ProductClothes, error)
	GetCatalogByID(ctx context.Context, UNIQUEID string) (*models.ProductClothes, error)
	AddCatalog(ctx context.Context, catalog models.ProductClothes) (models.ProductClothes, error)
	UpdateCatalog(ctx context.Context, catalog models.ProductClothes) (models.ProductClothes, error)
	Order(ctx context.Context, request models.OrderMenuRequest) (models.Order, error)
	GetOrderInfo(ctx context.Context, request models.GetOrderInfoRequest) (models.Order, error)
 	AdminGetAllOrder(ctx context.Context) ([]models.Order, error) // ✅ Tidak perlu parameter
}