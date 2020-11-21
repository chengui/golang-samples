package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type EchoClient struct {
	Host string
	Port int
}

func NewEchoClient(host string, port int) *EchoClient {
	return &EchoClient{
		Host: host,
		Port: port,
	}
}

func (c *EchoClient) Start() error {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	readline := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input: ")
		input, err := readline.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			return err
		}
		if input == "quit" {
			break
		}
		nw, err := conn.Write([]byte(input))
		if err != nil {
			return err
		}
		log.Printf("Sent %d bytes\n", nw)
		buf := make([]byte, 1024)
		nr, err := conn.Read(buf)
		if err != nil {
			return err
		}
		data := buf[:nr]
		fmt.Printf("Received: %s\n", data)
	}
	return nil
}
