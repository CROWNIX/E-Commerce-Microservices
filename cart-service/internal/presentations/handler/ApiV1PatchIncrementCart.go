package handler

import (
	"cart-service/generated/api"
	"fmt"
	"net/http"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

func (h *Handler) ApiV1PatchIncrementCart(c *gin.Context) {
	var userID uint64

	err := runtime.BindStyledParameterWithOptions("simple", "userID", c.Param("userID"), &userID, runtime.BindStyledParameterOptions{Explode: false, Required: true})

	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter userID: %w", err)))

		return
	}

	var productID uint64

	err = runtime.BindStyledParameterWithOptions("simple", "productID", c.Param("productID"), &productID, runtime.BindStyledParameterOptions{Explode: false, Required: true})

	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter productID: %w", err)))

		return
	}

	err = h.Service.CartService.IncrementCart(c.Request.Context(), userID, productID)
	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1PatchIncrementCartResponseBody{
		Message: "Increment cart Successfully",
	}

	c.JSON(http.StatusOK, resp)
}
