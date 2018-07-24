package request

import (
	"net/http"
	"encoding/json"
	"user-event-store/utility/validation"
)

func DecodeAndValidate(r *http.Request, v validation.InputValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	return v.Validate(r)
}
