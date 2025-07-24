package auth

import (
	// "errors"
	// "fmt"

	// "github.com/google/uuid"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/takumifahri/RESTful-API-GO/internal/models"

	// "github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	// "github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
	// "github.com/takumifahri/RESTful-API-GO/internal/repository/order"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/users/auth"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
)

type authUsecase struct {
	userRepo  auth.Repository
}

func GetUsecase( userRepo auth.Repository) Usecase {
	return &authUsecase{
		userRepo:  userRepo,
	}
}

func (au *authUsecase) RegisterUser(ctx context.Context, request models.RegisterRequest) (models.User, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "RegisterUser")
	defer span.End()
	// cek apakah user sudah register atau belum
	RegisteredUser, err := au.userRepo.CheckRegistered(ctx, request.Name)
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
	hashUser, err := au.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		fmt.Println("Error generating user hash:", err)
		return models.User{}, err
	}

	userData, err := au.userRepo.RegisterUser(ctx, models.User{
		UniqueID: "USR-" + uuid.New().String(),
		Email:    request.Email,
		Name:     request.Name,
		Hash:     hashUser,
		Phone:    request.Phone,
		Address:  request.Address,
	})
	if err != nil {
		fmt.Println("Error registering user:", err)
		return models.User{}, err
	}
	fmt.Println("User registered successfully:", userData)
	return userData, nil
}
func (au *authUsecase) LoginUser(ctx context.Context, request models.LoginRequest) (models.UserSession, error) {
    ctx, span := tracing.CreateSpanWrapper(ctx, "LoginUser")
    defer span.End()

    fmt.Println("🔍 DEBUG: Starting login for email:", request.Email)
    
    fmt.Println("🔍 DEBUG: Getting user data...")
    userData, err := au.userRepo.GetMe(ctx, request.Email)
    if err != nil {
        fmt.Println("❌ ERROR getting user data:", err)
        return models.UserSession{}, err
    }
    fmt.Println("✅ DEBUG: User data retrieved:", userData.Name)

    fmt.Println("🔍 DEBUG: Verifying login...")
    verified, err := au.userRepo.VerifyUserLogin(ctx, request.Email, request.Password, userData)
    if err != nil {
        fmt.Println("❌ ERROR verifying user login:", err)
        return models.UserSession{}, err
    }
    fmt.Println("✅ DEBUG: Login verified:", verified)

    if !verified {
        fmt.Println("❌ DEBUG: Invalid credentials")
        return models.UserSession{}, errors.New("invalid email or password")
    }

    fmt.Println("🔍 DEBUG: Creating user session...")
    userSession, err := au.userRepo.CreateUserSession(ctx, userData.UniqueID)
    if err != nil {
        fmt.Println("❌ ERROR creating user session:", err)
        return models.UserSession{}, err
    }
    fmt.Println("✅ DEBUG: Session created successfully")

    return userSession, nil
}


func (au *authUsecase) CheckSession(ctx context.Context, data models.UserSession) (userUniqueID string, err error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "CheckSession")
	defer span.End()
  	userUniqueID, err = au.userRepo.CheckSession(ctx, data)
	// Debug
	if err != nil {
		fmt.Println("❌ ERROR checking session:", err)
		return "", err
	}
	return userUniqueID, nil
}