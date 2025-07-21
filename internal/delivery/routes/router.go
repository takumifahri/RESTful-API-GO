package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest/user"
)

func LoadRoutes(e *echo.Echo, handler *rest.Handler, authHandler *user.AuthHandler) {
	// Catalog
	e.GET("/clothes", handler.GetAllCatalogList)
	e.GET("/clothes/:unique_id", handler.GetCatalogByID)
	e.POST("/clothes", handler.AddCatalog)
	e.PATCH("/clothes/:unique_id", handler.UpdateCatalog)

	//Order
	e.POST("/order", handler.Order)
	e.GET("/orders", handler.AdminGetAllOrder)
	e.GET("/order/:unique_id", handler.GetOrderInfo)


	// auth 
	e.POST("/auth/register", authHandler.RegisterUser)
}