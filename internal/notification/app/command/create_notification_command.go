package command

import (
	"github.com/tiago-bitten/notification-service/internal/notification/domain/notification"
)

type userCommand struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Channel string `json:"channel"`
}

type CreateNotificationCommand struct {
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Type  notification.Type `json:"type"`
	Users []userCommand     `json:"users"`
}

type CreateNotificationHandler struct {
	notificationRepo notification.Repository
}

func NewCreateNotificationHandler(notificationRepo notification.Repository) *CreateNotificationHandler {
	return &CreateNotificationHandler{
		notificationRepo: notificationRepo,
	}
}

func (h *CreateNotificationHandler) Handle(cmd CreateNotificationCommand, projectKey string) error {
	var users []notification.User
	for _, u := range cmd.Users {
		users = append(users, *notification.NewUser(u.ID, u.Name, u.Channel))
	}

	n := notification.NewNotification(cmd.Title, cmd.Body, cmd.Type, projectKey, users)
	return h.notificationRepo.Save(n)

	// create a strategy to notify (email_notify_strategy.go, sms_notify_strategy.go and push_notify_strategy.go)
}
