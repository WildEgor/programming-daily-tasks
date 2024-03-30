package map_intersect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = struct {
	input1 map[int]struct{}
	input2 map[int]struct{}
	output []int
}{
	map[int]struct{}{
		1: {},
		2: {},
		3: {},
	},
	map[int]struct{}{
		1: {},
		2: {},
		3: {},
	},
	[]int{1, 2, 3},
}

func Test_mapKeyIntersect(t *testing.T) {
	assert.Equal(t, tsc.output, mapKeyIntersect(tsc.input1, tsc.input2))
}
