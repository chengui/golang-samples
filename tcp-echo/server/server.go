package server

import (
	"fmt"
	"log"
	"net"
)

type EchoServer struct {
	Host string
	Port int
}

func NewEchoServer(host string, port int) *EchoServer {
	return &EchoServer{
		Host: host,
		Port: port,
	}
}

func (s *EchoServer) Start() error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listen.Close()
	log.Printf("Listen at %s...", addr)

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		go s.Handler(conn)
	}
	return nil
}

func (s *EchoServer) Handler(conn net.Conn) error {
	log.Printf("Connection established.")
	defer conn.Close()
	defer log.Printf("Connection closed.")

	for {
		buf := make([]byte, 1024)
		nr, err := conn.Read(buf)
		if err != nil {
			return err
		}
		data := buf[:nr]
		log.Printf("Received data: %s", data)
		nw, err := conn.Write(data)
		if err != nil {
			return err
		}
		log.Printf("Reply %d bytes", nw)
	}
	return nil
}
