package request

import (
	"net/http"
	"user-event-store/utility/response"
	"user-event-store/utility/response/model"
	"user-event-store/utility/validation"
	"fmt"
)

func Handle(r *http.Request, w http.ResponseWriter, v validation.InputValidation, postProcess func() (data interface{}, err error)) {
	validationError := DecodeAndValidate(r, v)

	result := new(model.GenericResult)

	if validationError == nil {
		result = wrapWithGenericTemplate(recoverOnAnyPanic(postProcess))
	} else {
		result = wrapWithGenericTemplate(nil, validationError)
	}

	response.WriteResult(w, result)
}

func recoverOnAnyPanic(postProcess func() (interface{}, error)) (data interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	return postProcess()
}

func wrapWithGenericTemplate(data interface{}, err error) *model.GenericResult {
	if err == nil {
		result := model.Success()
		result.Data = data
		return result
	} else {
		return model.Error(err)
	}
}
