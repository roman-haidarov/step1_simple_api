package api

import (
	"encoding/json"
	"net/http"
	"step1_simple_api/internal/types"
	"step1_simple_api/internal/users/password"
	generatedUsers "step1_simple_api/internal/web/users"
	"time"
)

func (api *API) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.users.GetUsers()
	if err != nil {
		api.WriteError(w, r, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := make([]generatedUsers.User, len(users))
	for i, user := range users {
		response[i] = api.convertToGeneratedUser(user)
	}

	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req generatedUsers.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	hashedPassword, salt, err := password.GeneratePasswordAndSalt(req.Password)
	if err != nil {
		api.WriteError(w, r, "Invalid password", http.StatusBadRequest)
		return
	}

	user := types.User{
		Email:     req.Email,
		Password:  &hashedPassword,
		Salt: 		 &salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userDB, err := api.users.CreateUser(user)
	if err != nil {
		api.WriteError(w, r, "Failed to create user", http.StatusBadRequest)
		return
	}

	response := api.convertToGeneratedUser(userDB)
	api.WriteJSON(w, r, response, http.StatusCreated)
}

func (api *API) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	user, err := api.users.GetUser(id)
	if err != nil {
		api.WriteError(w, r, "User not found", http.StatusNotFound)
		return
	}

	response := api.convertToGeneratedUser(user)
	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req generatedUsers.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	existingUser, err := api.users.GetUser(id)
	if err != nil {
		api.WriteError(w, r, "User not found", http.StatusNotFound)
		return
	}

	updatedUser := existingUser

	if req.Email != nil {
		updatedUser.Email = *req.Email
	}
	updatedUser.UpdatedAt = time.Now()

	if err = api.users.UpdateUser(updatedUser); err != nil {
		api.WriteError(w, r, "Failed to update user", http.StatusInternalServerError)
		return
	}

	response := api.convertToGeneratedUser(updatedUser)
	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) DestroyUser(w http.ResponseWriter, r *http.Request, id int) {
	if err := api.users.DestroyUser(id); err != nil {
		api.WriteError(w, r, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *API) convertToGeneratedUser(user types.User) generatedUsers.User {	
	generatedUser := generatedUsers.User{
		Id: user.ID,
		Email: user.Email,
	}

	if !user.CreatedAt.IsZero() {
		generatedUser.CreatedAt = &user.CreatedAt
	}
	if !user.UpdatedAt.IsZero() {
		generatedUser.UpdatedAt = &user.UpdatedAt
	}
	if user.DeletedAt != nil {
		generatedUser.DeletedAt = user.DeletedAt
	}

	return generatedUser
}
