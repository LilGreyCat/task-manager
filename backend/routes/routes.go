package routes

import (
	"github.com/LilGreyCat/task-manager/handlers"
	"github.com/gorilla/mux"
)

type Handlers struct {
	UserHandler *handlers.UserHandler
}

func SetupRoutes(h *Handlers) *mux.Router {
	r := mux.NewRouter()

	// Register User routes
	RegisterUserRoutes(r, h.UserHandler)

	return r
}
