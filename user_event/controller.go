package user_event

import (
	"net/http"
	"user-event-store/validation"
	"user-event-store/response"
	"user-event-store/user_event/model"
)

func CreateUserEvent(w http.ResponseWriter, r *http.Request) {
	var userEvent model.UserEvent
	validationError := validation.DecodeAndValidate(r, &userEvent)

	response.HandleGenericResult(w, validationError)
}
