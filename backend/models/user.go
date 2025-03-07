package models

import (
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID        uuid.UUID `json:"id"`         // UUID as primary key
	Name      string    `json:"name"`       // User's name
	Email     string    `json:"email"`      // Unique email
	Password  string    `json:"password"`   // Hashed password
	CreatedAt time.Time `json:"created_at"` // Timestamp of user creation
}

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compares a plaintext password with a hashed one.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
