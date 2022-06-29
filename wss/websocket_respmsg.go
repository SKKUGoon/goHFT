package wss

/*
	Aggregate Trade Streams
	- Aggregate Trade Stream - push trade information that is aggregated for fills
	- with same price and taking side
	- <symbol>@aggTrade
*/
type AggTradeStream struct {
	EventType    string `json:"e"`
	EventTime    int    `json:"E"`
	Symbol       string `json:"s"`
	AggTradeID   int    `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeId int    `json:"f"`
	LastTradeId  int    `json:"l"`
	TradeTime    int    `json:"T"`
	MarketMaker  bool   `json:"m"`
}

/*
	Mark Price Stream
	- Mark price and funding rate for a single symbol. 3 seconds or every second
	- <symbol>@markPrice
	- <symbol>@markPrice@1s 3s
*/
type MarkPriceStream struct {
	EventType       string `json:"e"`
	EventTime       int    `json:"E"`
	Symbol          string `json:"s"`
	MarkPrice       string `json:"p"`
	IndexPrice      string `json:"i"`
	EstSettlePrice  string `json:"P"`
	FundingRate     string `json:"r"`
	NextFundingTime int    `json:"T"`
}

/*
	Kline/Candlestick Streams
	- push updates to the current klines/candlestick every 250 milliseconds (if existing)
*/
type KlineCandleStickStream struct {
	EventType string `json:"e"`
	EventTime int    `json:"E"`
	Symbol    string `json:"s"`
	Kline     kline  `json:"k"`
}

type kline struct {
	KlineStartTime           int    `json:"t"`
	KlineCloseTime           int    `json:"T"`
	Symbol                   string `json:"s"`
	Interval                 string `json:"i"`
	FirstTradeID             int    `json:"f"`
	LastTradeId              int    `json:"L"`
	OpenPrice                string `json:"o"`
	HighPrice                string `json:"h"`
	LowPrice                 string `json:"l"`
	BaseAssetVolume          string `json:"v"`
	NumberTrade              int    `json:"n"`
	KlineClose               bool   `json:"x"`
	QuoteAssetVolume         string `json:"q"`
	TakerBuyBaseAssetVolume  string `json:"V"`
	TakerBuyQuoteAssetVolume string `json:"Q"`
}

/*
	Partial Book Depth Streams
	- Stream Names
	- <symbol>@depth<levels>  levels in {5, 10, 20}
	- <symbol>@depth<levels>@<speeds> speeds in 250ms, 500ms or 100ms
*/
type PartialBookDepthStream struct {
	EventType               string     `json:"e"`
	EventTime               int        `json:"E"`
	TransactionTime         int        `json:"T"`
	Symbol                  string     `json:"s"`
	FirstUpdateId           int        `json:"U"`
	FinalUpdateId           int        `json:"u"`
	FinalUpdateIdLastStream int        `json:"pu"`
	Bids                    [][]string `json:"b"`
	Asks                    [][]string `json:"a"`
}
