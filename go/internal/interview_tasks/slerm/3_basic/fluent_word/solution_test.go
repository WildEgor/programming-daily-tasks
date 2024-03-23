package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = struct {
	input  string
	output string
}{
	"test word test",
	"test",
}

func Test_frequentWord(t *testing.T) {
	assert.Equal(t, tsc.output, frequentWord(tsc.input))
}
