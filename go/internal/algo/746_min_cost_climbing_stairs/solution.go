package template

/**
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
@link https://leetcode.com/problems/min-cost-climbing-stairs/
*/

func minCostClimbingStairs(cost []int) int {
	result := make([]int, len(cost)+1)
	result[0] = 0
	result[1] = 0

	for i := 2; i <= len(cost); i++ {
		result[i] = min(result[i-1]+cost[i-1], result[i-2]+cost[i-2])
	}

	return result[len(cost)]
}

var Solution = minCostClimbingStairs
