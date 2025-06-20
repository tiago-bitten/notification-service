package command

import (
	"github.com/tiago-bitten/notification-service/internal/notification/domain/notification"
	"github.com/tiago-bitten/notification-service/internal/project/domain/project"
)

type CreateNotificationCommand struct {
	ProjectID string
	Title     string
	Body      string
	Type      string
	Users     []string
}

type CreateNotificationHandler struct {
	notificationRepo notification.Repository
	projectRepo      project.Repository
}

func (h *CreateNotificationHandler) Handle(cmd CreateNotificationCommand) error {
	return nil
}
