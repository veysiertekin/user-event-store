package user_event

import (
	"net/http"
	"user-event-store/utility/validation"
	"user-event-store/utility/response"
	"user-event-store/endpoint/user_event/model"
)

func CreateUserEvent(w http.ResponseWriter, r *http.Request) {
	var userEvent model.UserEvent
	validationError := validation.DecodeAndValidate(r, &userEvent)

	response.HandleGenericResult(w, validationError)
}
