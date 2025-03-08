package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LilGreyCat/task-manager/app"
	"github.com/LilGreyCat/task-manager/config"
	"github.com/LilGreyCat/task-manager/routes"
	"github.com/gorilla/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.LoadConfig()

	appInstance, err := app.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer appInstance.DB.Close()

	// Set up routes
	router := routes.SetupRoutes(&routes.Handlers{
		UserHandler: appInstance.Handlers.UserHandler,
	})

	// Enable CORS for Vite (port 5173)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	fmt.Println("Server running at http://localhost" + cfg.Port + " 🚀")
	err = http.ListenAndServe(cfg.Port, corsHandler)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
