package handler

import (
	"auth-service/generated/api"
	"auth-service/internal/services/auth"
	"net/http"

	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ApiV1PostRegister(c *gin.Context) {
	req := api.ApiV1PostRegisterRequestBody{}
	if ok := ginx.MustShouldBind(c, &req); !ok {
		return
	}

	err := h.Service.AuthService.Register(c.Request.Context(), auth.RegisterInput{
		ApiV1PostRegisterRequestBody: req,
	})

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1PostRegisterResponseBody{
		Message: "Register Successfully",
	}

	c.JSON(http.StatusOK, resp)
}
