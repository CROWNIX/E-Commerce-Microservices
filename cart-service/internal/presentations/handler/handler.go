package handler

import (
	// "cart-service/generated/api"
	"cart-service/internal/presentations/handler/cms"
	"cart-service/internal/services"

	// "github.com/CROWNIX/go-utils/utils/primitive"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service   *services.Service
	*cms.CmsHandler
}

type Options struct {
	Service    *services.Service
}

func NewHandler(opts Options) *Handler {
	return &Handler{	
		Service:   opts.Service,
		CmsHandler: cms.NewHandler(cms.Options{
			Service: opts.Service,
		}),
	}
}

// func bindToPaginationResponse(input primitive.PaginationOutput) api.Pagination {
// 	return api.Pagination{
// 		Page:      input.Page,
// 		PageSize:  input.PageSize,
// 		Total:     input.TotalData,
// 		TotalPage: input.PageCount,
// 	}
// }

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.POST("/carts", h.ApiV1PostCart)
	v1.PATCH("/users/:userID/carts/:productID/increment", h.ApiV1PatchIncrementCart)
	v1.PATCH("/users/:userID/carts/:productID/decrement", h.ApiV1PatchDecrementCart)
	v1.DELETE("/carts/:cartID", h.ApiV1DeleteCart)
}
