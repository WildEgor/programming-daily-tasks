package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	rooms [][]int
	exp   bool
}{
	{
		rooms: [][]int{{1}, {2}, {3}, {}},
		exp:   true,
	},
}

func Test_canVisitAllRooms_1(t *testing.T) {
	for i, s := range tsc {
		if got := Solution(s.rooms); got != s.exp {
			assert.Equal(t, s.exp, got, "case %d failed", i+1)
		}
	}
}
