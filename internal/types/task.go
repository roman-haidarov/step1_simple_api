package types

import "time"

type Task struct {
	UUID        string    `gorm:"type:uuid;primaryKey" json:"uuid"`
	Description string    `json:"description" validate:"required,min=1"`
	IsDone      bool      `gorm:"default:false" json:"is_done"`
	UserId      int       `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
