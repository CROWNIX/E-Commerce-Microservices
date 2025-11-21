package handler

import (
	"fmt"
	"net/http"
	"product-service/generated/api"
	"product-service/internal/services/product"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/CROWNIX/go-utils/utils/generic"
	"github.com/CROWNIX/go-utils/utils/primitive"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

func (h *Handler) ApiV1GetProducts(c *gin.Context) {
	var page int64
	if paramValue := c.Query("page"); paramValue == "" {
		ginx.ErrorResponse(c, apperror.BadRequest("Query argument page is required, but not found"))
		return
	}

	err := runtime.BindQueryParameter("form", true, true, "page", c.Request.URL.Query(), &page)
	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter page: %w", err)))
		return
	}

	var pageSize int64
	if paramValue := c.Query("page_size"); paramValue == "" {
		ginx.ErrorResponse(c, apperror.BadRequest("Query argument page_size is required, but not found"))
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page_size", c.Request.URL.Query(), &pageSize)
	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter page_size: %w", err)))
		return
	}

	var sortField string
	if paramValue := c.Query("sort_field"); paramValue == "" {
		ginx.ErrorResponse(c, apperror.BadRequest("Query argument sort_field is required, but not found"))
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "sort_field", c.Request.URL.Query(), &sortField)
	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter sort_field: %w", err)))
		return
	}

	var sortDirection string
	if paramValue := c.Query("sort_field"); paramValue == "" {
		ginx.ErrorResponse(c, apperror.BadRequest("Query argument sort_field is required, but not found"))
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "sort_direction", c.Request.URL.Query(), &sortDirection)
	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter sort_direction: %w", err)))
		return
	}

	var categoryID *uint64
	err = runtime.BindQueryParameter("form", true, false, "category_id", c.Request.URL.Query(), &categoryID)
	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter category_id: %w", err)))
		return
	}

	output, err := h.Service.ProductService.GetProducts(c.Request.Context(), product.GetProductsInput{
		Pagination: primitive.PaginationInput{
			Page:     page,
			PageSize: pageSize,
		},
		Sorting:    primitive.NewSortingFromQueryParams(sortDirection, sortField),
		CategoryID: categoryID,
	})

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1GetProductsResponseBody{
		Message: "Get Products Successfully",
		Items: generic.TransformSlice(output.Items, func(product product.GetProduct) api.ApiV1GetProduct {
			return api.ApiV1GetProduct{
				Id:         product.ID,
				Name:       product.Name,
				Images:     product.Images.V,
				Price:      product.Price,
				FinalPrice: product.FinalPrice,
			}
		}),
		Metadata: api.MetaDataOnlyPagination{
			Pagination: bindToPaginationResponse(output.PaginationOutput),
		},
	}

	c.JSON(http.StatusOK, resp)
}
