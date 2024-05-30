package main

import (
	"fmt"
	"sort"
)

// kSumSimple using 2 loops O(n2)
func kSumSimple(nums []int, k int) [][]int {
	r := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == k {
				r = append(r, []int{nums[i], nums[j]})
			}
		}
	}

	return r
}

func bs(in []int, n int) (int, bool) {
	fp, lp := 0, len(in)-1

	for fp <= lp {
		mp := ((lp - fp) / 2) + fp

		if in[mp] == n {
			return in[mp], true
		} else if in[mp] < n {
			fp = mp + 1
		} else if in[mp] > n {
			lp = mp - 1
		}
	}

	return -1, false
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func kSumBinarySearch(nums []int, k int) [][]int {
	r := make([][]int, 0)

	sort.Ints(nums) // HINT: sort nums ASC

	for i := 0; i < len(nums); i++ {
		base := nums[i]                       // HINT: ex. k = 5 and nums[i] = -5. 5 - (-5) = 10, so find 10
		in := append(nums[:i], nums[i+1:]...) // HINT: remove nums[i] from base
		fe := k - base

		le, exists := bs(in, fe)
		if !exists {
			continue
		}

		r = append(r, []int{base, le})
	}

	return r
}

func main() {
	in := []int{1, 4, 3, 5, -1, 0, 2}
	k := 2

	r := kSumBinarySearch(in, k)

	fmt.Println(r)
}
