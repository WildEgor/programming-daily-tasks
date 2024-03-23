package freq_rune

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = struct {
	input  string
	output rune
}{
	"test",
	't',
}

func Test_frequentRune(t *testing.T) {
	assert.Equal(t, tsc.output, frequentRune(tsc.input))
}
