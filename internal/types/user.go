package types

import "time"

type User struct {
	ID        int        `gorm:"primaryKey" json:"id,omitempty"`
	Email     string     `json:"email" validate:"required,min=10,max=40"`
	Password  *string    `gorm:"password" validate:"required,min=10,max=40" json:"password,omitempty"`
	Salt      *string    `gorm:"salt" json:"salt,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
