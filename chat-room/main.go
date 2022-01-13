package main

import (
	"log"
	"net/http"

	"chat-room/chat"
)

func main() {
	addr := ":8000"
	handler := chat.NewHandler()
	http.HandleFunc("/", handler.Serve)
	http.HandleFunc("/chat", handler.HandleChat)
	log.Printf("Start server at %v", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
