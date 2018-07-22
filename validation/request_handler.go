package validation

import (
	"net/http"
	"encoding/json"
)

func DecodeAndValidate(r *http.Request, v InputValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	return v.Validate(r)
}
