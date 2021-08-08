package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Method string         `json:"method"`
	Params registerParams `json:"params"`
	Id     int64          `json:"id"`
}

type registerParams struct {
	Name string `json:"name"`
}

func main() {
	var req Request

	fmt.Println("Enter the method (register or list):")
	fmt.Fscan(os.Stdin, &req.Method)
	if req.Method == "register" {
		fmt.Println("Enter the name:")
		fmt.Fscan(os.Stdin, &req.Params.Name)
	}

	// number generator
	now := time.Now()
	sec := now.Unix()
	req.Id = sec

	json_data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	// connect to server
	resp, err := http.Post(
		"http://localhost:8081/rpc",
		"application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		log.Fatal(err)
	}

	// get resp body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}
