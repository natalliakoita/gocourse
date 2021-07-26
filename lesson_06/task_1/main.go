package main

import (
	"encoding/json"
	"net/http"
)

func requestInfo(w http.ResponseWriter, req *http.Request) {
	var ins response

	ins.Host = req.Host
	v := req.Header["Accept"]
	v1 := req.Header["User-Agent"]
	ins.Headers.UserAgent = v1
	ins.Headers.Accept = v
	ins.RequestURI = req.RequestURI
	ins.UserAgent = v1[0]

	js, err := json.Marshal(ins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type requestHeaders struct {
	Accept    []string `json:"Assept"`
	UserAgent []string `json:"User-Agent"`
}

type response struct {
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	Headers    requestHeaders
}

func main() {
	http.HandleFunc("/headers", requestInfo)

	http.ListenAndServe(":8081", nil)
}
