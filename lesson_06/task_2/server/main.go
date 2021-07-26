package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Member struct {
	Value string `json:"value"`
}

func HandlerX2(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dat Member

	if err := json.Unmarshal(reqBody, &dat); err != nil {
		panic(err)
	}
	var result Member

	b, err := strconv.Atoi(dat.Value)
	if err != nil {
		c := strings.ToUpper(dat.Value)
		result.Value = c
	} else {
		x := b * 2
		s := strconv.Itoa(x)
		result.Value = s
	}

	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/x2", HandlerX2)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
