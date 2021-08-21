package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type AddContactRequest struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type AddContactResponse struct {
	Message string `json:"message"`
}

// write web response to client
func (c AddContactResponse) writeToWeb(w http.ResponseWriter) {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

// handle add contact request
func (u *ContactHandler) AddContact(w http.ResponseWriter, req *http.Request) {
	bodyData, err := ParseBodyRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req1 := &AddContactRequest{}
	if err = json.Unmarshal(bodyData, req1); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.ds.AddContact(req.Context(), req1.Name, req1.Number)
	if err != nil {
		result := AddContactResponse{
			Message: err.Error(),
		}
		result.writeToWeb(w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := AddContactResponse{
		Message: "Contact added",
	}
	result.writeToWeb(w)

	w.WriteHeader(http.StatusOK)
}
