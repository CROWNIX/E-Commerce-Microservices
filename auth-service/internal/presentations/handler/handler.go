package handler

import (
	"auth-service/generated/api"
	"auth-service/internal/presentations/handler/cms"
	"auth-service/internal/presentations/middleware"
	"auth-service/internal/services"

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
	v1.POST("/register", h.ApiV1PostRegister)
	v1.POST("/login", h.ApiV1PostLogin)
}
