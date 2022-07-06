# 0.0.1
<p>
initial commit
</p>

[Add]
- ./wss
  - c.go
    - constants
  - endpoint.go
    - func genTestWxURL
    - func genMainWxURL
    - func GetWxStream
  - websocket.go
    - func ReceiveHandler
  - websocket_respmsg.go
    - struct AggTradeStream
    - struct MarkPriceStream
    - struct KlineCandleStickStream
    - struct kline
    - struct PartialBookDepthStream
- main.go
  - func main

[Change]

[Fix]

[Remove]


# 0.1.0
<p>
Major Overhaul + Newly Activated Functions
</p>

[Add]
- ./wss
  - c.go
    - const IsTest
    - const ImpactNotional
    - const WebSocketTO - TO for TimeOut
    - var BookDepthChan
    - var AggTradeChan
    - var VolumePowerChan
    - var Interruption
    - var Ticking
  - utility_idxCalc.go
    - func ProcessVolPow
    - func calcImpactDepth
    - func calcDepthInfo
    - func calcDepthWeight
  - websocket_endpoint.go
    - func getWxURL
    - func GetDepthWx
    - func GetAggTradeWx
  - websocket_service.go
    - func searchSymbol 
    - func WxServe (main)
    - WxSAggServe
    - WxDepthServe
  - websocket_handler.go
    - func DepthHandler
    - func AggTradeHandler  - still json writer
    - KeepAlive - ping pong 

[Change]
- ./wss
  - websocket_endpoint.go - changed from endpoint.go
  - websocket_respmsg.go
    - struct AggTradeStream - add Placeholder. if there's no placeholder, buyermaker field will always parse 'true'.
    - struct PartialBookDepthStream - change `json structure'. The official doc was wrong.(not updated)

[Fix]

[Remove]


# 0.1.1
[Add]
- ./wss
  - utility_idxCalc.go
    - func ProcessVolPower

[Change]
- ./wss
  - c.go
    - chan PremiumChan - renamed from VolPowerChan
    - chan VolPowerChan
  - utility_idxCalc.go
    - func ProcessPremium - rename from ProcessVolPow
    - func ProcessVolPower
  - websocket_service.go
    - func WxAggServe - add ProcessVolPower goroutine

ProcessVolPow 

[Fix]

[Remove]

# 0.1.2
<p>
new go routine. Calculate Volume power every 500 milliseconds
</p>

[Add]

[Change]
- ./wss
  - utility_idxCalc.go
    - func ProcessVolPower 
      - if get tradeStream { add up } if get tick { push VolPow to chan }. 
  - websocket_service.go
    - func WxServe
      - pull premium from PremiumChan
      - pull volPower from VolPowerChan
    - func WxAggServe
      - start another go routine ProcessVolPower

[Fix]

[Remove]
- ./wss
  - c.go
    - Interruption channel

# 0.1.3
<p>
1 TODO, 1 FIXME added. tech is here. need strategy
</p>

[Add]

[Change]
- ./wss
  - utility_idxCalc.go
    - case Ticking.C: add instance count to find trend
    - case Gathering.C: trend finding? it's little wierd. 
    - NEED TO FIND TREND BY ANALYZING ADDING SEQUENCE. U U D D -> should be D or U?

[Fix]

[Remove]



