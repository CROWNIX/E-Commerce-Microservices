package handler

import (
	"auth-service/generated/api"
	"auth-service/internal/services/auth"
	"net/http"

	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ApiV1PostLogin(c *gin.Context) {
	req := api.ApiV1PostLoginRequestBody{}
	if ok := ginx.MustShouldBind(c, &req); !ok {
		return
	}

	accessToken, err := h.Service.AuthService.Login(c.Request.Context(), auth.LoginInput{
		ApiV1PostLoginRequestBody: req,
	})

	if err != nil {
		ginx.ErrorResponse(c, err)
		return
	}

	resp := api.ApiV1PostLoginResponseBody{
		Message:     "Login Successfully",
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, resp)
}
