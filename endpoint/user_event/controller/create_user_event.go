package controller

import (
	"net/http"
	"user-event-store/utility/request"
	"user-event-store/endpoint/user_event/model"
	"user-event-store/endpoint/user_event/service"
)

func CreateUserEvent(w http.ResponseWriter, r *http.Request) {
	var event model.UserEvent
	request.Handle(r, w, &event, func() (interface{}, error) {
		return service.CreateUserEvent(event)
	})
}
