package ui

import (
	"net/http"

	"github.com/iufb/goth/internal/handlers"
	"github.com/iufb/goth/views/auth"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) error {
	return handlers.Render(w, r, auth.Register(), http.StatusOK)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return handlers.Render(w, r, auth.Login(), http.StatusOK)
}
