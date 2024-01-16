package evaluate_division

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	eq [][]string
	v  []float64
	q  [][]string
	r  []float64
}{
	{
		[][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}},
		[]float64{6, 0.5, -1},
	},
}

func Test_calcEquation_1(t *testing.T) {
	for _, tc := range tsc {
		assert.Equal(t, tc.r, Solution(tc.eq, tc.v, tc.q))
	}
}
