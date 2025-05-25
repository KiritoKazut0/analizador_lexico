package routers

import (
	"net/http"
	"github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, controller *controllers.UserController) {

	routerUser := router.PathPrefix("/users").Subrouter()

	routerUser.HandleFunc("",controller.GetAllUser).Methods(http.MethodGet)
	routerUser.HandleFunc("/{id}", controller.GetUserByID).Methods(http.MethodGet)
	routerUser.HandleFunc("", controller.CreateUser).Methods(http.MethodPost)
	routerUser.HandleFunc("", controller.UpdateUser).Methods(http.MethodPut)
	routerUser.HandleFunc("/{id}", controller.DeleteUser).Methods(http.MethodDelete)

	router.HandleFunc("/lotes", controller.CreateUsersBatch).Methods(http.MethodPost)
	router.HandleFunc("/paginated", controller.GetAllUsersPaginated).Methods(http.MethodGet)

}