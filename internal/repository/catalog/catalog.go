package catalog

import (
	"errors"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetAllCatalogList(tipe string) ([]models.ProductClothes, error) {
    var catalogData []models.ProductClothes

    // Mulai dengan query dasar tanpa filter
    query := m.db

    // Jika parameter 'tipe' tidak kosong, tambahkan filter WHERE
    if tipe != "" {
        query = query.Where("type_clothes = ?", tipe)
    }

    // Eksekusi query yang sudah dibentuk
    if err := query.Find(&catalogData).Error; err != nil {
        return nil, err
    }
    return catalogData, nil
}

func (m *menuRepo) GetAllCatalog(orderCode string) (models.ProductClothes, error) {
    var productData models.ProductClothes

    // Ambil data produk berdasarkan orderCode
    if err := m.db.Where("unique_id = ?", orderCode).First(&productData).Error; err != nil {
        return productData, err
    }
    return productData, nil
}
func (m *menuRepo) GetCatalogByID(UNIQUEID string) (*models.ProductClothes, error) {
    var catalogData models.ProductClothes

    // Ambil data produk berdasarkan UNIQUEID
    if err := m.db.Where("unique_id = ?", UNIQUEID).First(&catalogData).Error; err != nil {
        return nil, err
    }
    return &catalogData, nil
}

func (m *menuRepo) CreateCatalog(catalog models.ProductClothes) error {
    // Simpan data produk ke database
    if err := m.db.Create(&catalog).Error; err != nil {
        return err
    }
    return nil
}

func (m *menuRepo) UpdateCatalog(uniqueID string, updateData map[string]interface{}) error {
    // GORM akan update hanya field yang ada di map
    result := m.db.Model(&models.ProductClothes{}).Where("unique_id = ?", uniqueID).Updates(updateData)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("catalog not found")
    }
    return nil
}
