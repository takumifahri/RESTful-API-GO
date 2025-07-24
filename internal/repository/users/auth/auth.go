package auth

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
	"gorm.io/gorm"
)

type authRepo struct {
	db      *gorm.DB
	gcm 	cipher.AEAD
	time    uint32
	memory  uint32
	threads uint8
	keylen  uint32
	signKey  *rsa.PrivateKey
	expTime time.Duration
}

func GetRepository(
	db *gorm.DB,
	secret string,
	time uint32,
	memory uint32,
	threads uint8,
	keylen uint32,
	signKey *rsa.PrivateKey,
	expTime time.Duration,
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
		signKey: signKey,
		expTime: expTime,
	}, nil
}

func (au *authRepo) RegisterUser(ctx context.Context, userData models.User) (models.User, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "RegisterUser")
	defer span.End()

	if err := au.db.WithContext(ctx).Create(&userData).Error; err != nil {
		return models.User{}, err
	}
	return userData, nil
}

func (au *authRepo) CheckRegistered(ctx context.Context, Email string) (bool, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "CheckRegistered")
	defer span.End()
	// kita ambil data kosongan dari user yakni
	var userData models.User
	if err := au.db.WithContext(ctx).Where("email = ?", Email).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User tidak ditemukan
		}
		return false, err // Terjadi error lain
	}

	return userData.UniqueID != "", nil // Jika UniqueID tidak kosong, berarti user terdaftar
}

// func (au *authRepo) GetMe(Email string) (models.User, error) {
// 	var userData models.User
// 	if err := au.db.Where("email = ?", Email).First(&userData).Error; err != nil {
// 		return userData, err
// 	}

// 	return userData, nil
// }

// func (au *authRepo) VerifyUserLogin(Email, Password string, userData models.User) (bool, error) {
// 	// Cek email apakah ada
// 	if Email != userData.Email {
// 		return false, nil
// 	}

// 	// Hashed atau enkripsi si password apakah sama atau tidak
// 	verified, err := au.comparePassword(Password, userData.Hash)
// 	if err != nil {
// 		return false, err
// 	}

// 	return verified, nil
// }


func (au *authRepo) GetMe(ctx context.Context, email string) (models.User, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "GetMe")
	defer span.End()
	var userData models.User

	if err := au.db.WithContext(ctx).Where("email = ?", email).First(&userData).Error; err != nil {
		fmt.Println("❌ REPO ERROR: User not found:", err)
		return userData, err
    }
    
    return userData, nil
}

func (au *authRepo) VerifyUserLogin(ctx context.Context, email, password string, userData models.User) (bool, error) {
    ctx, span := tracing.CreateSpanWrapper(ctx, "VerifyUserLogin")
    defer span.End()

    verified, err := au.comparePassword(ctx, password, userData.Hash)
    if err != nil {
        fmt.Println("❌ REPO ERROR: Compare password failed:", err)
        return false, err
    }
    
    fmt.Println("✅ REPO DEBUG: Password comparison result:", verified)
    return verified, nil
}