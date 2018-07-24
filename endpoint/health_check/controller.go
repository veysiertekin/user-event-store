package health_check

import (
	"net/http"
	"user-event-store/utility/response/model"
	"user-event-store/utility/response"
)

func SendHeartbeat(w http.ResponseWriter, r *http.Request) {
	response.WriteResult(w, model.Success())
}
