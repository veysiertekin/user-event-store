package model

type GenericResult struct {
	Status       Status `json:"status,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func Success() *GenericResult {
	result := new(GenericResult)
	result.Status = SUCCESS
	return result
}

func Error(errorMessage string) *GenericResult {
	result := new(GenericResult)
	result.Status = ERROR
	result.ErrorMessage = errorMessage
	return result
}
