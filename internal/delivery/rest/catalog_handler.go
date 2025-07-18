package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/utils"
	// "github.com/takumifahri/RESTful-API-GO/internal/models"
)

func (h *Handler) GetAllCatalogList(c echo.Context) error {

	// Params 
	clothesType := c.QueryParam("TypeClothes")
	catalogData, err := h.storeUsecase.GetAllCatalogList(clothesType)
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

func (h *Handler) GetCatalogByID(c echo.Context) error {
    uniqueID := c.Param("unique_id")
    if uniqueID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Catalog ID is required",
        })
    }

    catalogData, err := h.storeUsecase.GetCatalogByID(uniqueID)

    if err != nil {
        fmt.Printf("Error fetching catalog by unique ID: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Failed to fetch catalog by unique ID",
            "error":   err.Error(),
        })
    }

    if catalogData == nil {
        return c.JSON(http.StatusNotFound, map[string]string{
            "message": "Catalog not found",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Catalog fetched successfully",
        "data":    catalogData,
    })
}

func (h *Handler) AddCatalog(c echo.Context) error {
	var request models.ProductClothes
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("Error decoding request body: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to decode request body",
			"error":   err.Error(),
		})
	}
	//Validasi
    if err := utils.ValidatorStruct(request); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Validation failed",
            "error":   err.Error(),
        })
    }
	catalogData, err := h.storeUsecase.AddCatalog(request)
	if err != nil {
		fmt.Printf("Error adding catalog: %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to add catalog",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Catalog added successfully",
		"data":    catalogData,
	})
	
}