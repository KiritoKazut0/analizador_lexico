package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/KiritoKazut0/analizador-lexico/src/users/application"
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserController struct {
	application *application.UserUseCase
}

func NewUserController(service *application.UserUseCase) *UserController {
	return &UserController{application: service}
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}, success bool, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    data,
		"success": success,
		"message": message,
	})
}

func (c UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := c.application.GetAllUsers()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, false, "Error to get users")
		return
	}
	writeJSON(w, http.StatusOK, users, true, "Users retrieved successfully")
}

func (c UserController) GetAllUsersPaginated(w http.ResponseWriter, r *http.Request) {
	// Obtener parámetros de query
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	// Valores por defecto
	page := 1
	perPage := 100

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if perPageStr != "" {
		if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 && pp <= 1000 {
			perPage = pp
		}
	}

	paginatedResponse, err := c.application.GetAllUsersPaginated(page, perPage)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, false, "Error to get paginated users")
		return
	}

	writeJSON(w, http.StatusOK, paginatedResponse, true, "Paginated users retrieved successfully")
}

func (c UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params == nil || params["id"] == "" {
		writeJSON(w, http.StatusBadRequest, nil, false, "Clave is required")
		return
	}

	user, err := c.application.GetUserByID(params["id"])
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			writeJSON(w, http.StatusNotFound, nil, false, "User not found")
			return
		}
		writeJSON(w, http.StatusInternalServerError, nil, false, "Error to get user: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user, true, "User retrieved successfully")
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, false, "Complete all fields: "+err.Error())
		return
	}

	if err := c.application.CreateUser(&user); err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, false, "Failed to create user: "+err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, user, true, "User created successfully")
}

func (c UserController) CreateUsersBatch(w http.ResponseWriter, r *http.Request) {
	var batchRequest entities.BatchCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&batchRequest); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, false, "Invalid request format: "+err.Error())
		return
	}

	if len(batchRequest.Users) == 0 {
		writeJSON(w, http.StatusBadRequest, nil, false, "No users provided")
		return
	}


	if len(batchRequest.Users) > 10000 {
		writeJSON(w, http.StatusBadRequest, nil, false, "Too many users in batch. Maximum 10,000 users per request")
		return
	}

	if err := c.application.CreateUsersBatch(batchRequest.Users); err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, false, "Failed to create users batch: "+err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"created_count": len(batchRequest.Users),
	}, true, "Users batch created successfully")
}

func (c UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, false, "Complete all fields")
		return
	}

	updatedUser, err := c.application.UpdateUser(&user)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil, false, "Failed to update user: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, updatedUser, true, "User updated successfully")
}

func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params == nil || params["id"] == "" {
		writeJSON(w, http.StatusBadRequest, nil, false, "Clave is required")
		return
	}

	err := c.application.DeleteUser(params["id"])
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, false, "Failed to delete user: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, nil, true, "User deleted successfully")
}