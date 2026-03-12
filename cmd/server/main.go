package main

import (
	
	"log"
	"net/http"
	"stardew_villagers/internal/handlers"
	"stardew_villagers/internal/utils"
)

//Convierte en json la respuesta que GO no puede enviar json, solo se uso para el endpoint de prueba
type Message struct{
	Message string `json:"message"`
}

func main(){

	//endpoint de prueba para verificar que si responde el servidor
	http.HandleFunc("/api/piringo", pingHandler)
	//handler para path aprameters, al registrar con una barra al final Go lo trare como prefijo
	http.HandleFunc("/api/villagers/", handlers.VillagerByIDHandler)
	//handler principal, filtra la peticion segun el metodo HTTP
	http.HandleFunc("/api/villagers", func(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
        case http.MethodGet:
            handlers.GetVillagers(w, r)
        case http.MethodPost:
            handlers.PostVillager(w, r)
        case http.MethodPut:
            handlers.PutVillager(w, r)
        case http.MethodDelete:
            handlers.DeleteVillager(w, r)
        default:
            // Si el método no es ninguno de los anteriores, devolvemos 405
            utils.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
        }
	})
	
	//mensaje que esta sirviendo en el puerto 24785, si esta ocupado me da un error
	log.Println("Server running in: 24785")
	log.Fatal(http.ListenAndServe(":24785",nil))

}

//con el httpRequest contiene lo que manda el usuario, y responde con el httpResponseWriter
func pingHandler(w http.ResponseWriter, r *http.Request){
	response := Message{
		Message: "porongo",
	}
	utils.WriteJSON(w, http.StatusOK, response)
}

