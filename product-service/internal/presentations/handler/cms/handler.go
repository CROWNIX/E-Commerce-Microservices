package cms

import (
	"product-service/internal/services"
)

type CmsHandler struct {
	Service *services.Service
}

type Options struct {
	Service *services.Service
}

func NewHandler(opts Options) *CmsHandler {
	return &CmsHandler{
		Service: opts.Service,
	}
}
