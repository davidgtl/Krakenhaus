package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	//(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func Respond(w http.ResponseWriter, data interface{})  {
	w.Header().Add("Content-Type", "application/json")
	//setupResponse(&w)

	json.NewEncoder(w).Encode(data)
}