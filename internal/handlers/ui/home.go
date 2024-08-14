package ui

import (
	"net/http"

	"github.com/iufb/goth/internal/handlers"
	"github.com/iufb/goth/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return handlers.Render(w, r, home.Index())
}
