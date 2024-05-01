package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		log.Fatal("Connection err: ", err)
	}

	//TODO: not properly send input to server, previous input append to new string. Maybe scanLn??
	go copyTo(os.Stdout, conn)
	copyTo(conn, os.Stdin)
	//conn.Write([]byte("pipi pupu"))
	//readBytes := make([]byte, 1024)
	//_, err = conn.Read(readBytes)
	//if err != nil {
	//	log.Fatal("Read err: ", err)
	//}
	//fmt.Print(string(readBytes))
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
