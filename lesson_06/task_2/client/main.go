package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Member struct {
	Value string `json:"value"`
}

func MakeRequest(serverURL string, v string) error {
	z := Member{Value: v}
	json_data, err := json.Marshal(z)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		serverURL,
		"application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dat Member

	if err := json.Unmarshal(respBody, &dat); err != nil {
		return err
	}

	fmt.Println(dat.Value)

	return nil
}

func main() {
	const serverURL = "http://localhost:8081/x2"
	const port = 8081

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		err := MakeRequest(serverURL, text)
		if err != nil {
			panic(err)
		}
	}
}
