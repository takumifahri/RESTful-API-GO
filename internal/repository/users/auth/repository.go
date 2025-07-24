package auth

import (
	"context"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

type Repository interface {
	RegisterUser(ctx context.Context, userData models.User) (models.User, error)
	CheckRegistered(ctx context.Context, Email string) (bool, error)
	GenerateUserHash(ctx context.Context, password string) (hash string, err error)
	VerifyUserLogin(ctx context.Context, Email, Password string, userData models.User) (bool, error)
	GetMe(ctx context.Context, Email string) (models.User, error)
	CreateUserSession(ctx context.Context, userUniqueID string) (models.UserSession, error)
	CheckSession(ctx context.Context, data models.UserSession) (userUniqueID string, err error)
}
