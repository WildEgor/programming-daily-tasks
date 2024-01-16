package koko_eating_bananas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	piles []int
	h     int
	r     int
}{
	{
		piles: []int{3, 6, 7, 11},
		h:     8,
		r:     4,
	},
	{
		piles: []int{30, 11, 23, 4, 20},
		h:     5,
		r:     30,
	},
	{
		piles: []int{30, 11, 23, 4, 20},
		h:     6,
		r:     23,
	},
}

func Test_minEatingSpeed_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.r, Solution(tc.piles, tc.h))
	}
}
