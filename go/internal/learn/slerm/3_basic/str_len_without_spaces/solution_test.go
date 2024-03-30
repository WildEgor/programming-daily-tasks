package str_len_without_spaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeSpaces(t *testing.T) {
	assert.Equal(t, 4, stringLengthWithoutSpaces("a b c \t"))
}
