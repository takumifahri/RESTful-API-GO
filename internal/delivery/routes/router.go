package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
)

func LoadRoutes(e *echo.Echo, handler *rest.Handler) {
	e.GET("/clothes", handler.GetAllCatalogList)
	e.GET("/clothes/:unique_id", handler.GetCatalogByID)
	e.POST("/clothes", handler.AddCatalog)
	e.PATCH("/clothes/:unique_id", handler.UpdateCatalog)
	e.POST("/order", handler.Order)
	e.GET("/order/:unique_id", handler.GetOrderInfo)
}