package user_event_repository

import (
	"user-event-store/endpoint/user_event/model"
	"user-event-store/cassandra"
)

func CreateUserEvent(event model.UserEvent) (*model.UserEvent, error) {
	if err := cassandra.Session.Query("INSERT INTO user_event (api_key, user_id, event_time) values(?,?,?)", event.ApiKey, event.UserId, event.EventTime).Exec(); err != nil {
		return nil, err
	}
	return &event, nil
}
