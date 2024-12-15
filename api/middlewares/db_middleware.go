package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {

    return func(c *gin.Context) {
        // Set the database in the context
        c.Set("db", db)
        c.Next()
    }
}
