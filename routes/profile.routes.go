package routes

import (
	"rest_api_app/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine, controller controllers.ProfileController) {
	router.POST("/api/profile/create", controller.CreateProfile)
	router.POST("/api/profile/edit", controller.EditProfile)
	router.DELETE("/api/profile/delete/:id", controller.DeleteProfile)
	router.GET("/api/profile/:id", controller.GetProfileById)
	router.POST("/api/profile/search", controller.GetProfileBySearch)
}
