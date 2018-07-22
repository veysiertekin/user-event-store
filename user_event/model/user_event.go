package model

import (
	"time"
	"net/http"
	"errors"
)

var (
	ErrInvalidApiKey    = errors.New("invalid api key")
	ErrInvalidUserId    = errors.New("invalid user id")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

type UserEvent struct {
	ApiKey    string    `json:"apiKey,omitempty"`
	UserId    int64     `json:"userId,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (user UserEvent) Validate(r *http.Request) error {
	if &user.ApiKey == nil || len(user.ApiKey) == 0 {
		return ErrInvalidApiKey
	}
	if &user.UserId == nil || user.UserId == 0 {
		return ErrInvalidUserId
	}
	if &user.Timestamp == nil || user.Timestamp.IsZero() {
		return ErrInvalidTimestamp
	}
	return nil
}
