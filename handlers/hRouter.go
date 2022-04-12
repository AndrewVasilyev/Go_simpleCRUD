package handlers

import (
	"GO_SIMPLECRUD/db"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	dbHandler := NewDBHandler(db.NewDB())

	router.HandleFunc("/plants", dbHandler.GetAllPlants).Methods(http.MethodGet)
	router.HandleFunc("/plants", dbHandler.AddPlant).Methods(http.MethodPost)
	router.HandleFunc("/plants/{id}", dbHandler.GetPlant).Methods(http.MethodGet)
	router.HandleFunc("/plants/{id}", dbHandler.UpdatePlant).Methods(http.MethodPut)
	router.HandleFunc("/plants/{id}", dbHandler.DeletePlant).Methods(http.MethodDelete)

	return router
}
