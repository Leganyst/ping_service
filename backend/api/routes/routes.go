package routes

import (
	"vktest/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/containers", handlers.GetContainers)
	r.POST("/containers", handlers.CreateContainer)
	r.DELETE("/containers/:ip", handlers.DeleteContainer)

	return r
}
