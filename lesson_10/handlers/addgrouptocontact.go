package handlers

import (
	"encoding/json"
	"net/http"
)

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
