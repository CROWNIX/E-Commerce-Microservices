package handler

import (
	"product-service/generated/api"
	"product-service/internal/presentations/handler/cms"
	"product-service/internal/presentations/middleware"
	"product-service/internal/services"

	"github.com/CROWNIX/go-utils/utils/primitive"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service   *services.Service
	Middlware *middleware.Middleware
	*cms.CmsHandler
}

type Options struct {
	Service    *services.Service
	Middleware *middleware.Middleware
}

func NewHandler(opts Options) *Handler {
	return &Handler{
		Service:   opts.Service,
		Middlware: opts.Middleware,
		CmsHandler: cms.NewHandler(cms.Options{
			Service: opts.Service,
		}),
	}
}

func bindToPaginationResponse(input primitive.PaginationOutput) api.Pagination {
	return api.Pagination{
		Page:      input.Page,
		PageSize:  input.PageSize,
		Total:     input.TotalData,
		TotalPage: input.PageCount,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/products", h.ApiV1GetProducts)
	v1.GET("/products/:productID", h.ApiV1GetDetailProduct)
}
