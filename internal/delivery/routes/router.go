package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest/user"
)

func LoadRoutes(e *echo.Echo, handler *rest.Handler, authHandler *user.AuthHandler) {
	// Grouping
	catalogGroup := e.Group("/catalog")
	orderGroup := e.Group("/order")
	authGroup := e.Group("/auth")
	// Catalog
	catalogGroup.GET("", handler.GetAllCatalogList)
	catalogGroup.GET("/:unique_id", handler.GetCatalogByID)
	catalogGroup.POST("", handler.AddCatalog)
	catalogGroup.PATCH("/:unique_id", handler.UpdateCatalog)

	//Order
	orderGroup.POST("", handler.Order)
	orderGroup.GET("", handler.AdminGetAllOrder)
	orderGroup.GET("/:unique_id", handler.GetOrderInfo)


	// auth 
	authGroup.POST("/register", authHandler.RegisterUser)
	authGroup.POST("/login", authHandler.LoginUser)
}