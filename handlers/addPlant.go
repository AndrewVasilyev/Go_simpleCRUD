package handlers

import (
	"GO_SIMPLECRUD/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h dbHandler) AddPlant(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	var plants models.Plants

	json.Unmarshal(body, &plants)

	if result := h.DB.Create(&plants); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{"error": "Record Can Not Be Created"})

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"data": "Record Created Successfuly"})

}
