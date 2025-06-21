package app

import (
	"github.com/tiago-bitten/notification-service/internal/notification/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateNotification *command.CreateNotificationHandler
}

type Queries struct {
}
