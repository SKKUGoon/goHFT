package wss

import (
	"os"
	"time"
)

/*
	OLD WEBSOCKET DOC
	* websocket base point 1
		baseUrl: wss://stream.binance.com:9443/ws
			1 raw stream example :
				- wss://stream.binance.com:9443/ws/bnbusdt@aggTrade
		baseUrl: wss://stream.binance.com:9443/stream?streams=
			multi raw stream example :
				- wss://stream.binance.com:9443/stream?streams=bnbusdt@aggTrade

	* connection
		single connection is valid for 24 hours. disconnect at 24 hour mark
		ping frame every 5 minutes - should pong within 15 min
		10 message per second
		single connection - 200 streams
*/

const (
	IsTest               = false
	MainNetBaseWssSingle = "wss://stream.binance.com:9443/ws"
	TestNetBaseWssSingle = "wss://testnet.binance.vision/ws"
)

const (
	ImpactNotional = 50_000
	WebsocketTO    = time.Second * 60
)

var (
	// Receive Messages
	BookDepthChan = make(chan PartialBookDepthStream)
	AggTradeChan  = make(chan AggTradeStream)

	// Custom Index
	PremiumChan  = make(chan float64)
	VolPowerChan = make(chan float64)

	// Others
	Interruption = make(chan os.Signal, 1)
	Ticking      = time.NewTicker(time.Second)
)
