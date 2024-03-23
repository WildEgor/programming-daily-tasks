package remove_spaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeSpaces(t *testing.T) {
	assert.Equal(t, "abc", removeSpaces("a b c"))
}
