package types

type Task struct {
	UUID        string `json:"uuid"`
	Description string `json:"description" validate:"required,min=1"`
}
