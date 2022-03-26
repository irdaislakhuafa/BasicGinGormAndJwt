package requests

import "encoding/json"

type DeleteRequest struct {
	TargetId json.Number `json:"target_id" validate:"required,number"`
}
