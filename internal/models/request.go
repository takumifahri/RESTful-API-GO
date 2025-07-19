package models

type tipeRequest string
// UpdateProductClothesRequest untuk update operations
// Menggunakan pointer agar bisa detect nil vs zero value
type UpdateProductClothesRequest struct {
    NamaPakaian *string                `json:"nama_pakaian,omitempty" validate:"omitempty,min=3,max=100"`
    Price       *int64                 `json:"price,omitempty" validate:"omitempty,min=1"`
    Deskripsi   *string                `json:"deskripsi,omitempty" validate:"omitempty,min=10,max=500"`
    Stock       *int                   `json:"stock,omitempty" validate:"omitempty,min=0"`
    TypeClothes *tipeRequest  `json:"type_clothes,omitempty" validate:"omitempty,oneof=shirt pants outerwear accessories shoes"`
}