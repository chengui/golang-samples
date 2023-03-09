package engine

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"engine-pool/pool"
)

var engine_pool *pool.EnginePool

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	defer conn.Close()

	sessionID := conn.RemoteAddr().String()
	engine_pool.Add(sessionID)
	log.Printf("Client %s connected.", sessionID)

	for {
		msg_t, msg_s, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ERROR: %v", err)
			break
		}
		log.Printf("RECV %s: %s", sessionID, msg_s)
		engine, ok := engine_pool.Get(sessionID)
		if !ok {
			log.Println("ERROR: engine failure")
			break
		}
		engine.Lock.Lock()
		msg_r, err := engine.Predict(string(msg_s[:]))
		engine.Lock.Unlock()
		if err != nil {
			log.Printf("ERROR: %v", err)
			break
		}
		log.Printf("SENT %s: %s", sessionID, msg_r)
		if err := conn.WriteMessage(msg_t, []byte(msg_r)); err != nil {
			log.Printf("ERROR: %v", err)
			break
		}
	}
}

func ExamplePool() {
	factory := func() *pool.Engine {
		return pool.NewEngine(0)
	}
	engine_pool = pool.NewEnginePool(10, factory)
	if err := engine_pool.Init(); err != nil {
		log.Fatal(err)
	}

	port := ":8000"
	http.HandleFunc("/echo", handler)
	log.Printf("Start websocket at %v", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
