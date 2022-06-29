package wss

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
	MAINNET_BASE_WSS_SINGLE = "wss://stream.binance.com:9443/ws"
	TESTNET_BASE_WSS_SINGLE = "wss://testnet.binance.vision/ws"
	MAINNET_BASE_WSS_MULTI  = "wss://stream.binance.com:9443/stream?streams="
	TESTNET_BASE_WSS_MULTI  = "wss://testnet.binance.vision/stream?streams="
)
