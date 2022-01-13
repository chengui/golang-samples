package chat

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const roomHtml = `<!doctype html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Websocket Chatroom</title>
    <script src="http://code.jquery.com/jquery-1.7.2.min.js"></script>
    <script>
        var ws = new WebSocket("ws://localhost:8000/chat");
        ws.onopen = function(e){
            console.log("onopen");
            console.dir(e);
        };
        ws.onmessage = function(e){
            console.log("onmessage");
            console.dir(e);
            $('#log').append('<p>'+e.data+'<p>');
            $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
        };
        ws.onclose = function(e){
            console.log("onclose");
            console.dir(e);
        };
        ws.onerror = function(e){
            console.log("onerror");
            console.dir(e);
        };
        $(function(){
            $('#msgform').submit(function(){
                ws.send($('#msg').val()+"\n");
                $('#log').append('<p style="color:red;">Me > '+$('#msg').val()+'<p>');
                $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
                $('#msg').val('');
                return false;
            });
        });
    </script>
</head>
<body>
    <div id="log" style="border:1px solid #ccc;height:365px;overflow: auto;margin: 20px 0">
    </div>
    <div>
        <form id="msgform">
            <input type="text" id="msg" style="width:810px" />
        </form>
    </div>
</body>
</html>
`

var (
	roomTmpl = template.Must(template.New("").Parse(roomHtml))
	upgrader = websocket.Upgrader{}
)

type Handler struct {
	Connects map[string]*websocket.Conn
}

func NewHandler() *Handler {
	return &Handler{
		Connects: make(map[string]*websocket.Conn),
	}
}

func (h *Handler) Serve(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	roomTmpl.Execute(w, nil)
}

func (h *Handler) HandleChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return
	}
	defer conn.Close()

	clientID := conn.RemoteAddr().String()
	h.Connects[clientID] = conn
	defer delete(h.Connects, clientID)

	welcome := fmt.Sprintf("Welcome %s join", clientID)
	if err := h.Broadcast("system", websocket.TextMessage, []byte(welcome)); err != nil {
		log.Printf("ERROR: %v", err)
		return
	}

	for {
		typ, dat, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ERROR: %v", err)
			break
		}
		if err := h.Broadcast(clientID, typ, dat); err != nil {
			log.Printf("ERROR: %v", err)
			break
		}
	}
}

func (h *Handler) Broadcast(clientID string, typ int, mesg []byte) error {
	merge := []byte(fmt.Sprintf("%s > %s", clientID, mesg))
	for key, conn := range h.Connects {
		if key == clientID {
			continue
		}
		if err := conn.WriteMessage(typ, merge); err != nil {
			log.Printf("ERROR: %v", err)
			return err
		}
	}
	return nil
}
