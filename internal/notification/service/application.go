package service

import (
	"github.com/tiago-bitten/notification-service/internal/notification/app"
	"github.com/tiago-bitten/notification-service/internal/notification/app/command"
	"github.com/tiago-bitten/notification-service/internal/notification/infra/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewApplication(db *mongo.Database) app.Application {
	notificationRepo := repository.NewMongoNotificationRepository(db)

	return app.Application{
		Commands: app.Commands{
			CreateNotification: command.NewCreateNotificationHandler(notificationRepo),
		},
	}
}
