package main

import (
  "dslic/app"
  "dslic/controllers"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "os"
)

func main() {

  router := mux.NewRouter()

  router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
  router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

  router.HandleFunc("/patients/new", controllers.CreatePatient).Methods("POST")
  router.HandleFunc("/patients/", controllers.ListPatients).Methods("GET")
  router.HandleFunc("/patients/", controllers.GetPatient).Methods("POST")
  router.HandleFunc("/patients/update", controllers.UpdatePatient).Methods("POST")
  router.HandleFunc("/patients/delete", controllers.DeletePatient).Methods("POST")
  //router.HandleFunc("/", controllers.GetContactsFor).Methods("GET")

  router.Use(app.JwtAuthentication) //attach JWT auth middleware

  port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally

  fmt.Println(port)

  err := http.ListenAndServe("127.0.0.1:" + port, router) //Launch the app, visit localhost:8000/api
  if err != nil {
    fmt.Print(err)
  }
}