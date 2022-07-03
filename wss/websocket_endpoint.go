package wss

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Pairs[T, U any] struct {
	First  T
	Second U
}

func genTestWxURL(endpoint string) string {
	return TestNetBaseWssSingle + "/" + endpoint
}

func genMainWxURL(endpoint string) string {
	return MainNetBaseWssSingle + "/" + endpoint
}

/*
	getWxURL
	- URL, JSON container
*/
func getWxURL(targetSymbol, info string, isTest bool) string {
	var endpoint string
	switch info {
	case "depth":
		endpoint = fmt.Sprintf("%s@depth20@100ms", targetSymbol)
	case "aggTrade":
		endpoint = fmt.Sprintf("%s@aggTrade", targetSymbol)
	default:
		return ""
	}
	if isTest {
		return genTestWxURL(endpoint)
	}
	return genMainWxURL(endpoint)
}

func GetDepthWx(symbol string) (*websocket.Conn, PartialBookDepthStream, error) {
	u := getWxURL(symbol, "depth", IsTest)
	fmt.Println(u)

	ct := PartialBookDepthStream{}
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	return conn, ct, err
}

func GetAggTradeWx(symbol string) (*websocket.Conn, AggTradeStream, error) {
	u := getWxURL(symbol, "aggTrade", IsTest)
	fmt.Println(u)

	ct := AggTradeStream{}
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	return conn, ct, err
}
