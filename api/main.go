package main

import (
	"Todo-api-Golang/api/docs"
	middleware "Todo-api-Golang/api/middlewares"
	"Todo-api-Golang/infrastructure/persistence/db"
	"Todo-api-Golang/infrastructure/utility"
	"log"

	_ "Todo-api-Golang/api/docs"
	"Todo-api-Golang/core/application/features"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version 1.0
// @description This is a simple API developed with Golang
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:10000
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "This is a simple API developed with Golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:10000"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
	router.GET("/todos/", features.GetTodos)   // Get all todos
	router.POST("/todos/", features.AddTodo)  // Add a new todo

	// Run the server
	router.Run(":10000")
}
