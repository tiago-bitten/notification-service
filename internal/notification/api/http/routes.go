package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/notification/app"
)

func RegisterNotificationRoutes(g *gin.RouterGroup, app app.Application) {
	handler := NewNotificationHandler(app)

	g.POST("/notifications", handler.CreateNotification)
}
