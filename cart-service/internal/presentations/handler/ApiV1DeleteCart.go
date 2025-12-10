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

func (h *Handler) ApiV1DeleteCart(c *gin.Context) {
	var cartID uint64

	err := runtime.BindStyledParameterWithOptions("simple", "cartID", c.Param("cartID"), &cartID, runtime.BindStyledParameterOptions{Explode: false, Required: true})

	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter cartID: %w", err)))

		return
	}

	err = h.Service.CartService.DeleteCart(c.Request.Context(), cartID)
	if err != nil{
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1DeleteCartResponseBody{
		Message: "Delete product in cart Successfully",
	}

	c.JSON(http.StatusOK, resp)
}
