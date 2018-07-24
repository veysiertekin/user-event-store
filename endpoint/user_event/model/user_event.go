package model

import (
	"time"
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
	EventTime time.Time `json:"timestamp,omitempty"`
}

func (event UserEvent) Validate() error {
	if &event.ApiKey == nil || len(event.ApiKey) == 0 {
		return ErrInvalidApiKey
	}
	if &event.UserId == nil || event.UserId == 0 {
		return ErrInvalidUserId
	}
	if &event.EventTime == nil || event.EventTime.IsZero() {
		return ErrInvalidTimestamp
	}
	return nil
}
