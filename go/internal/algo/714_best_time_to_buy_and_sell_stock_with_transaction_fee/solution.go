package best_time_to_buy_and_sell_stock_with_transaction_fee

import "math"

/**
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
Note:
You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
The transaction fee is only charged once for each stock purchase and sale.
@link https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
*/

// TODO: need more explanation
func maxProfit(prices []int, fee int) int {
	// edge case
	if len(prices) == 0 {
		return 0
	}

	withShare := float64(prices[0] * -1)
	withNo := float64(0)
	for i := 0; i < len(prices); i++ {
		withNo = math.Max(withNo, withShare+float64(prices[i]-fee))
		withShare = math.Max(withShare, withNo-float64(prices[i]))
	}

	return int(withNo)
}

var Solution = maxProfit
