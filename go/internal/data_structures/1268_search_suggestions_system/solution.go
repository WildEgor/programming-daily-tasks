package template

import (
	"sort"
)

/**
You are given an array of strings products and a string searchWord.
Design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.
Return a list of lists of the suggested products after each character of searchWord is typed.
@link https://leetcode.com/problems/search-suggestions-system/
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	// sort by lexicographically
	sort.Strings(products)

	collect := make(map[int][]string)
	from := 0
	for _, p := range products {
		n := lcp(searchWord, p)
		for i := from; i < n; i++ {
			ln := len(collect[i])
			if ln < 3 {
				collect[i] = append(collect[i], p)
				ln++
			}

			if ln == 3 {
				from = i + 1
			}
		}
	}

	ans := make([][]string, len(searchWord))
	for i := range ans {
		ans[i] = collect[i]
	}

	return ans
}

// longest common prefix return the length of longest common prefix
func lcp(a, b string) int {
	n := min(len(a), len(b))
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var Solution = suggestedProducts
