package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Convierte en json la respuesta que GO no puede enviar json
type Message struct{
	Message string `json:"message"`
}

func main(){

	//respuesta de prueba de pin-pong, si recibe la ruta /api/ping devuelve el resultado de la funcion pinghandler
	http.HandleFunc("/api/piringo", pingHandler)
	
	//mensaje en cnsola que esta sirviendo en el puerto 24785, si esta ocupado me da un error
	log.Println("Server running in: 24785")
	log.Fatal(http.ListenAndServe(":24785",nil))

}

//con el httpRequest contiene lo que manda el usuario, y responde con el httpResponseWriter
func pingHandler(w http.ResponseWriter, r *http.Request){
	response := Message{
		Message: "porongo",
	}
	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil{
		http.Error(w,"Internal Server Erro", http.StatusInternalServerError)
	}
}