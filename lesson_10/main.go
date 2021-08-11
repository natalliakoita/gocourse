package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/natalliakoita/gocourse/lesson_10/datasource"
	"github.com/natalliakoita/gocourse/lesson_10/handlers"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	client := datasource.NewDB(ctx)
	defer client.Disconnect(ctx)

	// create struct
	ds := datasource.NewDS(client)

	// create migrations
	addMigration := flag.String("migration", "false", "create migration")
	flag.Parse()
	if *addMigration == "true" {
		ds.AddGroups(ctx)
	}

	// create new handler
	h := handlers.NewContactHandler(&ds, ctx)

	router := mux.NewRouter()
	router.HandleFunc("/api/v0/contacts", h.AddContact).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/contacts", h.AddGroupToContact).Methods(http.MethodPut)
	router.HandleFunc("/api/v0/contacts", h.GetContactsOrderByGroup).Methods(http.MethodGet)

	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
