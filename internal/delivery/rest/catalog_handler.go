package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/utils"
    "github.com/takumifahri/RESTful-API-GO/internal/tracing"
	// "github.com/takumifahri/RESTful-API-GO/internal/models"
)

func (h *Handler) GetAllCatalogList(c echo.Context) error {
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "GetAllCatalogList")
    defer span.End()
	// Params 
	clothesType := c.QueryParam("TypeClothes")
	catalogData, err := h.storeUsecase.GetAllCatalogList(ctx, clothesType)
	if err != nil {
		fmt.Printf("Error fetching catalog %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]string{
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
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "GetCatalogByID")
    defer span.End()
    uniqueID := c.Param("unique_id")
    if uniqueID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Catalog ID is required",
        })
    }

    catalogData, err := h.storeUsecase.GetCatalogByID(ctx, uniqueID)

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
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "AddCatalog")
    defer span.End()
    
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
    if err := utils.ValidateStruct(request); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Validation failed",
            "error":   err.Error(),
        })
    }
	catalogData, err := h.storeUsecase.AddCatalog(ctx, request)
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
func (h *Handler) UpdateCatalog(c echo.Context) error {
    ctx, span := tracing.CreateSpanWrapper(c.Request().Context(), "UpdateCatalog")
    defer span.End()

    uniqueID := c.Param("unique_id")
    if uniqueID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Catalog ID is required",
        })
    }

    // Buat struct untuk partial update dengan pointer
    var request struct {
        NamaPakaian  *string `json:"nama_pakaian,omitempty"`
        Price        *int64    `json:"price,omitempty"`
        Deskripsi    *string `json:"deskripsi,omitempty"`
        Stock        *int                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              `json:"stock,omitempty"`
        TypeClothes  *models.Tipe `json:"type_clothes,omitempty"`
    }

    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err != nil {
        fmt.Printf("Error decoding request body: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Failed to decode request body",
            "error":   err.Error(),
        })
    }

    // Get existing catalog data
    existingCatalog, err := h.storeUsecase.GetCatalogByID(ctx, uniqueID)
    if err != nil {
        fmt.Printf("Error fetching existing catalog: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Failed to fetch existing catalog",
            "error":   err.Error(),
        })
    }

    if existingCatalog == nil {
        return c.JSON(http.StatusNotFound, map[string]string{
            "message": "Catalog not found",
        })
    }

    // Update only fields that are provided
    updateRequest := *existingCatalog
    updateRequest.UNIQUEID = uniqueID

    if request.NamaPakaian != nil {
        updateRequest.NamaPakaian = *request.NamaPakaian
    }
    if request.Price != nil {
        updateRequest.Price = *request.Price
    }
    if request.Deskripsi != nil {
        updateRequest.Deskripsi = *request.Deskripsi
    }
    if request.Stock != nil {
        updateRequest.Stock = *request.Stock
    }
    if request.TypeClothes != nil {
        updateRequest.TypeClothes = models.Tipe(*request.TypeClothes)
    }

    catalogData, err := h.storeUsecase.UpdateCatalog(ctx, updateRequest)
    if err != nil {
        fmt.Printf("Error updating catalog: %s\n", err.Error())
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Failed to update catalog",
            "error":   err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Catalog updated successfully",
        "data":    catalogData,
    })
}