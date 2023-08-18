package main

import (
	"context"
	"fmt"
	"log"
	"rest_api_app/config"
	"rest_api_app/constants"
	"rest_api_app/controllers"
	"rest_api_app/routes"
	"rest_api_app/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

// func initApp(mongoClient *mongo.Client){
// 	ctx = context.TODO()
// 	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("transactions")
// 	profileService := services.NewProfileServiceInit(profileCollection, ctx)
// 	profileController := controllers.InitProfileController(profileService)
// 	routes.ProfileRoute(server,profileController)
// }

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	transactCollection := mongoClient.Database(constants.DatabaseName).Collection("transactions")
	transactService := services.NewTransServiceInit(transactCollection, ctx)
	transactController := controllers.InitTransactController(transactService)
	routes.TransactRoute(server, transactController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
