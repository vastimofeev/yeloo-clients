package routes

import (
	"yeloo-clients/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, profileController *controllers.ProfileController) {
	router.GET("/profile/:id", profileController.GetProfile)
	router.POST("/profile", profileController.CreateProfile)
	router.DELETE("/profile/:id", profileController.DeleteProfile)
}
