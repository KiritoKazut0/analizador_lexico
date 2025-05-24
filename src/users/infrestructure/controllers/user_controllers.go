package controllers

import (
	"encoding/json"
	"net/http"
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
