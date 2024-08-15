package utils

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("Error : %s", err.Error())))
}
