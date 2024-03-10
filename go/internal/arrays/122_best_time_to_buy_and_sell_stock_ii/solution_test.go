package best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	input  []int
	output int
}{
	{
		input:  []int{7, 1, 5, 3, 6, 4},
		output: 7,
	},
	{
		input:  []int{1, 2, 3, 4, 5},
		output: 4,
	},
	{
		input:  []int{7, 6, 4, 3, 1},
		output: 0,
	},
}

func Test_solution(t *testing.T) {
	for _, test := range tsc {
		assert.Equal(t, test.output, maxProfit(test.input))
	}
}
