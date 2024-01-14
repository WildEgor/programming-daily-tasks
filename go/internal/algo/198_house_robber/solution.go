package house_robber

/**
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
@link https://leetcode.com/problems/house-robber
*/

/*
*
So, you should not just rob the odd-indexed or even-indexed houses,
but instead, for each house, you decide to rob it or not based on whether
it maximizes the total amount robbed so far.
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		// compare the current house with the previous house
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}

	return dp[len(nums)]
}

var Solution = rob
