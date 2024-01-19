package remove_element

import (
	"cmp"
	"slices"
)

/**
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
@link https://leetcode.com/problems/remove-element
*/

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			k++
		} else {
			nums[i] = 0
		}
	}

	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return k
}

func removeElementAlt(nums []int, val int) int {
	k := 0

	for i, num := range nums {
		if num != val {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	return k
}

var Solution = removeElementAlt
