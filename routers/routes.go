package routers

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	// Create a new Chi router
	router := chi.NewRouter()

	// Register the blogs routes
	router.Mount("/v1", V1Router())

	// Return the initialized router
	return router
}
