package wss

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func DepthHandler(conn *websocket.Conn, depthStream chan PartialBookDepthStream, container *PartialBookDepthStream) {
	defer close(depthStream)
	for {
		err := conn.ReadJSON(&container)
		if err != nil {
			log.Println("Error in unmarshal : depth :", err)
		}
		depthStream <- *container
	}
}

func AggTradeHandler(conn *websocket.Conn, tradeStream chan AggTradeStream, container *AggTradeStream) {
	defer close(tradeStream)
	for {
		err := conn.ReadJSON(&container)
		if err != nil {
			log.Println("Error in unmarshal : trade :", err)
		}
		fmt.Println(container)
	}
}

func KeepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
