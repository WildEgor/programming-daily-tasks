package domino_and_tromino_tiling

/**
You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.
Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.
In a tiling, every square must be covered by a tile. Two tilings are different if and only
if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.
@link https://leetcode.com/problems/domino-and-tromino-tiling
*/

/*
From n more than 3 on ward we can calculate it by looping x time and calculate with this formula
(how many unique way to solve i) * (how mane solution to solve n-i) + 2
unique way to solve 1 * result of 3 = 1 * 5 = 5
unique way to solve 2 * result of 2 = 1 * 2 = 2
unique way to solve 3 * result of 1 = 2 * 1 = 2
so the answer for n = 4 is 5 + 2 + 2 + 2 = 11
*/
func numTilings(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	memo := make([]int, n+1)
	memo[1] = 1
	memo[2] = 2

	result := 0
	for i := 3; i <= n; i++ {
		result = 2
		for j := 1; j <= i; j++ {
			unique := 2
			if j < 3 {
				unique = 1
			}
			result += unique * memo[i-j]
		}
		memo[i] = result % 1000000007 // it's from task
	}
	return memo[n]
}

var Solution = numTilings
