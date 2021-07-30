package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")

	for {
		fmt.Print("Input number-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		err := MakeRequest(text)
		if err != nil {
			panic(err)
		}
	}
}

func MakeRequest(message string) error {
	// create a dialer
	var d net.Dialer

	// server port number
	const port = 8081

	// message - removes '\n' for logging
	fmt.Printf("Sending message: %s; to port: %d\n", message, port)

	// create call context that should close when timeout reached
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// call cancel function to context when we end with our tasks
	defer cancel()

	// connect to server with context
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	// call close to connection when we end with our tasks
	defer conn.Close()

	// send some data to server
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		return err
	}

	// create buffer and read message from server
	getMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Printf("Message recieved: %s\n", getMessage[:len(getMessage)-1])
	return nil
}
