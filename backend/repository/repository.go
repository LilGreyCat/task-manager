package repository

import "database/sql"

// Holds references to all repositories
type RepositoryManager struct {
	UserRepo UserRepository
}

// Initializes all repositories
func NewRepositoryManager(db *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		UserRepo: NewUserRepository(db),
	}
}
