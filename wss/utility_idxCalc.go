package wss

import (
	"strconv"
)

/*
	ProcessVolPow
	- goroutine that
		1) gets data from depthStream
		2) calc volumePower = BidImpactDepth / AskImpactDepth
		3) relays volumePower to vpStream
*/
func ProcessVolPow(depthStream chan PartialBookDepthStream, vpStream chan float64) {
	defer close(vpStream)
	for {
		select {
		case m := <-depthStream:
			vp := calcImpactDepth(m.Bids) / calcImpactDepth(m.Asks) * 100
			vpStream <- vp
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

func calcDepthWeight(data [][]string) float64 {
	var (
		depthProd float64
		depthQty  float64
	)
	for _, r := range data {
		sp, _, q := calcDepthInfo(r)
		depthProd += sp
		depthQty += q
	}
	if depthQty != 0 {
		return depthProd / depthQty
	} else {
		return -1
	}

}
