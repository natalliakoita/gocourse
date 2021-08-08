package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Storage = map[string]Request{}

type Request struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type Response struct {
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

func parseRequest(r *http.Request) (*Request, error) {
	body := r.Body
	defer body.Close()
	bodyData, err := io.ReadAll(body)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("failed to read request body")
	}
	req := &Request{}
	if err = json.Unmarshal(bodyData, req); err != nil {
		log.Print(err.Error())
		return nil, errors.New("request format is wrong")
	}
	return req, nil
}

func (r Response) writeToWeb(w http.ResponseWriter) {
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func returnOK(w http.ResponseWriter, data interface{}) {
	res := Response{
		Error:  nil,
		Result: data,
	}
	res.writeToWeb(w)
}

func returnErr(w http.ResponseWriter, data interface{}) {
	res := Response{
		Error:  data,
		Result: nil,
	}
	res.writeToWeb(w)
}

func saveToken(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		returnErr(w, err.Error())
		return
	}
	_, ok := Storage[req.Token]
	if ok {
		returnErr(w, err.Error())
		return
	}
	Storage[req.Token] = *req
	
	returnOK(w, "Token successfully added")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/save", saveToken).Methods(http.MethodPost)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
