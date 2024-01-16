package template

import "slices"

/**
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:
Insert a character
Delete a character
Replace a character
@link https://leetcode.com/problems/edit-distance
*/

// TODO:
func minDistance(word1 string, word2 string) int {
	wl := len(word2)
	pre := make([]int, wl+1)
	cur := make([]int, wl+1)
	for i := 0; i < len(pre); i++ {
		pre[i] = i
	}

	for i := 1; i <= len(word1); i++ {
		cur[0] = i
		for j := 1; j < len(pre); j++ {
			if word1[i-1] != word2[j-1] {
				cur[j] = min(cur[j-1], pre[j-1], pre[j]) + 1
			} else {
				cur[j] = pre[j-1]
			}
		}
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		pre = tmp
	}
	return pre[wl]
}

func min(nums ...int) int {
	return slices.Min(nums)
}

var Solution = minDistance
