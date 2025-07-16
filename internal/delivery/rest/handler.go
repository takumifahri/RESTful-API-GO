package rest

import "github.com/takumifahri/RESTful-API-GO/internal/usecase/store"

type handler struct {
	storeUsecase store.Usecase
}

func NewHandler(storeUsecase store.Usecase) *handler {
	return &handler{
		storeUsecase: storeUsecase,
	}
}