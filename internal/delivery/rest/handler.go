package rest

import "github.com/takumifahri/RESTful-API-GO/internal/usecase/store"

type Handler struct {
	storeUsecase store.Usecase
}

func NewHandler(storeUsecase store.Usecase) *Handler {
	return &Handler{
		storeUsecase: storeUsecase,
	}
}