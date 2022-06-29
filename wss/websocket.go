package wss

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func ReceiveHandler(conn *websocket.Conn, Msg chan interface{}, container any) {
	defer close(Msg)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error in receive: ", err)
		}
		err = json.Unmarshal(msg, container)
		if err != nil {
			log.Println("Error in unmarshal: ", err)
		}
		Msg <- container
	}
}
