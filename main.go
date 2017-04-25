package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"github.com/Latias94/taskmanager/common"
	"github.com/Latias94/taskmanager/routers"
)

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
