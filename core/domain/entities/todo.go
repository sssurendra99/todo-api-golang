package entities

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

	// @description The timestamp when the record was created
	CreatedAt time.Time `json:"createdAt"`

	// @description The timestamp when the record was last updated
	UpdatedAt time.Time `json:"updatedAt"`

	// @description The timestamp when the record was deleted (for soft deletes)
	DeletedAt time.Time `json:"deletedAt"`

	// @description Name or title of the Todo
	// @example "Buy groceries"
	Name string `json:"name"`

	// @description Whether the Todo is completed
	// @example false
	IsCompleted bool `gorm:"default:false" json:"isCompleted"`
}

func (Todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	Todo.ID = uuid.NewString()
	return
}
