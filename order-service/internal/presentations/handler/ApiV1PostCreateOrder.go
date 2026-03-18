package handler

import (
	"fmt"
	"net/http"
	"order-service/generated/api"
	"order-service/internal/services/order"

	utilGeneric "github.com/CROWNIX/go-utils/utils/generic"

	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ApiV1PostCreateOrder(c *gin.Context) {
	req := api.CreateOrderServiceInput{}
	if ok := ginx.MustShouldBind(c, &req); !ok {
		return
	}

	_, err := h.Service.OrderService.CreateOrder(c.Request.Context(), order.CreateOrderServiceInput{
		UserID: req.UserId,
		AddressID: req.AddressId,
		PaymentMethodID: req.PaymentMethodId,
		GrandTotal: req.GrandTotal,
		Items:   utilGeneric.TransformSlice(req.Items, func(item api.Item) order.Item {
			return order.Item{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			}
		}),
	})

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiPostOrderResponseBody{
		Message: "Create order Successfully",
	}

	c.JSON(http.StatusOK, resp)
}
