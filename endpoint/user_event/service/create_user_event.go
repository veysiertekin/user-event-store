package service

import (
	"errors"
	"user-event-store/endpoint/user_event/model"
	"user-event-store/endpoint/user_event/repository/api_key_repository"
	"user-event-store/endpoint/user_event/repository/user_event_repository"
)

func CreateUserEvent(event model.UserEvent) (data interface{}, err error) {
	if _, err := api_key_repository.GetKey(event.ApiKey); err != nil {
		return nil, errors.New("invalid api key")
	}
	if data, err = user_event_repository.CreateUserEvent(event); err != nil {
		return nil, errors.New("user event could not be stored")
	}
	return
}
