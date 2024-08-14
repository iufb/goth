package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/iufb/goth/internal/db"
	"github.com/iufb/goth/internal/handlers"
	"github.com/iufb/goth/internal/handlers/api"
	"github.com/iufb/goth/internal/handlers/ui"
	"github.com/iufb/goth/internal/repositories"
	"github.com/iufb/goth/internal/services"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"http://127.0.0.1:7331"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// routes for ui
	r.Get("/", handlers.Make(ui.HandleHome))
	r.Get("/register", handlers.Make(ui.HandleRegister))

	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}
	// routes for rest api
	r.Route("/api/v1", func(r chi.Router) {
		userRepo := repositories.NewGormUserRepository(db)
		userService := services.NewUserService(userRepo)
		r.Post("/register", handlers.Make(api.RegisterHandler(userService)))
	})

	// serve static files
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public", fs))

	listerAddr := os.Getenv("PORT")
	slog.Info("Server started on port", "listerAddr", listerAddr)
	http.ListenAndServe(listerAddr, r)
}
