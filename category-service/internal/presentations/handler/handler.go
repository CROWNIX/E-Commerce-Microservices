package handler

import (
	"category-service/generated/api"
	"category-service/internal/presentations/handler/cms"
	"category-service/internal/presentations/middleware"
	"category-service/internal/services"

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
	v1.GET("/categories", h.ApiV1GetCategories)
	v1.GET("/categories/:categoryID", h.ApiV1GetParentCategory)
}
