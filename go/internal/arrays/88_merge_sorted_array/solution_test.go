package merge_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	want  []int
}{
	{
		[]int{1, 2, 3, 0, 0, 0},
		3,
		[]int{2, 5, 6},
		3,
		[]int{1, 2, 2, 3, 5, 6},
	},
}

func Test_example_1(t *testing.T) {
	for _, tc := range tcs {
		Solution(tc.nums1, tc.m, tc.nums2, tc.n)
		assert.Equal(t, tc.want, tc.nums1)
	}
}
