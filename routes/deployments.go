package routes

import (
	"github.com/SpaceTesla/orbis-be/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/deployments", controllers.GetDeployments)
		api.POST("/deployments", controllers.CreateDeployment)
	}
}
