package rest

import (
	"github.com/takumifahri/RESTful-API-GO/internal/usecase/auth"
	"github.com/takumifahri/RESTful-API-GO/internal/usecase/store"
)

type Handler struct {
	storeUsecase store.Usecase
	AuthUsecase  auth.Usecase
}

func NewHandler(storeUsecase store.Usecase, AuthUsecase auth.Usecase) *Handler {
	return &Handler{
		storeUsecase: storeUsecase,
		AuthUsecase:  AuthUsecase,
	}
}
