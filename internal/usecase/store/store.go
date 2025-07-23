package store

import (
	"errors"
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

func (s *storeUsecase) UpdateCatalog(catalog models.ProductClothes) (models.ProductClothes, error) {
	// 1. Cek apakah catalog exist
	existingCatalog, err := s.menuRepo.GetCatalogByID(catalog.UNIQUEID)
	if err != nil {
		return models.ProductClothes{}, err
	}
	if existingCatalog == nil {
		return models.ProductClothes{}, errors.New("catalog not found")
	}

	// 2. Update data
	updateData := make(map[string]interface{})
	updateData["nama_pakaian"] = catalog.NamaPakaian
	updateData["price"] = catalog.Price
	updateData["deskripsi"] = catalog.Deskripsi
	updateData["stock"] = catalog.Stock
	updateData["type_clothes"] = catalog.TypeClothes

	// 3. Update di repository
	if err := s.menuRepo.UpdateCatalog(catalog.UNIQUEID, updateData); err != nil {
		return models.ProductClothes{}, err
	}

	// 4. Return updated data
	updatedCatalog, err := s.menuRepo.GetCatalogByID(catalog.UNIQUEID)
	if err != nil {
		return models.ProductClothes{}, err
	}

	return *updatedCatalog, nil
}

func (s *storeUsecase) Order(request models.OrderMenuRequest) (models.Order, error) {
    productOrderData := make([]models.ProductOrder, len(request.OrderProduct))

    // Generate Order UNIQUEID sekali di awal
    orderUniqueID := "ORD-" + uuid.New().String()

    //  Kita loop 
    for i, orderProduct := range request.OrderProduct {
        // Gunakan GetCatalogByID untuk UNIQUEID
        menuData, err := s.menuRepo.GetCatalogByID(orderProduct.ProductID)
        if err != nil {
            return models.Order{}, err
        }
        
        // Cek apakah catalog ditemukan
        if menuData == nil {
            return models.Order{}, fmt.Errorf("product with ID %s not found", orderProduct.ProductID)
        }
    
        productOrderData[i] = models.ProductOrder{
            OrderUniqueID: orderUniqueID,              // ✅ GUNAKAN OrderUNIQUEID, bukan OrderID
            ProductID: menuData.UNIQUEID,              
            NamaPakaian: menuData.NamaPakaian,         
            Quantity: int64(orderProduct.Quantity),    // ✅ Convert ke int64
            TotalPrice: menuData.Price * orderProduct.Quantity,
            Status: constant.ProductOrderStatusPending,
        }
    }

    orderData := models.Order{
        UNIQUEID: orderUniqueID,      
		UserUniqueID: request.UserUniqueID,                 
        Status: constant.OrderStatusPending,
        ProductOrder: productOrderData,
		ReferenceID: request.ReferenceID,
    }
	if orderData.ReferenceID == "" {
		return models.Order{}, errors.New("reference ID is required")
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
	if orderData.UserUniqueID != request.UserUniqueID {
		return models.Order{}, fmt.Errorf("order with ID %s does not belong to user %s", request.OrderID, request.UserUniqueID)
	}
	return orderData, nil
}

func (s *storeUsecase) AdminGetAllOrder() ([]models.Order, error) {
    // Tidak perlu parameter, langsung ambil semua
    orderData, err := s.orderRepo.AdminGetAllOrder()
    if err != nil {
        return nil, err
    }

    return orderData, nil
}

