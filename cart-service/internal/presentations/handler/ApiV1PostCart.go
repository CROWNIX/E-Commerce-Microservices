package handler

import (
	"cart-service/generated/api"
	"cart-service/internal/services/cart"
	"net/http"

	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ApiV1PostCart(c *gin.Context) {
	req := api.ApiV1PostCartRequestBody{}
	if ok := ginx.MustShouldBind(c, &req); !ok {
		return
	}

	err := h.Service.CartService.CreateCart(c.Request.Context(), cart.CreateCartInput{
		ApiV1PostCartRequestBody: req,
	})

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1PostCartResponseBody{
		Message: "Add product to cart Successfully",
	}

	c.JSON(http.StatusOK, resp)
}
