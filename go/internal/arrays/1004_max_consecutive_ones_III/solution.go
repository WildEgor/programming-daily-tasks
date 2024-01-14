package max_consecutive_ones_III

/**
  Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
  @link https://leetcode.com/problems/max-consecutive-ones-iii/
*/

func longestOnes(nums []int, k int) int {
	var m, left, right, count int

	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}

		if count > k {
			if nums[left] == 0 {
				count--
			}

			left++
		}

		m = max(m, right-left+1)

		right++
	}

	return m
}

var Solution = longestOnes
