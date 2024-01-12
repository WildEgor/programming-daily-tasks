package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteMiddle_1(t *testing.T) {
	heads := []int{1, 3, 4, 7, 1, 2, 6}
	linked := NewListNode()

	for _, val := range heads {
		linked.Insert(val)
	}

	head := Solution(linked)
	result := head.List()

	assert.ElementsMatch(t, []int{1, 3, 4, 1, 2, 6}, result)
}
