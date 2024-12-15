package main

import (
	middleware "Todo-api-Golang/api/middlewares"
	"Todo-api-Golang/infrastructure/persistence/db"
	"Todo-api-Golang/infrastructure/utility"
	"log"

	_ "Todo-api-Golang/docs"
	"Todo-api-Golang/core/application/features"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {


	utility.LoadingEnv()

	db, err := db.ConnectGormDB()
	if err != nil {
		log.Println("Connection error.")
		return
	}

	router := gin.Default()

	// Use the DB middleware for all routes
	router.Use(middleware.DBMiddleware(db))

	// Swagger UI route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	router.GET("/todos", features.GetTodos)   // Get all todos
	router.POST("/todos", features.AddTodo)  // Add a new todo

	// Run the server
	router.Run(":10000")
}
