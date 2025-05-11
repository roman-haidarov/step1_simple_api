package types

type Body struct {
	Task string `json:"task" validate:"required,min=1"`
}
