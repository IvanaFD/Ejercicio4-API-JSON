package utils


import (
	"stardew_villagers/internal/models"
	"encoding/json"
	"net/http"
)

type ErrorResponse struct{
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, message string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Error:message,
	})
}


func ValidateVillager(v models.Villager) string {

	if v.Name == "" {
		return "name is required"
	}

	if v.Birthday == "" {
		return "birthday is required"
	}

	if v.Location == "" {
		return "location is required"
	}

	if len(v.BestGifts) == 0 {
		return "best_gifts must contain at least one item"
	}

	return ""
}