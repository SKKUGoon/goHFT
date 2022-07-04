package wss

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
)

type AggTradeWS func(*websocket.Conn, chan AggTradeStream, *AggTradeStream)
type DepthWS func(*websocket.Conn, chan PartialBookDepthStream, *PartialBookDepthStream)

func searchSymbol() string {
	tic := os.Args
	if len(tic) == 1 {
		log.Panicln("not enough arguments")
	}
	return tic[1]
}

/*
	WxServe
	- main actuator
*/
func WxServe() {
	go WxAggServe(AggTradeHandler)
	go WxDepthServe(DepthHandler)

	for {
		select {
		case volPower := <-PremiumChan:
			log.Println(volPower)
		case <-Ticking.C:
		}
	}
}

/*
	WxAggServe
	- Process AggTrade endpoint information
	- spawn 3 goroutine: KeepAlive & Callback (JSON formatter) func & ProcessVolPower
*/
func WxAggServe(io AggTradeWS) {
	conn, contain, err := GetAggTradeWx(searchSymbol())
	if err != nil {
		log.Panicln("Error connecting to AggTrade", err)
	}

	go KeepAlive(conn, WebsocketTO)
	go io(conn, AggTradeChan, &contain)
	go ProcessVolPower(AggTradeChan)
}

/*
	WxDepthServe
	- Process depth<> endpoint information
	- spawn 3 goroutine: KeepAlive & Callback (JSON formatter) & VolumePower Index calc
*/
func WxDepthServe(io DepthWS) {
	conn, contain, err := GetDepthWx(searchSymbol())
	if err != nil {
		log.Panicln("Error connecting to PartialDepth", err)
	}

	go KeepAlive(conn, WebsocketTO)
	go io(conn, BookDepthChan, &contain)
	go ProcessPremium(BookDepthChan, PremiumChan)
}
