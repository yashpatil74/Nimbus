package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/handlers"
	"github.com/yashpatil74/nimbus/internal/api/middlewares"
)

func SetupFileRoutes(router *gin.RouterGroup, fileHandler *handlers.FileHandler) {
	fileRoutes := router.Group("/file")
	fileRoutes.Use(middlewares.AuthMiddleware())
	{
		fileRoutes.POST("/upload", fileHandler.UploadFile)
		fileRoutes.GET("/download/:id", fileHandler.DownloadFile)
		fileRoutes.GET("/list", fileHandler.ListFiles)
		fileRoutes.DELETE("/delete/:id", fileHandler.DeleteFile)
	}
}
