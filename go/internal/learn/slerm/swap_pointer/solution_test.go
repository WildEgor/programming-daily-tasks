package swap_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swap(t *testing.T) {
	a := 3
	b := 2
	swap(&a, &b)

	assert.Equal(t, 2, a)
	assert.Equal(t, 3, b)
}
