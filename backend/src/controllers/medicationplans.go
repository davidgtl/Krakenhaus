package controllers

import (
	"dslic/models"
	u "dslic/utils"
	"encoding/json"
	"net/http"
)

var CreateMedicationPlan = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	medicationplan := &models.MedicationPlan{}

	err := json.NewDecoder(r.Body).Decode(medicationplan)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Create(medicationplan)
	u.Respond(w, resp)
}

var ListMedicationPlans = func(w http.ResponseWriter, r *http.Request){
	medicationplans := &[]models.MedicationPlan{}

	resp := models.List(medicationplans)
	u.Respond(w, resp)
}

var GetMedicationPlan = func(w http.ResponseWriter, r *http.Request){
	medicationplan := &models.MedicationPlan{}

	err := json.NewDecoder(r.Body).Decode(medicationplan)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Get(medicationplan)
	u.Respond(w, resp)
}


var UpdateMedicationPlan = func(w http.ResponseWriter, r *http.Request){
	obj := &models.MedicationPlan{}

	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Update(obj)
	u.Respond(w, resp)
}

var DeleteMedicationPlan = func(w http.ResponseWriter, r *http.Request){
	medicationplan := &models.MedicationPlan{}

	err := json.NewDecoder(r.Body).Decode(medicationplan)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	models.Delete(medicationplan)
	u.Respond(w, medicationplan)
}