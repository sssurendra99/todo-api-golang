package features

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	entities "Todo-api-Golang/core/domain/entities"
)

// GetTodos godoc
//	@Summary		Get all todos
//	@Description	Retrieve all todo items from the database
//	@Tags			Todos
//	@Produce		json
//	@Success		200	{array}		entities.Todo	"List of Todo items"
//	@Failure		500	{object}	shared.ResponseError
//	@Router			/todos [get]
func GetTodos(c *gin.Context) {
	// Retrieve the database instance
    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database not found"})
        return
    }

    // Type assert the database
    database := db.(*gorm.DB)

    var todos []entities.Todo
    if err := database.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todos)
}