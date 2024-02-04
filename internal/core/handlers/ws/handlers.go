package wshandlers

import (
	"log"
	"net/http"
	"random_number_server/internal/core/ports"

	"github.com/gorilla/websocket"
)

type WSHandlers struct {
	generatorRandomNumber ports.GeneratorRandomNumber
	upgrader              *websocket.Upgrader
}

const (
	defaultMSG = "Hello, World!"
)

func (wsh *WSHandlers) Connection(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	number := wsh.generatorRandomNumber.RandNumber()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error: ", err)
			break
		}
		switch string(message) {
		case "get_random_number":
			err = conn.WriteMessage(websocket.TextMessage, []byte(number.String()))
			if err != nil {
				log.Println("WebSocket write error: ", err)
			}
		default:
			err = conn.WriteMessage(websocket.TextMessage, []byte(defaultMSG))
			if err != nil {
				log.Println("WebSocket write error: ", err)
			}
		}

	}
	conn.Close()

}

func NewWSHandlers(generatorRandomNumber ports.GeneratorRandomNumber) *WSHandlers {
	return &WSHandlers{
		generatorRandomNumber: generatorRandomNumber,
		upgrader:              &websocket.Upgrader{},
	}
}
