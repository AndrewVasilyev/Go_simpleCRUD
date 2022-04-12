package handlers

import (
	"GO_SIMPLECRUD/models"
	"encoding/json"
	"net/http"
)

func (h dbHandler) GetAllPlants(w http.ResponseWriter, r *http.Request) {

	var plants []models.Plants

	if result := h.DB.Find(&plants); result.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Can Not Represent Plants"})

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"plant": plants})
}
