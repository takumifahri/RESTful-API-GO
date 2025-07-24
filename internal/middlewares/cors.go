package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)


func CorsMiddleware(c *echo.Echo) {
	c.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// KTIA BUAT CORS UNTUK ALLOW ORIGIN
		AllowOrigins: []string{"*"},
	}))

	c.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}