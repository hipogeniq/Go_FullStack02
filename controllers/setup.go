package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hipogeniq/Go_FullStack02/utils"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("Get")
	router.HandleFunc("/quests", GetAllQuests).Methods("GET")
	router.HandleFunc("/quest/{id}", GetQuest).Methods("GET")
	router.HandleFunc("/quest", CreateQuest).Methods("POST")
	router.HandleFunc("/quest/{id}", UpdateQuest).Methods("PUT")
	router.HandleFunc("/quest/{id}", DeleteQuest).Methods("DELETE")

	return router
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	utils.RespondWithError(w, http.StatusFound, "Home Of Quests")
}
