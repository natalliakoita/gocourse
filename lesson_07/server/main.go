package main

import (
	"log"
	"net/http"

	handlers "github.com/natalliakoita/gocourse/lesson_07/server/handle"
)

func main() {
	addr := "localhost:8081"
	endpoint := "/rpc"
	http.HandleFunc(endpoint, handlers.Handle)
	log.Printf("Started JSON RPC server on %s%s", addr, endpoint)
	log.Fatal(http.ListenAndServe(addr, nil))
}
