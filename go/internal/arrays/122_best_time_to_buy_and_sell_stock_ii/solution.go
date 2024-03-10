package best_time_to_buy_and_sell_stock_ii

import "math"

func maxProfit(prices []int) int {
	hold, notHold := math.MinInt64, 0

	for _, price := range prices {
		prevHold, prevNotHold := hold, notHold

		hold = max(prevHold, prevNotHold-price)
		notHold = max(prevNotHold, prevHold+price)
	}

	return notHold
}

var Solution = maxProfit
