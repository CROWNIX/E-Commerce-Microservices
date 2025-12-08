package middleware

import (
	"fmt"
	"net/http"

	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/pkg/utils/jwtutils"
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/learn-go-microservices/pkg/httperror"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JwtUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := m.getTokenFromHeader(ctx)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		_, err = m.jwtUtil.ParseAndVerifyWithRedis(ctx, accessToken)
		if err != nil {
			ctx.Error(httperror.NewUnauthorizedError())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func(m *AuthMiddleware) JwtMiddlewareUsingRedis(c *gin.Context){
	token, err := m.getTokenFromHeader(c)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return 
	}
	fmt.Println("token exists in header")
	
	_, err = m.jwtUtil.ParseAndVerifyWithRedis(c.Request.Context(), token)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.Next()
}

func (m *AuthMiddleware) getTokenFromHeader(ctx *gin.Context) (string, error) {
	accessToken := ctx.GetHeader("Authorization")
	fmt.Println(accessToken)
	if accessToken == "" {
		return "", httperror.NewUnauthorizedError()
	}

	return accessToken, nil
}
