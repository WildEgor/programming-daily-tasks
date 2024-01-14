package template

/**
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
ans[i] is the number of 1's in the binary representation of i.
@link https://leetcode.com/problems/counting-bits/
*/

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		/**
		i>>1 is the same as i/2
		i&1 is the same as i%2
		*/
		result[i] = result[i>>1] + i&1
	}

	return result
}

var Solution = countBits
