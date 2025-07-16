package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/clothes", handler.GetAllCatalog)
}