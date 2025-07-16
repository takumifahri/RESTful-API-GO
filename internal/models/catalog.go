package models

type tipe string

type ProductClothes struct {
	ID          uint `gorm:"primaryKey"`
	NamaPakaian string
	Price       int64
	Deskripsi   string
	TypeClothes tipe

}