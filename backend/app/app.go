package app

import (
	"database/sql"

	"github.com/LilGreyCat/task-manager/config"
	"github.com/LilGreyCat/task-manager/handlers"
	"github.com/LilGreyCat/task-manager/repository"
)

type App struct {
	DB       *sql.DB
	Handlers *handlers.HandlerManager
}

func Initialize(cfg *config.Config) (*App, error) {
	// Open SQLite database
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	repoManager := repository.NewRepositoryManager(db)

	// Initialize handlers
	handlerManager := handlers.NewHandlerManager(repoManager)

	return &App{
		DB:       db,
		Handlers: handlerManager,
	}, nil
}
