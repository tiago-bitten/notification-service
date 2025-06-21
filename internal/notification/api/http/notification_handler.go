package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiago-bitten/notification-service/internal/notification/app"
	"github.com/tiago-bitten/notification-service/internal/notification/app/command"
	"net/http"
)

type NotificationHandler struct {
	app app.Application
}

func NewNotificationHandler(app app.Application) *NotificationHandler {
	return &NotificationHandler{
		app: app,
	}
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var cmd command.CreateNotificationCommand
	if err := c.ShouldBindBodyWithJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectKey := c.GetHeader("X-Project-Key")
	err := h.app.Commands.CreateNotification.Handle(cmd, projectKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, gin.H{"success": true})
}
