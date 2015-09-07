package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jf-marino/simple.go.server/controllers"
	"github.com/jf-marino/simple.go.server/routes"
	"github.com/jf-marino/simple.go.server/storage"
)

func main() {
	log.Println("Starting server...")
	storage.Initialize()

	log.Println("Server initialized...")
	log.Println("Moving on to setting up routes...")
	people := new(routes.PeopleHandler)
	root := new(routes.RootHandler)

	http.Handle("/", root)
	http.Handle("/people", people)
	http.HandleFunc("/events", createEvent)
	http.HandleFunc("/events/all", selectEvents)

	log.Println("Finished setup...")
	http.ListenAndServe(":8080", nil)
	log.Println("Listening...")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	controllers.CreateEvent("SomeEventName")
	json.NewEncoder(w).Encode(nil)
}

func selectEvents(w http.ResponseWriter, r *http.Request) {
	events := controllers.GetEvents()
	json.NewEncoder(w).Encode(events)
}
