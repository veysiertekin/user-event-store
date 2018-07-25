package request

import (
	"fmt"
	"net/http"
	"user-event-store/graphite"
	"user-event-store/utility/response"
	"user-event-store/utility/response/model"
	"user-event-store/utility/validation"
	"time"
	"user-event-store/utility/time_helper"
	"user-event-store/utility/panic_helper"
)

const maxWaitDurationInMilliseconds = 100

func Handle(r *http.Request, w http.ResponseWriter, v validation.InputValidation, serviceName string, postProcess func() (data interface{}, err error)) {
	startTime := time.Now()
	defer sendMetric(startTime, serviceName)

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

func sendMetric(startTime time.Time, serviceName string) {
	duration := time.Since(startTime)
	timePassed := time_helper.ToMilliSeconds(duration)
	graphite.Client.SendMetricAsync("user-event-store.service."+serviceName, fmt.Sprintf("%.6f", timePassed))
}
