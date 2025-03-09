package routes

import (
	"net/http"

	"github.com/LilGreyCat/task-manager/handlers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userHandler *handlers.UserHandler) {
	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/users/email", userHandler.GetUserByEmail).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)
}
