package wss

import (
	"fmt"
	"log"
	"math"
)

type Pairs[T, U any] struct {
	First  T
	Second U
}

func genSymbol(symbols, info []string) []string {
	minLen := int(math.Min(float64(len(symbols)), float64(len(info))))
	s := make([]string, minLen)
	for i := 0; i < minLen; i++ {
		s[i] = symbols[i] + "@" + info[i]
	}
	return s
}

func genTestWxURL(endpoint string) string {
	return TESTNET_BASE_WSS_SINGLE + "/" + endpoint
}

func genMainWxURL(endpoint string) string {
	return MAINNET_BASE_WSS_SINGLE + "/" + endpoint
}

func GetWxStream(targetSymbol, info string, isTest bool) string {
	endpoint := fmt.Sprintf("%s@%s", targetSymbol, info)
	log.Println(endpoint)
	if isTest {
		return genTestWxURL(endpoint)
	}
	return genMainWxURL(endpoint)
}
