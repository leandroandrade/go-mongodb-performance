package main

import (
	"github.com/leandroandrade/go-mongodb/service"
	"github.com/leandroandrade/go-mongodb/database"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	mongo := database.NewMongoInstance()
	handler := service.NewHandler(mongo)

	router := mux.NewRouter().StrictSlash(true)
	users := router.PathPrefix("/users").Subrouter()
	users.Methods("POST").HandlerFunc(handler.Home)

	metrics := router.PathPrefix("/metrics").Subrouter()
	metrics.Methods("GET").HandlerFunc(handler.PrintMemory)

	negr := negroni.Classic()

	negr.UseHandler(router)
	negr.Run(":3000")
}
