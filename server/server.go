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
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024) // буфер для чтения клиентских данных
	for {
		_, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		conn.Write([]byte(fmt.Sprintln("You send:", string(buf)))) // пишем в сокет
	}
}
