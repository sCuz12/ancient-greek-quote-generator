package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondToJson(w,code,errResponse{
		Error: message,
	})
}
func respondToJson(w http.ResponseWriter,code int,payload interface{}) {
	data,err := json.Marshal(payload)

	if err != nil{ 
		log.Println("Failed to marshal JSON response")
		w.WriteHeader(500) //internal server error
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)	
	w.Write(data)

}