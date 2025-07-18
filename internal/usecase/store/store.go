package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/order"
)

type storeUsecase struct {
	menuRepo catalog.Repository
	orderRepo order.Repository
}

func GetUsecase(menuyRepo catalog.Repository, orderRepo order.Repository) Usecase {
	return &storeUsecase{
		menuRepo: menuyRepo,
		orderRepo: orderRepo,
	}
}

func (s *storeUsecase) GetAllCatalogList(tipe string) ([]models.ProductClothes, error) {
	return s.menuRepo.GetAllCatalogList(tipe)
}
func (s *storeUsecase) GetCatalogByID(UNIQUEID string) (*models.ProductClothes, error) {
	catalogData, err := s.menuRepo.GetCatalogByID(UNIQUEID)
	if err != nil {
		return nil, err
	}
	if catalogData == nil {
		return nil, nil // Return nil if no catalog found
	}
	return catalogData, nil
}

func (s *storeUsecase) AddCatalog(catalog models.ProductClothes) (models.ProductClothes, error) {
    // 1. Generate UUID (ID akan di-handle oleh database)
    catalog.UNIQUEID = fmt.Sprintf("PRD-%s", uuid.New().String())
    
    // 2. Save ke repository
    if err := s.menuRepo.CreateCatalog(catalog); err != nil {
        return models.ProductClothes{}, err
    }
    
    // 3. PENTING: Query kembali data yang baru saja disimpan untuk mendapatkan ID yang benar
    savedCatalog, err := s.menuRepo.GetCatalogByID(catalog.UNIQUEID)
    if err != nil {
        return models.ProductClothes{}, err
    }
    
    return *savedCatalog, nil
}

func (s *storeUsecase) Order(request models.OrderMenuRequest) (models.Order, error) {
	productOrderData := make([]models.ProductOrder, len(request.OrderProduct))

	//  Kita loop 
	for i, orderProduct := range request.OrderProduct {
		menuData, err := s.menuRepo.GetAllCatalog(orderProduct.OrderCode)
		if err != nil {
			return models.Order{}, err
		}
	
		productOrderData[i] = models.ProductOrder{
			OrderID: uuid.New().String(),
			ProductID: menuData.UNIQUEID,
			Quantity: orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status: constant.ProductOrderStatusPending,
		}
	}

	orderData := models.Order{
		UNIQUEID: uuid.New().String(),
		Status: constant.OrderStatusPending,
		ProductOrder: productOrderData,
	}

	createData, err := s.orderRepo.CreateOrder(orderData)
	if err != nil {
		return models.Order{}, err
	}
	return createData, nil 
}

func (s *storeUsecase) GetOrderInfo(request models.GetOrderInfoRequest) (models.Order, error) {
	orderData, err := s.orderRepo.GetInfoOrder(request.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

