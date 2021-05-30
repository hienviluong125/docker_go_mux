package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hienviluong125/docker_go_mux/controllers"
	"github.com/hienviluong125/docker_go_mux/database"
	"github.com/joho/godotenv"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	if goEnv := os.Getenv("GO_ENV"); goEnv != "production" {
		log.Println("Starting go server in development mode")
		err := godotenv.Load(".env")

		if err != nil {
			panic("Env vars error. Please ensure .env file is created")
		}
	} else if goEnv == "production" {
		log.Println("Starting go server in production mode")
	}

	database.Connect()

	r := mux.NewRouter()
	postController := controllers.PostController{}

	r.Use(loggingMiddleware)
	r.HandleFunc("/api/posts", postController.Index).Methods("GET")
	r.HandleFunc("/api/posts/{id}", postController.Show).Methods("GET")
	r.HandleFunc("/api/posts", postController.Create).Methods("POST")
	r.HandleFunc("/api/posts/{id}", postController.Update).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", postController.Destroy).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
