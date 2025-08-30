package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-dilve/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/download-json", controllers.Download_file)
	router.GET("/", controllers.Success)
	router.GET("save-products", controllers.SaveProducts)
	router.GET("download", controllers.DownloadDirectory)
}
