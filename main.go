package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"tv-guide/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	port, exists := os.LookupEnv("PORT")
	if exists == false {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	var err error

	log.Println("Application is running on localhost:" + port)

	log.Fatal(srv.ListenAndServe())

	if err != nil {
		log.Fatal(err)
	}
}
