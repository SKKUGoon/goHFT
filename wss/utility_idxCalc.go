package wss

import (
	"log"
	"strconv"
)

/*
	ProcessPremium
	- goroutine that
		1) gets data from depthStream
		2) calc premium = BidImpactDepth / AskImpactDepth
		3) relays volumePower to vpStream
*/
func ProcessPremium(depthStream chan PartialBookDepthStream, premiumStream chan float64) {
	defer close(premiumStream)
	for {
		select {
		case m := <-depthStream:
			premium := calcImpactDepth(m.Bids) / calcImpactDepth(m.Asks) * 100
			premiumStream <- premium
		}
	}
}

/*
	ProcessVolPower
	- goroutine that
		1) gets data from tradeStream
		2) calc
*/
func ProcessVolPower(tradeStream chan AggTradeStream, resultStream chan float64) {
	defer close(tradeStream)

	var buyTick float64
	var sellTick float64
	var trend = map[string]int{
		"u": 0,
		"d": 0,
		"n": 0,
	}
	for {
		select {
		case t := <-tradeStream:
			q, _ := strconv.ParseFloat(t.Quantity, 64)
			if t.BuyerMaker {
				sellTick += q
			} else {
				buyTick += q
			}
		case <-Ticking.C:
			// update trend info by 500 milliseconds
			switch {
			case (buyTick == 0) && (sellTick == 0):
				trend["n"] += 1
			case (buyTick / sellTick) > 5:
				trend["u"] += 1
			case (buyTick / sellTick) <= 0.2:
				trend["d"] += 1
			default:
				trend["n"] += 1
			}
			//resultStream <- buyTick / sellTick
			//fmt.Println("tick", buyTick/sellTick)
			buyTick, sellTick = 0, 0
		case <-Gathering.C:
			// TODO: develop logic that identifies upward downward trend
			t := calcTrend(trend)
			log.Println(t)
			trend["u"], trend["d"], trend["n"] = 0, 0, 0

		}
	}
}

/*
	calcImpactDepth
	- impact depth: average price to settle ImpactNotional(constant)
*/
func calcImpactDepth(data [][]string) float64 {
	var (
		depthProd float64
		impact    float64
		depth     = 0
	)
	for dp, r := range data {
		sp, p, _ := calcDepthInfo(r)
		depthProd += sp
		impact += p
		depth = dp
		if depthProd > ImpactNotional {
			break
		}

	}
	return impact / float64(depth+1)
}

/*
	calcDepthInfo
	- most information comes in string for accuracy.
	- change string - depth price and quantity information into float64
*/
func calcDepthInfo(frag []string) (float64, float64, float64) {
	prc, qty := frag[0], frag[1]
	prcF, _ := strconv.ParseFloat(prc, 64)
	qtyF, _ := strconv.ParseFloat(qty, 64)
	return prcF * qtyF, prcF, qtyF
}

func calcTrend(data map[string]int) string {
	var maxKey string
	var maxi = -1
	for k, v := range data {
		if v > maxi {
			maxi = v
			maxKey = k
		}
	}
	return maxKey
}
