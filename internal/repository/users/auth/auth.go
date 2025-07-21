package auth

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"gorm.io/gorm"
)

type authRepo struct {
	db      *gorm.DB
	gcm 	cipher.AEAD
	time    uint32
	memory  uint32
	threads uint8
	keylen  uint32
}

func GetRepository(
	db *gorm.DB,
	secret string,
	time uint32,
	memory uint32,
	threads uint8,
	keylen uint32,
) (Repository, error) {
	// TODO: implement AES-GCM cipher initialization with secret
	block , err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &authRepo{
		db:      db,
		gcm:     gcm,
		time:    time,
		memory:  memory,
		threads: threads,
		keylen:  keylen,
	}, nil
}

func (au *authRepo) RegisterUser(userData models.User) (models.User, error) {
	if err := au.db.Create(&userData).Error; err != nil {
		return models.User{}, err
	}
	return userData, nil
}

func (au *authRepo) CheckRegistered(Name string) (bool, error) {
	// ktia ambil data kosongan dari user yakni
	var userData models.User
	if err := au.db.Where("name = ?", Name).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User tidak ditemukan
		}
		return false, err // Terjadi error lain
	}

	return userData.UniqueID != "", nil // Jika UniqueID tidak kosong, berarti user terdaftar
}
