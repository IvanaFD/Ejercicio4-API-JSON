package storage

import(
	"encoding/json"
	"os"
	"log"
	"stardew_villagers/internal/models"
)

//accede al villagers.josn, los lee y convierte JSON a estructura de datos de GO, que se representanta con un slice de villager
func LoadVillagers()([]models.Villager, error){

	file, err := os.ReadFile("./data/villagers.json")
	if err != nil{
		log.Fatal("Error reading file", err)
	}

	var villagers []models.Villager

	err = json.Unmarshal(file, &villagers)

	if err != nil{
		log.Fatal("Error parsing JSON", err)
	}

	return villagers, nil
}