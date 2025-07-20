package middlewares

import (
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func LoadMiddleware(c *echo.Echo) {
	// c.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	// KTIA BUAT CORS UNTUK ALLOW ORIGIN
	// 	AllowOrigins: []string{"*"},
	// }))
}

func RoleChecker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Implementasi logika role checker di sini
		// Contoh: cek apakah user memiliki role yang sesuai
		return next(c)
	}
}