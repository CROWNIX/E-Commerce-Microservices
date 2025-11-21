package handler

import (
	"fmt"
	"net/http"
	"product-service/generated/api"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

func (h *Handler) ApiV1GetDetailProduct(c *gin.Context) {
	var productID uint64

	err := runtime.BindStyledParameterWithOptions("simple", "productID", c.Param("productID"), &productID, runtime.BindStyledParameterOptions{Explode: false, Required: true})

	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter productID: %w", err)))

		return
	}

	output, err := h.Service.ProductService.GetDetailProduct(c.Request.Context(), productID)

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1GetDetailProductResponseBody{
		Id:              productID,
		Name:            output.Name,
		Images:          output.Images.V,
		Description:     output.Description,
		Price:           output.Price,
		Stock:           output.Stock,
		FinalPrice:      output.FinalPrice,
		DiscountPercent: output.DiscountPercent,
		MinimumPurchase: output.MinimumPurchase,
		MaximumPurchase: output.MaximumPurchase,
		Message:         "Get detail product successfully",
	}

	c.JSON(http.StatusOK, resp)
}
