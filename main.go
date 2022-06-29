package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goHFT/wss"
	"log"
	"os"
	"time"
)

var IS_TEST = true

func main() {
	var testTicker = 0
	socketUrl := wss.GetWxStream("btcusdt", "aggTrade", IS_TEST)
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket", err)
	}
	defer conn.Close()

	var agg = wss.AggTradeStream{}
	var msgChannel = make(chan interface{})
	go wss.ReceiveHandler(conn, msgChannel, &agg)

	interrupt := make(chan os.Signal, 1)
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case m := <-msgChannel:
			fmt.Println("h container", m)
		case t := <-ticker.C:
			fmt.Println("ticker", t, "testTicker", testTicker)
		case <-interrupt:
			log.Println("interrupt")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close: ", err)
				return
			}
			return
		}
	}
}
