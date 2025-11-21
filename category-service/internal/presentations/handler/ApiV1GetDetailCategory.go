package handler

import (
	"category-service/generated/api"
	"category-service/internal/services/category"
	"net/http"

	"fmt"

	"github.com/CROWNIX/go-utils/apperror"
	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/CROWNIX/go-utils/utils/generic"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

func (h *Handler) ApiV1GetParentCategory(c *gin.Context) {
	var categoryID uint64

	err := runtime.BindStyledParameterWithOptions("simple", "categoryID", c.Param("categoryID"), &categoryID, runtime.BindStyledParameterOptions{Explode: false, Required: true})

	if err != nil {
		ginx.ErrorResponse(c, apperror.BadRequest(fmt.Sprintf("invalid format for parameter categoryID: %w", err)))

		return
	}

	output, err := h.Service.CategoryService.GetParentCategory(c.Request.Context(), categoryID)

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1GetParentCategoryResponseBody{
		Id:      categoryID,
		Name:    output.Name,
		Message: "Get Category Successfully",
		Children: generic.TransformSlice(output.Children, func(category category.GetCategoryChildren) api.ApiV1GetCategoryChildren {
			return api.ApiV1GetCategoryChildren{
				Id:   category.ID,
				Name: category.Name,
			}
		}),
	}

	c.JSON(http.StatusOK, resp)
}
