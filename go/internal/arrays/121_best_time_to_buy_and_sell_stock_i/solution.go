package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	result := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > result {
				result = profit
			}
		}
	}

	return result
}

func maxProfitAlt(prices []int) int {
	result := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - minPrice

		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if diff > result {
			result = diff
		}
	}

	return result
}

var Solution = maxProfitAlt
