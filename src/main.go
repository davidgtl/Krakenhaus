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
  router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
  router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
  router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
  router.HandleFunc("/", controllers.GetContactsFor).Methods("GET")

  router.Use(app.JwtAuthentication) //attach JWT auth middleware

  port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally

  fmt.Println(port)

  err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
  if err != nil {
    fmt.Print(err)
  }
}