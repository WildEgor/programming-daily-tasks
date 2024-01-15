package maximum_subsequence_score

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tsc = []struct {
	nums1 []int
	nums2 []int
	k     int
	ans   int64
}{
	//{
	//	nums1: []int{1, 3, 3, 2},
	//	nums2: []int{2, 1, 3, 4},
	//	k:     3,
	//	ans:   12,
	//},
	{
		nums1: []int{2, 1, 14, 12},
		nums2: []int{11, 7, 13, 6},
		k:     3,
		ans:   168,
	},
}

func Test_maxScore_1(t *testing.T) {
	for _, tc := range tsc {
		result := Solution(tc.nums1, tc.nums2, tc.k)
		assert.Equal(t, tc.ans, result)
	}
}
