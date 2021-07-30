package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// server port number
	const port = 8081

	fmt.Printf("Launching server on port: %d \n\n", port)

	// create listener
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// log error and exit with error code
		log.Fatal(err)
	}

	// endless loop for listening connections
	for {
		// accept new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// serve connection
		go func(c net.Conn) {
			// call close function after serving
			defer c.Close()

			// read string from client(string should end with '\n' symbol)
			message, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				log.Print(err)
				return
			}

			fmt.Printf("Message received: %s \n", message)

			var result []byte

			msg := strings.TrimSpace(message)
			b, err := strconv.Atoi(msg)
			if err != nil {
				c := strings.ToUpper(msg)
				result = []byte(c + "\n")
			} else {
				x := b * 2
				s := strconv.Itoa(x)
				result = []byte(s + "\n")
			}

			// write message back(appends with '\n', because it deleted in previous step)
			_, err = c.Write(result)
			if err != nil {
				log.Print(err)
				return
			}
		}(conn)
	}
}
