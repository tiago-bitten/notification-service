package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/project/domain/project"
)

type ProjectAuthMiddleware struct {
	projectRepo project.Repository
}

func NewProjectAuthMiddleware(projectRepo project.Repository) *ProjectAuthMiddleware {
	return &ProjectAuthMiddleware{projectRepo: projectRepo}
}

func (m *ProjectAuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Project-Key")
		if key == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		p, err := m.projectRepo.FindByKey(key)
		if err != nil || p == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
			return
		}

		c.Set("project", p)
		c.Next()
	}
}
