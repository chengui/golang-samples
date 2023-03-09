package echo

import (
	"log"
)

func Example() {
	host, port := "localhost", 8000
	server := NewEchoServer(host, port)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

	client := NewEchoClient(host, port)
	err = client.Start()
	if err != nil {
		log.Fatal(err)
	}
}
