package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tcs = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{2, 1, 3},
		output: []int{1, 2, 3},
	},
}

func Test_bubble(t *testing.T) {
	for _, tc := range tcs {
		assert.Equal(t, tc.output, bubbleSort(tc.input))
	}
}

func Test_insert(t *testing.T) {
	for _, tc := range tcs {
		insertSortAlt(tc.input)

		assert.Equal(t, tc.output, tc.input)
	}
}

func Test_insertTime(t *testing.T) {
	defer TimeTrack(time.Now(), "insertSortAlt")

	insertSortAlt([]int{2, 1, 3, 4, 2, 5, 4, 32, 23, 1})

	//defer TimeTrack(time.Now(), "insertSort")
	//
	//insertSort([]int{2, 1, 3})

	assert.Equal(t, true, true)
}
