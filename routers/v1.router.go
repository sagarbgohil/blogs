package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sagarbgohil/go-backend/api/blogs"
)

func V1Router() *chi.Mux {
	// Create a new Chi router
	router := chi.NewRouter()

	// Register the blogs routes
	router.Mount("/blogs", blogs.V1Router())


	// Return the initialized router
	return router
}