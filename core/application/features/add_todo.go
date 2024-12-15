package features

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Todo-api-Golang/core/domain/entites"
)

// AddTodo godoc
// @Summary Create a new Todo
// @Description Add a new todo to the database
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body entities.Todo true "Todo Model"
// @Success 201 {object} entities.Todo
// @Router /todos [post]
func AddTodo(c *gin.Context) {
	db, _ := c.Get("db")
    database := db.(*gorm.DB)

    var todo entites.Todo
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