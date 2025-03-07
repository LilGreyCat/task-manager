package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LilGreyCat/task-manager/models"

	"github.com/gofrs/uuid"
)

// UserRepository defines the database operations for users.
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// userRepo is the concrete implementation of UserRepository.
type userRepo struct {
	db *sql.DB
}

// NewUserRepository returns an instance of UserRepository.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (id, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.ID.String(), user.Name, user.Email, user.Password, user.CreatedAt)
	return err
}

func (r *userRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id.String())

	var user models.User
	var idStr string
	if err := row.Scan(&idStr, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	user.ID, _ = uuid.FromString(idStr)
	return &user, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = ?"
	row := r.db.QueryRowContext(ctx, query, email)

	var user models.User
	var idStr string
	if err := row.Scan(&idStr, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	user.ID, _ = uuid.FromString(idStr)
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID.String())
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}
