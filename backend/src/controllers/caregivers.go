package controllers

import (
	"dslic/models"
	u "dslic/utils"
	"encoding/json"
	"net/http"
)

var CreateCaregiver = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	caregiver := &models.Caregiver{}

	err := json.NewDecoder(r.Body).Decode(caregiver)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Create(caregiver)
	u.Respond(w, resp)
}

var ListCaregivers = func(w http.ResponseWriter, r *http.Request){
	caregivers := &[]models.Caregiver{}

	resp := models.List(caregivers)
	u.Respond(w, resp)
}

var GetCaregiver = func(w http.ResponseWriter, r *http.Request){
	caregiver := &models.Caregiver{}

	err := json.NewDecoder(r.Body).Decode(caregiver)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Get(caregiver)
	u.Respond(w, resp)
}


var UpdateCaregiver = func(w http.ResponseWriter, r *http.Request){
	obj := &models.Caregiver{}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Update(obj)
	u.Respond(w, resp)
}

var DeleteCaregiver = func(w http.ResponseWriter, r *http.Request){
	caregiver := &models.Caregiver{}

	err := json.NewDecoder(r.Body).Decode(caregiver)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	models.Delete(caregiver)
	u.Respond(w, caregiver)
}