package health_check

import (
	"net/http"
	"user-event-store/response/model"
	"encoding/json"
)

func SendHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resultEncoder := json.NewEncoder(w)
	resultEncoder.Encode(model.Success())
}
