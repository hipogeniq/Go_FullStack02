package main

import (
	"log"
	"net/http"

	"github.com/hipogeniq/Go_FullStack02/controllers"
	"github.com/hipogeniq/Go_FullStack02/models"
)

func main() {

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:3334",
		Handler: handler,
	}

	models.ConnectDatabase()

	err := server.ListenAndServe()
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
