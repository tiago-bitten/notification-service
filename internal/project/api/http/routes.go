package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/project/app"
)

func RegisterProjectRoutes(g *gin.RouterGroup, app app.Application) {
	handler := NewProject(app)

	g.POST("/projects", handler.CreateProject)
	g.GET("/projects", handler.FindAllProjects)
}
