package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
)

func (h *Handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "Order")
	defer span.End() 

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
	value := c.Request().Context().Value(constant.AuthcontextKey)
	fmt.Printf("DEBUG: context value = %#v, type = %T\n", value, value)
	userUniqueID, ok := value.(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized: key is of invalid type",
		})
	}
	request.UserUniqueID = userUniqueID
	orderData, err := h.storeUsecase.Order(ctx, request)
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
	ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "GetOrderInfo")
	defer span.End() // Pastikan span diakhiri
	orderID := c.Param("unique_id")
	value := c.Request().Context().Value(constant.AuthcontextKey)
	fmt.Printf("DEBUG: context value = %#v, type = %T\n", value, value)
	userUniqueID, ok := value.(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized: key is of invalid type",
		})
	}

	orderData, err := h.storeUsecase.GetOrderInfo(ctx, models.GetOrderInfoRequest{
		UserUniqueID: userUniqueID,
		OrderID:     orderID,
	})

	if err != nil {
		// fmt.Printf("Error fetching order info: %s\n", err.Error())
		// change to logrus
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[GetOrderInfo] Error fetching order info")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch order info",
			"error":   err.Error(),
		}) // Jangan gunaka loggingg untuk hal data sensitif. contoh password.
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order info fetched successfully",
		"data":    orderData,
	})

}
func (h *Handler) AdminGetAllOrder(c echo.Context) error {
	ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "AdminGetAllOrder")
	defer span.End() // Pastikan span diakhiri
    // Tidak perlu decode request body karena tidak ada parameter

    orderData, err := h.storeUsecase.AdminGetAllOrder(ctx)
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