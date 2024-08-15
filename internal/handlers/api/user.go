package api

import (
	"fmt"
	"net/http"

	"github.com/iufb/goth/internal/handlers"
	"github.com/iufb/goth/internal/models"
	"github.com/iufb/goth/internal/services"
	"github.com/iufb/goth/pkg/utils"
	"github.com/iufb/goth/views/shared"
)

func RegisterHandler(service *services.UserService) handlers.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var payload models.AuthPayload
		payload.Email = r.FormValue("email")
		payload.Password = r.FormValue("password")
		status, err := service.RegisterUser(&payload)
		if err != nil {
			// utils.WriteError(w, status, err)
			handlers.Render(w, r, shared.Error("Щи"), status)
			return fmt.Errorf("Register user error :  %v", err)
		}
		return utils.WriteJSON(w, http.StatusCreated, "Registered")
	}
}
