package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
    "github.com/takumifahri/RESTful-API-GO/internal/tracing"
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
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "RegisterUser")
    defer span.End() // Pastikan span diakhiri
    var request models.RegisterRequest

    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return c.JSON(400, echo.Map{
            "message": "Invalid request body",
            "error":   err.Error(),
        })
    }

    userData, err := h.AuthUsecase.RegisterUser(ctx, request)  // ✅ Sekarang bisa akses
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

func (h *AuthHandler) LoginUser(c echo.Context) error {
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "LoginUser")
    defer span.End() // Pastikan span diakhiri
    var request models.LoginRequest

    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "Invalid request body",
            "error":   err.Error(),
        })
    } 

    sessionData, err := h.AuthUsecase.LoginUser(ctx, request) // ✅ Sekarang bisa akses
    if err != nil {
        fmt.Println("Error logging in user:", err)
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "Failed to log in user",
            "error":   err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "User logged in successfully",
        "data":    sessionData,
    })
}