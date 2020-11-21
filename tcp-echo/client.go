package main

import (
	"bufio"
	"fmt"
	"flag"
	"log"
	"net"
	"os"
	"strings"
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

	readline := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input: ")
		input, err := readline.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			return
		}
		if input == "quit" {
			break
		}
		nw, err := conn.Write([]byte(input))
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Sent %d bytes\n", nw)
		buf := make([]byte, 1024)
		nr, err := conn.Read(buf)
		if err != nil {
			log.Panic(err)
		}
		data := buf[:nr]
		fmt.Printf("Received: %s\n", data)
	}
}
