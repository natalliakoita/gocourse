package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func requestInfo(w http.ResponseWriter, req *http.Request) {
	var ins response

	m := make(map[string][]string)

	for i, header := range req.Header {
		m[i] = header
	}

	ins.Host = req.Host
	ins.Headers = m
	ins.RequestURI = req.RequestURI
	ins.UserAgent = req.Header["User-Agent"][0]

	js, err := json.Marshal(ins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type response struct {
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	Headers    map[string][]string
}

func main() {
	port := 8081
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", requestInfo)
	log.Printf("Started  server on %d", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
