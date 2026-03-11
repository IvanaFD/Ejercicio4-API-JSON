package handlers

import(
	"net/http"
	"stardew_villagers/internal/storage"
	"stardew_villagers/internal/utils"
)


//handler HTTP que responde a la solicitud del endpoint /api/villagers, sirve para solicitar los datos al almacenamiento por el LoadVillagers y se mandan en JSON

func GetVillagers(w http.ResponseWriter, r *http.Request){

	villagers, err := storage.LoadVillagers()
	if err != nil {
		http.Error(w,"Error loading villagers", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, villagers)
}