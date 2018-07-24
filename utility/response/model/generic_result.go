package model

type GenericResult struct {
	Status       Status      `json:"status,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func Success() *GenericResult {
	result := new(GenericResult)
	result.Status = SUCCESS
	return result
}

func Error(err error) *GenericResult {
	result := new(GenericResult)
	result.Status = ERROR
	result.ErrorMessage = err.Error()
	return result
}
