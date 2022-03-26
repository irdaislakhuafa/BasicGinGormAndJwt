package utils

type ResponseMessage struct {
	StatusCode int                    `json:"status_code"`
	Error      interface{}            `json:"error"`
	Data       interface{}            `json:"data"`
	Fields     map[string]interface{} `json:"fields"`
}
