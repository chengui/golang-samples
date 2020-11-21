package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	host := flag.String("host", "127.0.0.1", "address to bind")
	port := flag.Int("port", 8080, "port to listen")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Printf("Listen at %s:%d...", *host, *port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	log.Printf("Connection established.")
	defer conn.Close()
	defer log.Printf("Connection closed.")

	for {
		buf := make([]byte, 1024)
		nr, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:nr]
		log.Printf("Received data: %s", data)
		nw, err := conn.Write(data)
		if err != nil {
			return
		}
		log.Printf("Reply %d bytes", nw)
	}
}
