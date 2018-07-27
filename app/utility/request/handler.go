package request

import (
	"time"
	"net/http"
	"user-event-store/app/utility/validation"
	"user-event-store/app/utility/time_helper"
	"user-event-store/app/utility/panic_helper"
	"user-event-store/app/utility/response"
	"user-event-store/app/utility/response/metric"
	"user-event-store/app/utility/response/model"
)

const maxWaitDurationInMilliseconds = 100

func Handle(r *http.Request, w http.ResponseWriter, v validation.InputValidation, serviceName string, postProcess func() (data interface{}, err error)) {
	startTime := time.Now()
	defer metric.SendResponseTimeMetricAsync(startTime, serviceName)

	result := validateAndExecute(r, v, postProcess)

	time_helper.WaitForARandomDuration(maxWaitDurationInMilliseconds)
	response.WriteResult(w, result)
}

func validateAndExecute(r *http.Request, v validation.InputValidation, method func() (data interface{}, err error)) *model.GenericResult {
	validationError := DecodeAndValidate(r, v)

	result := new(model.GenericResult)
	if validationError == nil {
		result = wrapWithGenericTemplate(processAndRecover(method))
	} else {
		result = wrapWithGenericTemplate(nil, validationError)
	}
	return result
}

func processAndRecover(method func() (interface{}, error)) (data interface{}, err error) {
	defer panic_helper.RecoverOnError(&err)
	return method()
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
