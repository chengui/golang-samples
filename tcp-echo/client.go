package main

import (
	"fmt"
	"flag"
	"log"
	"net"
)

func main() {
	host := flag.String("host", "127.0.0.1", "server ip")
	port := flag.Int("port", 8080, "server port")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Panic(err)
	}

	for {
		fmt.Print("Please input: ")
		var input string
		fmt.Scanln(&input)
		log.Println("debug: ", input)
		if input == "quit" {
			break
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Panic(err)
		}
		buf := make([]byte, 1024)
		nr, err := conn.Read(buf)
		if err != nil {
			log.Panic(err)
		}
		data := buf[:nr]
		fmt.Printf("Received: %s\n", data)
	}
}
