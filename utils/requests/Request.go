package requests

import "encoding/json"

type Request[T any] struct {
	TargetId json.Number `json:"target_id" validate:"required,number"`
	Data     T           `json:"data" validate:"required"`
}
