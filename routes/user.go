package routes

import (
	"go-gin/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controller.RegisterUser)
	}
}
