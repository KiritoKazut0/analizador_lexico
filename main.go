package main

import (
	"github.com/KiritoKazut0/analizador-lexico/src/core"
	UserUseCase "github.com/KiritoKazut0/analizador-lexico/src/users/application"
	entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	UserController "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/controllers"
	UserDB "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/database"
	UserRouter "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// configure initialization
	db, err := core.ConnectMysql()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Starting migration...")
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Migration done successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}
	defer sqlDB.Close()

	userRepository := UserDB.NewUserMysqlRepository(db)
	userUseCase := UserUseCase.NewUserUseCase(userRepository)
	userControllor := UserController.NewUserController(userUseCase)

	router := mux.NewRouter()
	UserRouter.UserRoutes(router, userControllor)

	log.Println("Server listenin in http://localhost:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Error to setup server: %v", err)
	}

}
