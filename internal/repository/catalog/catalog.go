package catalog

import (
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

func (m *menuRepo) GetAllCatalog(tipe string) ([]models.ProductClothes, error) {
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