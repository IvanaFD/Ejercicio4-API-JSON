package handlers

import (
	"net/http"
	"stardew_villagers/internal/utils"
)

func APIInfoHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"api": "Stardew Villagers API",
		"version": "1.0",
		"author":"Ivana Figueroa",
		"description": "Simple REST API to manage Stardew Valley villagers",
		"endpoints": []string{
			"GET /api/villagers",
			"GET /api/villagers?id=1",
			"GET /api/villagers/{id}",
			"POST /api/villagers",
			"PUT /api/villagers?id=1",
			"DELETE /api/villagers?id=1",
		},
	}

	utils.WriteJSON(w, http.StatusOK, response)
}