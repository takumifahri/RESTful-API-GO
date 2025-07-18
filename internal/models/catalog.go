package models

// import "github.com/google/uuid"

type tipe string

type ProductClothes struct {
	ID          uint      `gorm:"primaryKey"`
    UNIQUEID    string   
	NamaPakaian string
	Price       int64
	Deskripsi   string
	Stock       int       `gorm:"default:0"`
	TypeClothes tipe
}