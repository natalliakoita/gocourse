package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/natalliakoita/gocourse/lesson_09/datasource"
	"github.com/natalliakoita/gocourse/lesson_09/handlers"
)

func main() {
	conn := datasource.NewDB()
	defer conn.Close()

	// create struct
	ds := datasource.NewDS(conn)

	// create new handler
	h := handlers.NewContactHandler(&ds)

	router := mux.NewRouter()
	router.HandleFunc("/api/v0/contacts", h.AddContact).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/contacts", h.AddGroupToContact).Methods(http.MethodPut)
	router.HandleFunc("/api/v0/group/{id}/contacts", h.GetContactsOrderByGroup).Methods(http.MethodGet)

	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
