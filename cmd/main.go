package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sagarbgohil/go-backend/config"
	"github.com/sagarbgohil/go-backend/routers"
)

func main() {
	// Load configuration
	_, err := config.Load()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Initialize MongoDB client
	client, err := config.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
	}()

	// Initialize the router
	router := routers.InitRouter()
	log.Println("Server is running on http://localhost:" + config.Constants.Port)
	log.Fatal(http.ListenAndServe(":"+config.Constants.Port, router))
}
