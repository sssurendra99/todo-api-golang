package extensions

import (
	
	"Todo-api-Golang/core/application/features"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRouting(router *gin.Engine) {

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("todo/", features.GetTodos)
	router.POST("todo/", features.AddTodo)
	router.Run(":10000")
	
}