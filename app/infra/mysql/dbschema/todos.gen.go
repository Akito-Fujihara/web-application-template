// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbschema

import (
	"time"
)

const TableNameTodo = "todos"

// Todo mapped from table <todos>
type Todo struct {
	ID          int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID      int64      `gorm:"column:user_id;type:bigint;not null;index:user_id,priority:1" json:"user_id"`
	Title       string     `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description *string    `gorm:"column:description;type:text" json:"description"`
	IsCompleted *bool      `gorm:"column:is_completed;type:tinyint(1)" json:"is_completed"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Todo's table name
func (*Todo) TableName() string {
	return TableNameTodo
}
