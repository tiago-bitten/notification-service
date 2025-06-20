package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FixedTokenAuthMiddleware struct {
	expectedToken string
}

func NewFixedTokenAuthMiddleware(expectedToken string) *FixedTokenAuthMiddleware {
	return &FixedTokenAuthMiddleware{expectedToken: expectedToken}
}

func (m *FixedTokenAuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != m.expectedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		c.Next()
	}
}
