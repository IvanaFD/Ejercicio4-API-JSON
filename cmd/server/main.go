package main

import (
	
	"log"
	"net/http"
	"stardew_villagers/internal/handlers"
	"stardew_villagers/internal/utils"
)

//Convierte en json la respuesta que GO no puede enviar json
type Message struct{
	Message string `json:"message"`
}

func main(){

	//respuesta de prueba de pin-pong, si recibe la ruta /api/ping devuelve el resultado de la funcion pinghandler
	http.HandleFunc("/api/piringo", pingHandler)
	http.HandleFunc("/api/villagers", func(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		handlers.GetVillagers(w, r)
		return
	}

	if r.Method == http.MethodPost {
		handlers.CreateVillager(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
})
	
	//mensaje en cnsola que esta sirviendo en el puerto 24785, si esta ocupado me da un error
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

