package handlers

import (
	"GO_PROJ/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h dbHandler) UpdatePlant(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	var updatedPlant models.Plants
	json.Unmarshal(body, &updatedPlant)

	var plant models.Plants

	if result := h.DB.First(&plant, id); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Defined plant was not found")

		return
	}

	plant = updatedPlant

	h.DB.Save(&plant)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Defined plant was updated")
}
