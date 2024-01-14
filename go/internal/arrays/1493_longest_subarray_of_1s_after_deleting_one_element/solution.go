package template

/**
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.
@link https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
*/

func longestSubarray(nums []int) int {
	var currentMax, previousMax, result int

	for _, num := range nums {
		if num == 1 {
			currentMax++
		} else {
			result = max(result, previousMax+currentMax)
			previousMax = currentMax
			currentMax = 0
		}
	}

	if currentMax == len(nums) {
		return currentMax - 1
	}

	return max(result, previousMax+currentMax)
}

func longestSubarrayAlt(nums []int) int {
	left, ans, zeroIdx := 0, 0, -1
	for i, v := range nums {
		if v == 0 {
			left = zeroIdx + 1
			zeroIdx = i
		}
		if i-left > ans {
			ans = i - left
		}
	}
	return ans
}

var Solution = longestSubarrayAlt
