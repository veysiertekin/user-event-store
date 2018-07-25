package api_key_repository

import (
	"user-event-store/app/client/cassandra"
	"user-event-store/app/endpoint/user_event/model"
)

func GetKey(apiKey string) (*model.ApiKey, error) {
	var key string
	if err := cassandra.Session.Query("SELECT api_key FROM api_key where api_key=?", apiKey).Scan(&key); err != nil {
		return nil, err
	}
	return &model.ApiKey{ApiKey: key}, nil
}
