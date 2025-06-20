package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/project/app"
)

func RegisterRoutes(router *gin.Engine, app app.Application) {
	handler := NewProject(app)

	router.POST("/projects", handler.CreateProject)
	router.GET("/projects", handler.FindAllProjects)
}
