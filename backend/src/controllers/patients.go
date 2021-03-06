package controllers

import (
	"dslic/models"
	u "dslic/utils"
	"encoding/json"
	"net/http"
)

var CreatePatient = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	patient := &models.Patient{}

	err := json.NewDecoder(r.Body).Decode(patient)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Create(patient)
	u.Respond(w, resp)
}

var ListPatients = func(w http.ResponseWriter, r *http.Request){
	patients := &[]models.Patient{}

	resp := models.List(patients)
	u.Respond(w, resp)
}

var GetPatient = func(w http.ResponseWriter, r *http.Request){
	patient := &models.Patient{}

	err := json.NewDecoder(r.Body).Decode(patient)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Get(patient)
	u.Respond(w, resp)
}


var UpdatePatient = func(w http.ResponseWriter, r *http.Request){
	obj := &models.Patient{}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Update(obj)
	u.Respond(w, resp)
}

var DeletePatient = func(w http.ResponseWriter, r *http.Request){
	patient := &models.Patient{}

	err := json.NewDecoder(r.Body).Decode(patient)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	models.Delete(patient)
	u.Respond(w, patient)
}