package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	respByteArray, err := json.Marshal(payload)
	fmt.Println(respByteArray)
	if err != nil {
		log.Printf("Failed to marshal to JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
 	w.WriteHeader(code)
	w.Write(respByteArray)
	fmt.Println(string(respByteArray))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with %v error:", msg)
	}

	type ErrorResponse struct {
		Error        string `json:"error"`
		ResponseCode int    `json:"code"`
	}

	respondWithJson(w, code, ErrorResponse{Error: msg, ResponseCode: code})
}
