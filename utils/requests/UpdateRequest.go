package requests

type UpdateRequest[T any] struct {
	TargetId int64 `json:"target_id" validate:"required,number"`
	NewData  T     `json:"new_data" validate:"required"`
}
