package metric

import (
	"time"
	"user-event-store/app/client/logstash"
	"user-event-store/app/utility/time_helper"
)

func SendResponseTimeMetricAsync(startTime time.Time, serviceName string) {
	duration := time.Since(startTime)
	timePassed := time_helper.ToMilliSeconds(duration)

	message := ServiceResponseTime{Application: "user-event-store", ServiceName: serviceName, TimeElapsed: timePassed}
	go logstash.Client.SendJsonAsync(message)
}

type ServiceResponseTime struct {
	Application string  `json:"application,omitempty"`
	ServiceName string  `json:"serviceName,omitempty"`
	TimeElapsed float64 `json:"timeElapsed,omitempty"`
}
