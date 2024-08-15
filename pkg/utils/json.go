package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteJSONError(w http.ResponseWriter, status int, err error) {
	errMessage := strings.Trim(err.Error(), "\"")
	WriteJSON(w, status, errMessage)
}
