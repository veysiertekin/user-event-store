package controller

import (
	"net/http"
	"user-event-store/app/utility/request"
	"user-event-store/app/endpoint/user_event/model"
	"user-event-store/app/endpoint/user_event/service"
)

func CreateUserEvent(w http.ResponseWriter, r *http.Request) {
	var event model.UserEvent
	request.Handle(r, w, &event, "create-user-event", func() (interface{}, error) {
		return service.CreateUserEvent(event)
	})
}
