package handler

import (
	"product-service/generated/api"
	"product-service/internal/models"
	"product-service/internal/services/category"
	"fmt"
	"net/http"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/CROWNIX/go-utils/utils/generic"
	"github.com/CROWNIX/go-utils/utils/primitive"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

func (h *Handler) ApiV1GetCategories(c *gin.Context) {
	var sortField string
	if paramValue := c.Query("sort_field"); paramValue == "" {
		ginx.ErrorResponse(c, apperror.BadRequest("Query argument sort_field is required, but not found"))
		return
	}

	err := runtime.BindQueryParameter("form", true, true, "sort_field", c.Request.URL.Query(), &sortField)
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

	output, err := h.Service.CategoryService.GetCategories(
		c.Request.Context(),
		category.GetCategoriesInput{
			Sorting: primitive.NewSortingFromQueryParams(sortDirection, sortField),
		},
	)

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1GetCategoriesResponseBody{
		Message: "Get Categories Successfully",
		Items: generic.TransformSlice(output.Items, func(cat models.Category) api.ApiV1GetCategory {
			return api.ApiV1GetCategory{
				Id:       cat.ID,
				Name:     cat.Name,
				Image:    cat.Image,
				ParentId: cat.ParentID,
			}
		}),
	}

	c.JSON(http.StatusOK, resp)
}
