package to_freq_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = struct {
	input  []string
	output map[string]int
}{
	[]string{"a", "b", "a"},
	map[string]int{
		"a": 2,
		"b": 1,
	},
}

func Test_toFrequencyMap(t *testing.T) {
	assert.Equal(t, tsc.output, toFrequencyMap(tsc.input))
}
