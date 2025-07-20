package middlewares


import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func CorsMiddleware(c *echo.Echo) {
	c.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// KTIA BUAT CORS UNTUK ALLOW ORIGIN
		AllowOrigins: []string{"*"},
	}))
}