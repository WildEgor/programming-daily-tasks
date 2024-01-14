package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  []int
}{
	{
		1,
		[]int{100, 80, 60, 70, 60, 75, 85},
		[]int{1, 1, 1, 2, 1, 4, 6},
	},
}

func Test_StockSpanner_1(t *testing.T) {

	spanner := Constructor()

	for _, c := range cases {
		for i, v := range c.input {
			assert.Equalf(t, c.want[i], spanner.NextAlt(v), "case %d, input %v", c.id, c.input)
		}
	}
}
