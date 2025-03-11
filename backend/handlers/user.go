package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/LilGreyCat/task-manager/models"
	"github.com/LilGreyCat/task-manager/repository"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	Repo repository.UserRepository
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request:", err)
		return
	}

	// Generate UUID if not provided
	if user.ID == uuid.Nil {
		newUUID, _ := uuid.NewV4()
		user.ID = newUUID
	}

	// Convert to Paris Time
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		return
	}
	user.CreatedAt = time.Now().In(loc)

	// Hash password
	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println("Error hashing password:", err)
		return
	}
	user.Password = hashedPassword

	// Insert into database
	ctx := context.Background()
	if err := h.Repo.CreateUser(ctx, &user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println("Database error:", err)
		return
	}

	// Hide password before returning response
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // Extract `id` from query params
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert string to UUID
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	// Fetch user from database
	ctx := context.Background()
	user, err := h.Repo.GetUserByID(ctx, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove password before returning user
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email") // Extract email from query params
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Fetch user from database
	ctx := context.Background()
	user, err := h.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove password before returning user
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	users, err := h.Repo.ListUsers(ctx)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure user ID is valid
	if user.ID == uuid.Nil {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Hash new password if provided
	if user.Password != "" {
		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}

	// Update user in database
	ctx := context.Background()
	if err := h.Repo.UpdateUser(ctx, &user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User updated successfully"}`))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := h.Repo.DeleteUser(ctx, id); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User deleted successfully"}`))
}
