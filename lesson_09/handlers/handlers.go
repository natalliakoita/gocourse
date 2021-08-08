package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/natalliakoita/gocourse/lesson_09/datasource"
)

type ContactHandler struct {
	ds *datasource.DS
}

func NewContactHandler(d *datasource.DS) ContactHandler {
	cont := ContactHandler{
		ds: d,
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
	Group  string `json:"group_id"`
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
	err = u.ds.AddContact(req1.Name, req1.Number)
	if err != nil {
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
	Name   string `json:"name"`
	Number string `json:"number"`
	Group  int    `json:"group_id"`
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

	err = u.ds.AddGroupToContact(reqStr.Name, reqStr.Number, reqStr.Group)
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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	GroupName string `json:"group_name"`
}

type ContactListResponse struct {
	Contacts []ContactGroup `json:"contacts"`
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
	vars := mux.Vars(req)
	group_id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(group_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	contacts, err := u.ds.GetContactsOrderByGroup(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := []ContactGroup{}
	for _, contact := range contacts {
		q := ContactGroup{}
		q.ID = contact.Contact.ID
		q.Name = contact.Contact.Name
		q.Number = contact.Contact.Number
		q.GroupName = contact.Group.Name

		response = append(response, q)
	}

	resp := ContactListResponse{}
	resp.Contacts = response

	resp.writeToWeb(w)

	w.WriteHeader(http.StatusOK)
}

