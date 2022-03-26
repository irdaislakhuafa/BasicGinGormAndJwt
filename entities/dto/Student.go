package dto

type Student struct {
	Nim  string `json:"nim" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}
