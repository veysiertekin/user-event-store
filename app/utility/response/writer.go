package response

import (
	"net/http"
	"encoding/json"
	"user-event-store/app/utility/response/model"
)

func WriteResult(w http.ResponseWriter, r *model.GenericResult) {
	w.Header().Set("Content-Type", "application/json")

	resultEncoder := json.NewEncoder(w)

	if r.Status == model.ERROR {
		w.WriteHeader(http.StatusConflict)
	}
	resultEncoder.Encode(r)
}
