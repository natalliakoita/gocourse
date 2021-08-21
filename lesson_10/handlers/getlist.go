package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Contact struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Group  string `json:"group_name"`
}

type ContactListResponse struct {
	Contacts []Contact `json:"contacts"`
}

// write web response to client
func (c ContactListResponse) writeToWeb(w http.ResponseWriter) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func (u *ContactHandler) GetContactsOrderByGroup(w http.ResponseWriter, req *http.Request) {
	contacts, err := u.ds.GetContactsOrderByGroup(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := []Contact{}
	for _, contact := range contacts {
		q := Contact{}
		contact_id := contact.ID.String()
		q.ID = contact_id
		q.Name = contact.Name
		q.Number = contact.Number
		q.Group = contact.Group.Name

		response = append(response, q)
	}

	resp := ContactListResponse{}
	resp.Contacts = response

	resp.writeToWeb(w)

	w.WriteHeader(http.StatusOK)
}
