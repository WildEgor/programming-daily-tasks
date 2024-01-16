package best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	prices []int
	fee    int
	result int
}{
	{
		[]int{1, 3, 2, 8, 4, 9},
		2,
		8,
	},
	{
		[]int{1, 3, 7, 5, 10, 3},
		3,
		6,
	},
}

func Test_maxProfit_1(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.result, Solution(tc.prices, tc.fee))
	}
}
