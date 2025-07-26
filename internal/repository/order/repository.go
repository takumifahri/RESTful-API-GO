package order


import (
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"context"
)
//go:generate mockgen -package=mocks -mock_names=Repository=MockOrderRepository -destination=../../mocks/order_repository_mock.go -source=repository.go

type Repository interface {
	CreateOrder(ctx context.Context, order models.Order) (models.Order, error)
	GetAllOrder(ctx context.Context, order models.Order) (models.Order, error)
	GetInfoOrder(ctx context.Context, orderID string) (models.Order, error)
  	AdminGetAllOrder(ctx context.Context) ([]models.Order, error) // ✅ Tidak perlu parameter
}

