package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/natalliakoita/gocourse/lesson_10/datasource"
)

type ContactHandler struct {
	ds  *datasource.DS
	ctx context.Context
}

func NewContactHandler(d *datasource.DS, ctx context.Context) ContactHandler {
	cont := ContactHandler{
		ds:  d,
		ctx: ctx,
	}
	return cont
}

func ParseBodyRequest(r *http.Request) ([]byte, error) {
	body := r.Body
	defer body.Close()
	bodyData, err := io.ReadAll(body)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("failed to read request body")
	}
	return bodyData, nil
}

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

type AddGroupToContactRequest struct {
	Name      string `json:"name"`
	Number    string `json:"number"`
	GroupName string `json:"group_name"`
}

type AddGroupToContactResponse struct {
	Message string `json:"message"`
}

func (u *ContactHandler) AddGroupToContact(w http.ResponseWriter, req *http.Request) {
	bodyData, err := ParseBodyRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqStr := &AddGroupToContactRequest{}
	if err = json.Unmarshal(bodyData, reqStr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = u.ds.AddGroupToContact(req.Context(), reqStr.Name, reqStr.Number, reqStr.GroupName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := AddContactResponse{
		Message: "Group for contact added",
	}
	result.writeToWeb(w)

	w.WriteHeader(http.StatusOK)
}

type ContactGroup struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	GroupName string `json:"group_name"`
}

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
