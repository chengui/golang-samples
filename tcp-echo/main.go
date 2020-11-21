package main

import (
	"flag"
	"log"

	"tcp-echo/server"
	"tcp-echo/client"
)

func main() {
	host := flag.String("host", "127.0.0.1", "address ip")
	port := flag.Int("port", 8080, "address port")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("invalid arguments")
	}

	subcommand := flag.Arg(0)
	switch (subcommand) {
	case "server":
		server := server.NewEchoServer(*host, *port)
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	case "client":
		client := client.NewEchoClient(*host, *port)
		err := client.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
