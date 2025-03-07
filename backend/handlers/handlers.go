package handlers

import "github.com/LilGreyCat/task-manager/repository"

// Holds references to all handlers
type HandlerManager struct {
	UserHandler *UserHandler
}

// Initializes all handlers
func NewHandlerManager(repos *repository.RepositoryManager) *HandlerManager {
	return &HandlerManager{
		UserHandler: NewUserHandler(repos.UserRepo),
	}
}
