package order

import (
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}


func (or *orderRepo) CreateOrder(order models.Order) (models.Order, error) {
	// Func buat create order
	if err := or.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (or *orderRepo) GetInfoOrder(orderID string) (models.Order, error) {
	var orderData models.Order // Ambil data order berdasarkan orderID
	if err := or.db.Where("id = ?", orderID).Preload("ProductOrder").First(&orderData).Error; err != nil {
		return orderData, err
	}
	return orderData, nil
}

func (or *orderRepo) GetAllOrder(order models.Order) (models.Order, error) {
	var orderData models.Order

	// Ambil semua data order
	if err := or.db.Preload("ProductOrder").Find(&orderData).Error; err != nil {
		return orderData, err
	}
	return orderData, nil
}