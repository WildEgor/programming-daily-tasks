package rotate_array

/**
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
@link https://leetcode.com/problems/rotate-array
*/

//func rotate(nums []int, k int) {
//	copy(nums, copyToK(nums, k))
//}
//
//func copyToK(nums []int, k int) []int {
//	var ans = make([]int, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		idx := (i + k) % len(nums) // the modulo operator gives us the overflow remainder and that delivers our idx
//		ans[idx] = nums[i]
//	}
//
//	return ans
//}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1) // reverse all array
	reverse(nums, 0, k-1) // reverse left part
	reverse(nums, k, n-1) // reverse back right part
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

var Solution = rotate
