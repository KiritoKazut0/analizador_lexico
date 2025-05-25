package main

import (
	core "github.com/KiritoKazut0/analizador-lexico/src/core"
	UserUseCase "github.com/KiritoKazut0/analizador-lexico/src/users/application"
	entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	UserRedis "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/cache"
	UserController "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/controllers"
	UserDB "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/database"
	UserRouter "github.com/KiritoKazut0/analizador-lexico/src/users/infrestructure/routers"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	// configure initialization
	db, err := core.ConnectMysql()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	redisClient, err := core.ConnectRedisClient()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
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
	userCacheRepository := UserRedis.NewUserRepository(redisClient)
	userUseCase := UserUseCase.NewUserUseCase(userRepository, userCacheRepository)
	userControllor := UserController.NewUserController(userUseCase)

	router := mux.NewRouter()
	UserRouter.UserRoutes(router, userControllor)

	log.Println("Server listenin in http://localhost:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Error to setup server: %v", err)
	}

}
