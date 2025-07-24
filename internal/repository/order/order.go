package order

import (
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"gorm.io/gorm"
	"context"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}


func (or *orderRepo) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "CreateOrder")
	defer span.End()
	// Func buat create order
	if err := or.db.WithContext(ctx).Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (or *orderRepo) GetInfoOrder(ctx context.Context, orderID string) (models.Order, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "GetInfoOrder")
	defer span.End()
	var orderData models.Order // Ambil data order berdasarkan orderID
	if err := or.db.WithContext(ctx).Where("unique_id = ?", orderID).Preload("ProductOrder").First(&orderData).Error; err != nil {
		return orderData, err
	}
	return orderData, nil
}

func (or *orderRepo) GetAllOrder(ctx context.Context, order models.Order) (models.Order, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "GetAllOrder")	
	defer span.End()

	var orderData models.Order

	// Ambil semua data order
	if err := or.db.WithContext(ctx).Preload("ProductOrder").Find(&orderData).Error; err != nil {
		return orderData, err
	}
	return orderData, nil
}


func (or *orderRepo) AdminGetAllOrder(ctx context.Context) ([]models.Order, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "AdminGetAllOrder")
	defer span.End()

    var orders []models.Order
    
    // Ambil semua order dengan preload ProductOrder
    if err := or.db.WithContext(ctx).Preload("ProductOrder").Find(&orders).Error; err != nil {
        return nil, err
    }
    
    return orders, nil
}
