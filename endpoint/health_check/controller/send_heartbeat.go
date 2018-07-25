package controller

import (
	"net/http"
	"user-event-store/utility/response"
	"user-event-store/utility/response/model"
)

func SendHeartbeat(w http.ResponseWriter, r *http.Request) {
	response.WriteResult(w, model.Success())
}
