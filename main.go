package main

import (
	"net/http"

	"github.com/hipogeniq/Go_FullStack02/controllers"
	"github.com/hipogeniq/Go_FullStack02/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:3334",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
