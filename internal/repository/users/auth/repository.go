package auth

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	RegisterUser(userData models.User) (models.User, error)
	CheckRegistered(Email string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	VerifyUserLogin(Email, Password string, userData models.User) (bool, error)
	GetMe(Email string) (models.User, error)
	CreateUserSession(userUniqueID string) (models.UserSession, error) 
}
