package server

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(msg))
		cmd := exec.Command("sh", "-c", "cd /root &&"+string(msg))
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		if err := ws.WriteMessage(websocket.TextMessage, output); err != nil {
			log.Println(err)
			break
		}
	}
}
