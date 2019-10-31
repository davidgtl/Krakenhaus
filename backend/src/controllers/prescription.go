package controllers

import (
	"dslic/models"
	u "dslic/utils"
	"encoding/json"
	"net/http"
)

var CreatePrescription = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	prescription := &models.Prescription{}

	err := json.NewDecoder(r.Body).Decode(prescription)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Create(prescription)
	u.Respond(w, resp)
}

var ListPrescriptions = func(w http.ResponseWriter, r *http.Request){
	prescriptions := &[]models.Prescription{}

	resp := models.List(prescriptions)
	u.Respond(w, resp)
}

var GetPrescription = func(w http.ResponseWriter, r *http.Request){
	prescription := &models.Prescription{}

	err := json.NewDecoder(r.Body).Decode(prescription)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Get(prescription)
	u.Respond(w, resp)
}


var UpdatePrescription = func(w http.ResponseWriter, r *http.Request){
	obj := &models.Prescription{}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Update(obj)
	u.Respond(w, resp)
}

var DeletePrescription = func(w http.ResponseWriter, r *http.Request){
	prescription := &models.Prescription{}

	err := json.NewDecoder(r.Body).Decode(prescription)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	models.Delete(prescription)
	u.Respond(w, prescription)
}