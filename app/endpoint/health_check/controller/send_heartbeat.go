package controller

import (
	"net/http"
	"user-event-store/app/utility/response"
	"user-event-store/app/utility/response/model"
)

func SendHeartbeat(w http.ResponseWriter, r *http.Request) {
	response.WriteResult(w, model.Success())
}
