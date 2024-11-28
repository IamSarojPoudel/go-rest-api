package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/internal/config"
	"rest-api/internal/database"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"
	"rest-api/internal/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading.env file")
	}
	cfg := config.LoadConfig()
	db := &database.Database{}
	db.ConnectDB(cfg.DatabaseURL)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	handlers.InitAuthHandler(authService)

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Server starting on", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatal(err)
	}

}
