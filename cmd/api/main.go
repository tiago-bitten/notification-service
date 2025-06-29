package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tiago-bitten/notification-service/docs"
	notificationHttp "github.com/tiago-bitten/notification-service/internal/notification/api/http"
	notificationService "github.com/tiago-bitten/notification-service/internal/notification/service"
	"github.com/tiago-bitten/notification-service/internal/project/infra/repository"
	"github.com/tiago-bitten/notification-service/internal/shared/config"
	"github.com/tiago-bitten/notification-service/internal/shared/http/middleware"

	projectHttp "github.com/tiago-bitten/notification-service/internal/project/api/http"
	projectService "github.com/tiago-bitten/notification-service/internal/project/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	appConfig := config.LoadConfig()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017/notificationservice?authSource=admin"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	db := client.Database(appConfig.Database)

	notificationApp := notificationService.NewApplication(db)
	projectApp := projectService.NewApplication(db)
	projectRepo := repository.NewMongoProjectRepository(db)

	projectAuthMiddleware := middleware.NewProjectAuthMiddleware(projectRepo)
	// fixedTokenAuthMiddleware := middleware.NewFixedTokenAuthMiddleware(appConfig.FixedAuthToken)

	r := gin.Default()
	// public := r.Group("/api/v1")

	projectProtectedGroup := r.Group("/api/v1")
	projectProtectedGroup.Use(projectAuthMiddleware.Handle())
	notificationHttp.RegisterNotificationRoutes(projectProtectedGroup, notificationApp)

	adminGroup := r.Group("/api/v1")
	// adminGroup.Use(fixedTokenAuthMiddleware.Handle())
	projectHttp.RegisterProjectRoutes(adminGroup, projectApp)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
