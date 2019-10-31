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

  cors := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
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
  //router.HandleFunc("/", controllers.GetContactsFor).Methods("GET")

  port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally

  fmt.Println(port)

  err := http.ListenAndServe("127.0.0.1:" + port, router) //Launch the app, visit localhost:8000/api
  if err != nil {
    fmt.Print(err)
  }
}