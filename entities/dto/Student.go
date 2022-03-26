package dto

type Student struct {
	Nim  string `json:"nim" validate:"required,number,min=10"`
	Name string `json:"name" validate:"required"`
}
