package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			// handle error
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	readBytes := make([]byte, 1024)
	_, err := conn.Read(readBytes)
	if err != nil {
		log.Fatal(err)
	}

	peps := string(readBytes)
	fmt.Print(string(readBytes))
}
