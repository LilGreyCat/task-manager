package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LilGreyCat/task-manager/app"
	"github.com/LilGreyCat/task-manager/config"
	"github.com/LilGreyCat/task-manager/routes"
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

	// Print clickable link
	fmt.Println("Server running at \033[1;34mhttp://localhost" + cfg.Port + "\033[0m 🚀")

	// Start HTTP server
	err = http.ListenAndServe(cfg.Port, router)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
