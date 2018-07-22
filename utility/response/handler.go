package response

import (
	"net/http"
	"encoding/json"
	"user-event-store/utility/response/model"
)

func HandleGenericResult(w http.ResponseWriter, validationError error) {
	w.Header().Set("Content-Type", "application/json")

	resultEncoder := json.NewEncoder(w)
	if validationError != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResult := model.Error(validationError.Error())
		resultEncoder.Encode(errorResult)
	} else {
		resultEncoder.Encode(model.Success())
	}
}
