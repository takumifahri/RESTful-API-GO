package models

// import "github.com/google/uuid"

type tipe string

type ProductClothes struct {
    ID          uint                 `gorm:"primaryKey" json:"id"`
    UNIQUEID    string               `gorm:"unique;not null;size:255" json:"unique_id"`
    NamaPakaian string               `json:"nama_pakaian" validate:"required,min=3,max=100"`
    Price       int64                `json:"price" validate:"required,min=1"`
    Deskripsi   string               `json:"deskripsi" validate:"required,min=10,max=500"`
    Stock       int                  `gorm:"default:0" json:"stock" validate:"min=0"`
    TypeClothes tipe 				 `gorm:"type:varchar(20)" json:"type_clothes" validate:"required,oneof=shirt pants outerwear accessories shoes"`
}