package features

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	entities "Todo-api-Golang/core/domain/entities"
)

// AddTodo godoc
//	@Summary		Create a new Todo
//	@Description	Add a new todo to the database
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		entities.Todo	true	"Todo Model"
//	@Success		201		{object}	entities.Todo
//	@Failure		400		{object}	shared.ResponseError
//	@Failure		500		{object}	shared.ResponseError
//	@Router			/todos [post]
func AddTodo(c *gin.Context) {
	db, _ := c.Get("db")
    database := db.(*gorm.DB)

    var todo entities.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"todo": todo})
}