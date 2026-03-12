package handlers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"stardew_villagers/internal/storage"
	"stardew_villagers/internal/models"
	"stardew_villagers/internal/utils"
	"strings"
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


//poder hacer un POST de un nuevo aldeano y se guarda en el villagers.json
func PostVillager(w http.ResponseWriter, r *http.Request){

	var newVillager models.Villager

	err := json.NewDecoder(r.Body).Decode(&newVillager)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		http.Error(w, "Error loading villagers", http.StatusInternalServerError)
		return
	}

	maxID := 0

	for _, v := range villagers {
		if v.ID > maxID {
			maxID = v.ID
		}
	}

	newVillager.ID = maxID + 1

	villagers = append(villagers, newVillager)
	
	err = storage.SaveVillagers(villagers)
	if err != nil {
		http.Error(w, "Error saving villagers", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, newVillager)
}


func PutVillager(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var updateVillager models.Villager

	err = json.NewDecoder(r.Body).Decode(&updateVillager)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		http.Error(w, "Error loading villagers", http.StatusInternalServerError)
		return
	}

	for i, villager := range villagers{
		if villager.ID == id {
			
			updateVillager.ID = id
			villagers[i] = updateVillager

			err = storage.SaveVillagers(villagers)
			if err != nil {
				http.Error(w, "Error saving villagers", http.StatusInternalServerError)
				return
			}

			utils.WriteJSON(w, http.StatusCreated, updateVillager)
			return
		}
	}
	http.Error(w, "Villager not found", http.StatusNotFound)
}



func DeleteVillager(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		http.Error(w, "Error loading villagers", http.StatusInternalServerError)
		return
	}

	for i, villager := range villagers{
		if villager.ID == id {
			
			villagers = append(villagers[:i], villagers[i+1:]...)

			err = storage.SaveVillagers(villagers)
			if err != nil {
				http.Error(w, "Error saving villagers", http.StatusInternalServerError)
				return
			}

			utils.WriteJSON(w,http.StatusOK, map[string]string{ "message":"Villager deleted",})
			return

		}
	}

	http.Error(w, "Villager not found", http.StatusNotFound)

}

func VillagerByIDHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	parts := strings.Split(path, "/")

	if len(parts) < 4 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	idParam := parts[3]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		http.Error(w, "Error loading villagers", http.StatusInternalServerError)
		return
	}

	for _, v := range villagers {
		if v.ID == id {
			utils.WriteJSON(w, http.StatusOK, v)
			return
		}
	}

	http.Error(w, "Villager not found", http.StatusNotFound)
}


