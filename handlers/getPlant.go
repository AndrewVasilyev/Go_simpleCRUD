package handlers

import (
	"GO_PROJ/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h dbHandler) GetPlant(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	var plant models.Plants

	if result := h.DB.First(&plant, id); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Defined plant was not found")

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plant)

}
