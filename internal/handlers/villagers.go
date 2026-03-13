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
		
		utils.WriteError(w,http.StatusInternalServerError, "Error loading villagers")
		return
	}

	query := r.URL.Query()

	idParam := query.Get("id")
	locationParam := query.Get("location")
	marriageableParam := query.Get("marriageable") 

	if idParam != ""{

		id, err := strconv.Atoi(idParam)
		if err != nil{
			
			utils.WriteError(w, http.StatusBadRequest, "Invalid id parameter")
			return
		}

		for _,  villager := range villagers {
			if villager.ID == id {
				utils.WriteJSON(w, http.StatusOK, villager)
				return
			}
		}
		utils.WriteError(w, http.StatusNotFound, "Villager not found")
		return
	}
	
	filtered := []models.Villager{}

	for _, v := range villagers{

		if locationParam != ""{
			if !strings.EqualFold(v.Location, locationParam){
			continue
			}
		}

		if marriageableParam != ""{
			m, err := strconv.ParseBool(marriageableParam)
			if err != nil{
				utils.WriteError(w, http.StatusBadRequest, "Invalid marriageable paramether")
				return
			}	

			if v.Marriageable != m {
				continue
			}
		}

		filtered = append(filtered, v)
	}

	if locationParam != "" || marriageableParam != ""{
		if len(filtered) == 0 {
			utils.WriteError(w, http.StatusNotFound, "Villager not found")
			return
		}			

		utils.WriteJSON(w, http.StatusOK, filtered)
		return
	}
	utils.WriteJSON(w, http.StatusOK, villagers)
}

//poder hacer un POST de un nuevo aldeano y se guarda en el villagers.json
func PostVillager(w http.ResponseWriter, r *http.Request){

	var newVillager models.Villager

	err := json.NewDecoder(r.Body).Decode(&newVillager)
	if err != nil {
		
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	validationError := utils.ValidateVillager(newVillager)
	if validationError != "" {
		utils.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		
		utils.WriteError(w, http.StatusInternalServerError, "Error loading villagers")
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
		
		utils.WriteError(w, http.StatusInternalServerError, "Error saving villagers")
		return
	}
	utils.WriteJSON(w, http.StatusCreated, newVillager)
}


func PutVillager(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		utils.WriteError(w,  http.StatusBadRequest, "Invalid id parameter")
		return
	}

	var updateVillager models.Villager

	err = json.NewDecoder(r.Body).Decode(&updateVillager)
	if err != nil {
		
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	validationError := utils.ValidateVillager(updateVillager)
	if validationError != "" {
		utils.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		
		utils.WriteError(w, http.StatusInternalServerError, "Error loading villagers")
		return
	}

	for i, villager := range villagers{
		if villager.ID == id {
			
			updateVillager.ID = id
			villagers[i] = updateVillager

			err = storage.SaveVillagers(villagers)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "Error saving villagers")
				return
			}

			utils.WriteJSON(w, http.StatusCreated, updateVillager)
			return
		}
	}
	utils.WriteError(w, http.StatusNotFound, "Villager not found")
}



func DeleteVillager(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query()
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		utils.WriteError(w,http.StatusBadRequest, "Invalid id parameter")
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error loading villagers")
		return
	}

	for i, villager := range villagers{
		if villager.ID == id {
			
			villagers = append(villagers[:i], villagers[i+1:]...)

			err = storage.SaveVillagers(villagers)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "Error saving villagers")
				return
			}

			utils.WriteJSON(w,http.StatusOK, map[string]string{ "message":"Villager deleted",})
			return

		}
	}

	utils.WriteError(w, http.StatusNotFound, "Villager not found")

}

func VillagerByIDHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	parts := strings.Split(path, "/")

	if len(parts) < 4 {
		utils.WriteError(w, http.StatusBadRequest, "Invalid path")
		return
	}

	idParam := parts[3]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	villagers, err := storage.LoadVillagers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error loading villagers")
		return
	}

	for _, v := range villagers {
		if v.ID == id {
			utils.WriteJSON(w, http.StatusOK, v)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "Villager not found")
}


