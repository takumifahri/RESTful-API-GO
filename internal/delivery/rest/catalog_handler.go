package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/takumifahri/RESTful-API-GO/internal/models"
)

func (h *handler) GetAllCatalog(c echo.Context) error {

	// Params 
	clothesType := c.QueryParam("TypeClothes")
	catalogData, err := h.storeUsecase.GetAllCatalog(clothesType)
	if err != nil {
		fmt.Printf("Error fetching catalog %s\n", err.Error())

		return  c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch catalog",
			"error":   err.Error(),	
		})
	}

	if len(catalogData) == 0 {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "No catalog found for the specified type",
			"data":    []any{},
		})
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Catalog fetched successfully",
		"data":    catalogData,
	})

}
