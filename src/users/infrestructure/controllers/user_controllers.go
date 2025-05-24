package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/KiritoKazut0/analizador-lexico/src/users/application"
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"github.com/gorilla/mux"
)

type UserController struct {
	application *application.UserUseCase
}

func NewUserController(service *application.UserUseCase) *UserController {
	return &UserController{application: service}
}

func (c UserController) GetAllUser(w http.ResponseWriter, r *http.Request){
	users, err := c.application.GetAllUsers()
	
	if err != nil {
		http.Error(w, "Error to get users" + err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, users, http.StatusOK)
}

func (c UserController) GetUserByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	if params == nil{
		writeJSON(w, "Is required the clave", http.StatusBadRequest)
		return
	}
	user, err := c.application.GetUserByID(params["id"])

	if err != nil {
		writeJSON(w, "error to get user" + err.Error(), http.StatusInternalServerError)
		return
	} else if (user.Clave == ""){
		writeJSON(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, user, http.StatusOK)

}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSON(w,"is required complete all fields" + err.Error(), http.StatusBadRequest)
		return
	}

	 err := c.application.CreateUser(&user)

	if err != nil {
		writeJSON(w, "Failed to create user"+ err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, "create user succefully", http.StatusCreated)


}


func (c UserController) UpdateUser(w http.ResponseWriter, r *http.Request){
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		writeJSON(w, "is required complete all fields", http.StatusBadRequest)
		return
	}

	newUser ,err := c.application.UpdateUser(&user)	

	if err != nil {
		writeJSON(w, "failed to update user", http.StatusBadRequest)
		return
	}

	writeJSON(w, newUser, http.StatusOK)

}

func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	if params == nil{
		writeJSON(w, "Is required the clave", http.StatusBadRequest)
		return
	}

     err :=c.application.DeleteUser(params["id"])

	 if err != nil {
		writeJSON(w, "User not found"+ err.Error(), http.StatusInternalServerError)
	 }

	 writeJSON(w, "delete user is succefully", http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}