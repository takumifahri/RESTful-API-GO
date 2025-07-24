package auth

import (
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"context"
)

type Usecase interface {
	RegisterUser(ctx context.Context, request models.RegisterRequest) (models.User, error)
	LoginUser(ctx context.Context, request models.LoginRequest) (models.UserSession, error)
	CheckSession(ctx context.Context, data models.UserSession) (userUniqueID string, err error)
}