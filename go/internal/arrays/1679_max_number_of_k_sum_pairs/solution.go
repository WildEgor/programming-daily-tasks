package arrays_1679_max_number_of_k_sum_pairs

/**
You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
Return the maximum number of operations you can perform on the array.
@link https://leetcode.com/problems/max-number-of-k-sum-pairs/
*/

func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	var opsCounter int

	for _, v := range nums {
		if m[k-v] > 0 {
			m[k-v]--
			opsCounter++
		} else {
			m[v]++
		}
	}

	return opsCounter
}

var Solution = maxOperations
