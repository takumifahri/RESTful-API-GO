package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

func (h *Handler) Order(c echo.Context) error {
	var request models.OrderMenuRequest
	// Kita ambil datanya menjadi Json
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("Error decoding request body: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to decode request body",
			"error":   err.Error(),
		})
		
	}

	orderData, err := h.storeUsecase.Order(request)
	if err != nil {
		fmt.Printf("Error processing order: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to process order",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order processed successfully",
		"data":    orderData,
	})
}

func (h *Handler) GetOrderInfo(c echo.Context) error {
	orderID := c.Param("unique_id")

	orderData, err := h.storeUsecase.GetOrderInfo(models.GetOrderInfoRequest{
		OrderID: orderID,
	})

	if err != nil {
		fmt.Printf("Error fetching order info: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch order info",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order info fetched successfully",
		"data":    orderData,
	})

}
func (h *Handler) AdminGetAllOrder(c echo.Context) error {
    // Tidak perlu decode request body karena tidak ada parameter

    orderData, err := h.storeUsecase.AdminGetAllOrder()
    if err != nil {
        fmt.Printf("Error fetching all orders: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Failed to fetch all orders",
            "error":   err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "All orders fetched successfully",
        "data":    orderData,
    })
}