package main

import (
  "dslic/app"
  "dslic/controllers"
  "fmt"
  "github.com/go-chi/chi"
  "github.com/go-chi/cors"
  "net/http"
  "os"
)

func main() {

  router := chi.NewRouter()

  /*cors := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
  })*/
  cors := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Authorization", "Content-Type"},
    AllowCredentials: true,
  })
  router.Use(cors.Handler)
  router.Use(app.JwtAuthentication)


  router.Post("/user/new", controllers.CreateAccount)
  router.Post("/user/update", controllers.UpdateAccount)
  router.Post("/user/login", controllers.Authenticate)

  router.Post("/patients/new", controllers.CreatePatient)
  router.Get("/patients/", controllers.ListPatients)
  router.Post("/patients/", controllers.GetPatient)
  router.Post("/patients/update", controllers.UpdatePatient)
  router.Post("/patients/delete", controllers.DeletePatient)

  router.Post("/caregivers/new", controllers.CreateCaregiver)
  router.Get("/caregivers/", controllers.ListCaregivers)
  router.Post("/caregivers/", controllers.GetCaregiver)
  router.Post("/caregivers/update", controllers.UpdateCaregiver)
  router.Post("/caregivers/delete", controllers.DeleteCaregiver)

  router.Post("/medication/new", controllers.CreateMedication)
  router.Get("/medication/", controllers.ListMedications)
  router.Post("/medication/", controllers.GetMedication)
  router.Post("/medication/update", controllers.UpdateMedication)
  router.Post("/medication/delete", controllers.DeleteMedication)

  router.Post("/prescription/new", controllers.CreatePrescription)
  router.Get("/prescription/", controllers.ListPrescriptions)
  router.Post("/prescription/", controllers.GetPrescription)
  router.Post("/prescription/update", controllers.UpdatePrescription)
  router.Post("/prescription/delete", controllers.DeletePrescription)

  router.Post("/medicationplans/new", controllers.CreateMedicationPlan)
  router.Get("/medicationplans/", controllers.ListMedicationPlans)
  router.Post("/medicationplans/", controllers.GetMedicationPlan)
  router.Post("/medicationplans/update", controllers.UpdateMedicationPlan)
  router.Post("/medicationplans/delete", controllers.DeleteMedicationPlan)
  //router.HandleFunc("/", controllers.GetContactsFor).Methods("GET")

  port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally

  fmt.Println(port)

  err := http.ListenAndServe("127.0.0.1:" + port, router) //Launch the app, visit localhost:8000/api
  if err != nil {
    fmt.Print(err)
  }
}