package main

import (
	"context"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tiago-bitten/notification-service/docs"

	projectHttp "github.com/tiago-bitten/notification-service/internal/project/api/http"
	projectService "github.com/tiago-bitten/notification-service/internal/project/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017/notificationservice?authSource=admin"))
	if err != nil {
		panic(err)
	}

	db := client.Database("notificationservice")

	projectApp := projectService.NewApplication(db)

	routes := gin.Default()

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	projectHttp.RegisterRoutes(routes, projectApp)

	err = routes.Run(":8080")
	if err != nil {
		panic(err)
	}
}
