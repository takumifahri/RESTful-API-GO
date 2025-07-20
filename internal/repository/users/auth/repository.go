package auth

import "github.com/takumifahri/RESTful-API-GO/internal/models"

type Repository interface {
	RegisterUser(userData models.User) (models.User, error)
	CheckRegistered(Name string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
}
