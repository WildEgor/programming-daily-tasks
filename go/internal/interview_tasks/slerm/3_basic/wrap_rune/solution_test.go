package wrap_rune

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, "cba", reverse("abc"))
}
