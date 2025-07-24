package catalog

import (
	"errors"
    "context"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
    "github.com/takumifahri/RESTful-API-GO/internal/tracing"
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

func (m *menuRepo) GetAllCatalogList(ctx context.Context, tipe string) ([]models.ProductClothes, error) {
    ctx, span := tracing.CreateSpanWrapper(ctx, "GetAllCatalogList")
    defer span.End()
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

func (m *menuRepo) GetAllCatalog(ctx context.Context, orderCode string) (models.ProductClothes, error) {
    ctx, span := tracing.CreateSpanWrapper(ctx, "GetAllCatalog")
    defer span.End()

    var productData models.ProductClothes

    // Ambil data produk berdasarkan orderCode
    if err := m.db.WithContext(ctx).Where("unique_id = ?", orderCode).First(&productData).Error; err != nil {
        return productData, err
    }
    return productData, nil
}
func (m *menuRepo) GetCatalogByID(ctx context.Context, UNIQUEID string) (*models.ProductClothes, error) {
    ctx, span := tracing.CreateSpanWrapper(ctx, "GetCatalogByID")
    defer span.End()

    var catalogData models.ProductClothes

    // Ambil data produk berdasarkan UNIQUEID
    if err := m.db.WithContext(ctx).Where("unique_id = ?", UNIQUEID).First(&catalogData).Error; err != nil {
        return nil, err
    }
    return &catalogData, nil
}

func (m *menuRepo) CreateCatalog(ctx context.Context, catalog models.ProductClothes) error {
    ctx, span := tracing.CreateSpanWrapper(ctx, "CreateCatalog")
    defer span.End()
    // Simpan data produk ke database
    if err := m.db.WithContext(ctx).Create(&catalog).Error; err != nil {
        return err
    }
    return nil
}

func (m *menuRepo) UpdateCatalog(ctx context.Context, uniqueID string, updateData map[string]interface{}) error {
    ctx, span := tracing.CreateSpanWrapper(ctx, "UpdateCatalog")
    defer span.End()
    // GORM akan update hanya field yang ada di map
    result := m.db.WithContext(ctx).Model(&models.ProductClothes{}).Where("unique_id = ?", uniqueID).Updates(updateData)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("catalog not found")
    }
    return nil
}
