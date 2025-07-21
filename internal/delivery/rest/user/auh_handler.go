package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

type AuthHandler struct {
	*rest.Handler
}
// ✅ Tambahkan constructor
func NewAuthHandler(handler *rest.Handler) *AuthHandler {
    return &AuthHandler{
        Handler: handler,
    }
}

func (h *AuthHandler) RegisterUser(c echo.Context) error {
    var request models.RegisterRequest

    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return c.JSON(400, echo.Map{
            "message": "Invalid request body",
            "error":   err.Error(),
        })
    }

    userData, err := h.AuthUsecase.RegisterUser(request)  // ✅ Sekarang bisa akses
    if err != nil {
        fmt.Println("Error registering user:", err)
        return c.JSON(500, echo.Map{
            "message": "Failed to register user",
            "error":   err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "User registered successfully",
        "data":    userData,
    })
}