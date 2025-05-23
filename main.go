package main

import (
	"log"
	"net/http"

	"github.com/KiritoKazut0/analizador-lexico/src/core"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// configure initialization

	db, err := core.ConnectMysql()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)

	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}
	sqlDB.Close()

	http.ListenAndServe(":3000", r)

}
