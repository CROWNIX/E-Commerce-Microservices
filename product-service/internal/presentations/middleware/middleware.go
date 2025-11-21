package middleware

import (
	"net/http"
	"product-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	TokenUtil *utils.TokenUtil
}

func NewMiddleware(tokenUtil *utils.TokenUtil) *Middleware {
	return &Middleware{TokenUtil: tokenUtil}
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	auth, err := m.TokenUtil.ParseToken(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.Set("auth", auth)
	c.Next()
}
