package handlers

import(
	"net/http"
	"strconv"
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

	query := r.URL.Query()
	idParam := query.Get("id")

	//si no hay id, da el .json completp, si tiene lo convierte a id y regresa al aldeado especifico del id
	if idParam == "" {
		utils.WriteJSON(w, http.StatusOK, villagers)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil{
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	for _,  villager := range villagers {
		if villager.ID == id {
			utils.WriteJSON(w, http.StatusOK, villager)
			return
		}
	}

	http.Error(w, "Villager not found", http.StatusNotFound)
}