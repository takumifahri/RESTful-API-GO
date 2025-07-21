package auth

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Usecase interface {
	RegisterUser(request models.RegisterRequest) (models.User, error)
	LoginUser(request models.LoginRequest) (models.UserSession, error)
}