package controllers

import (
	"dslic/models"
	u "dslic/utils"
	"encoding/json"
	"net/http"
)

var CreateMedication = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	medication := &models.Medication{}

	err := json.NewDecoder(r.Body).Decode(medication)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Create(medication)
	u.Respond(w, resp)
}

var ListMedications = func(w http.ResponseWriter, r *http.Request){
	medications := &[]models.Medication{}

	resp := models.List(medications)
	u.Respond(w, resp)
}

var GetMedication = func(w http.ResponseWriter, r *http.Request){
	medication := &models.Medication{}

	err := json.NewDecoder(r.Body).Decode(medication)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Get(medication)
	u.Respond(w, resp)
}


var UpdateMedication = func(w http.ResponseWriter, r *http.Request){
	obj := &models.Medication{}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Update(obj)
	u.Respond(w, resp)
}

var DeleteMedication = func(w http.ResponseWriter, r *http.Request){
	medication := &models.Medication{}

	err := json.NewDecoder(r.Body).Decode(medication)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	models.Delete(medication)
	u.Respond(w, medication)
}