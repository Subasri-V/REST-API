package routes

import (
	"rest_api_app/controllers"

	"github.com/gin-gonic/gin"
)

func TransactRoute(router *gin.Engine, controller controllers.TransactController) {
	router.POST("/api/transact/create", controller.CreateTransaction)
	router.POST("/api/profile/get", controller.GetAllTransaction)
}
