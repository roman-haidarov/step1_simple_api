package types

type Task struct {
	UUID        string `gorm:"primaryKey" json:"uuid"`
	Description string `json:"description" validate:"required,min=1"`
}
