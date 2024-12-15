package entites

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


// Todo represents a task or item in the system
// @Description A task or to-do item with completion status
type Todo struct {
	// @description Unique identifier for the Todo
	// @example "3fcdfe19-0dc6-4b9a-9354-c8b7e0b8a7e5"
	ID string `gorm:"primaryKey" json:"id"`

	// @description Name or title of the Todo
	// @example "Buy groceries"
	Name string `json:"name"`

	// @description Whether the Todo is completed
	// @example false
	IsCompleted bool `gorm:"default:false" json:"isCompleted"`

	// @description Timestamp when the Todo was created
	// @example "2024-12-15T12:34:56.789Z"
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`

	// @description Timestamp when the Todo was last updated
	// @example "2024-12-15T13:45:56.789Z"
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
func (Todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	Todo.ID = uuid.NewString()
	return
}

