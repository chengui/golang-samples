package simple

import (
	"log"
	"net/http"
)

func Example() {
	addr := ":8000"
	handler := NewHandler()
	http.HandleFunc("/", handler.Serve)
	http.HandleFunc("/chat", handler.HandleChat)
	log.Printf("Start server at %v", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
