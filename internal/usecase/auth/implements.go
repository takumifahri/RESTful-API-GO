package auth

import (
	// "errors"
	// "fmt"

	// "github.com/google/uuid"
	"fmt"

	"github.com/google/uuid"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	// "github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	// "github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
	// "github.com/takumifahri/RESTful-API-GO/internal/repository/order"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/users/auth"
)

type authUsecase struct {
	userRepo  auth.Repository
}

func GetUsecase( userRepo auth.Repository) Usecase {
	return &authUsecase{
		userRepo:  userRepo,
	}
}

func (au *authUsecase) RegisterUser(request models.RegisterRequest) (models.User, error) {
	// cek apakah user sudah register atau belum
	RegisteredUser, err := au.userRepo.CheckRegistered(request.Name)
	//debug jika ada error
	if err != nil {
		fmt.Println("Error checking registered user:", err)
		return models.User{}, err
	}

	if RegisteredUser {
		return models.User{}, fmt.Errorf("user with name %s already registered", request.Name)
	}

	// Jika belum terdaftar, buat user baru
	// kita buat hash nya dlu untuk password
	hashUser, err := au.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		fmt.Println("Error generating user hash:", err)
		return models.User{}, err
	}

	userData, err := au.userRepo.RegisterUser(models.User{
		UniqueID: "USR-" + uuid.New().String(),
		Email:    request.Email,
		Name:     request.Name,
		Hash:     hashUser,
		Phone:    request.Phone,
		Address: request.Address,
	})
	if err != nil {
		fmt.Println("Error registering user:", err)
		return models.User{}, err
	}
	fmt.Println("User registered successfully:", userData)
	return userData, nil
}
