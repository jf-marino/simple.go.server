package controllers

import (
	"log"

	"github.com/jf-marino/simple.go.server/storage"
	"github.com/jf-marino/simple.go.server/storage/models"
)

func CreateEvent(name string) {
	store := storage.Orm()
	event := new(models.Event)
	event.Name = name
	id, err := store.Insert(event)
	if err == nil {
		log.Println("Id for the new event is", id)
		event.Id = id
	} else {
		log.Println(err)
	}
}

func GetEvents() (events []*models.Event) {
	store := storage.Orm()
	num, err := store.QueryTable("event").All(&events)
	if err == nil {
		log.Println(num, " events found")
		printEvents(events)
	}
	return
}

func printEvents(events []*models.Event) {
	for _, evt := range events {
		log.Println(evt.Id, " -> ", evt.Name)
	}
}
