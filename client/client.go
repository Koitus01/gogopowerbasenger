package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		log.Fatal("Connection err: ", err)
	}
	conn.Write([]byte("pipi pupu"))
	readBytes := make([]byte, 1024)
	conn.Read(readBytes)
	print(string(readBytes))
	return
}
