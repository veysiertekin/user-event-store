package request

import (
	"net/http"
	"user-event-store/utility/response"
	"user-event-store/utility/response/model"
	"user-event-store/utility/validation"
)

func Handle(r *http.Request, w http.ResponseWriter, v validation.InputValidation, postProcess func() (interface{}, error)) {
	validationError := DecodeAndValidate(r, v)

	result := new(model.GenericResult)

	if validationError == nil {
		result = convert(postProcess())
	} else {
		result = convert(nil, validationError)
	}

	response.WriteResult(w, result)
}

func convert(data interface{}, err error) *model.GenericResult {
	if err == nil {
		result := model.Success()
		result.Data = data
		return result
	} else {
		return model.Error(err)
	}
}
