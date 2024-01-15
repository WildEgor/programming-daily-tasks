package smallest_number_in_infinite_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SmallestInfiniteSet_1(t *testing.T) {
	result := make([]int, 0)
	sis := Constructor()
	sis.AddBack(2)                             // 2 is already in the set, so no change is made.
	result = append(result, sis.PopSmallest()) // return 1, since 1 is the smallest number, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 2, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 3, and remove it from the set.
	sis.AddBack(1)                             // 1 is added back to the set.
	result = append(result, sis.PopSmallest()) // return 1, since 1 was added back to the set and
	// is the smallest number, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 4, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 5, and remove it from the set.

	assert.Equal(t, []int{1, 2, 3, 1, 4, 5}, result)
}

func Test_SmallestInfiniteSet_2(t *testing.T) {
	result := make([]int, 0)
	sis := NewSmallestInfiniteSetAlt()
	sis.AddBack(2)                             // 2 is already in the set, so no change is made.
	result = append(result, sis.PopSmallest()) // return 1, since 1 is the smallest number, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 2, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 3, and remove it from the set.
	sis.AddBack(1)                             // 1 is added back to the set.
	result = append(result, sis.PopSmallest()) // return 1, since 1 was added back to the set and
	// is the smallest number, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 4, and remove it from the set.
	result = append(result, sis.PopSmallest()) // return 5, and remove it from the set.

	assert.Equal(t, []int{1, 2, 3, 1, 4, 5}, result)
}
