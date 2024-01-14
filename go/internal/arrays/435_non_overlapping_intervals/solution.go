package template

import "sort"

/**
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
@link https://leetcode.com/problems/non-overlapping-intervals/
*/

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	// sort by end
	sort.Slice(intervals, func(i, j int) bool {
		// if end is equal, sort by start
		return intervals[i][1] < intervals[j][1]
	})

	// greedy
	count := 1
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}

	return n - count
}

var Solution = eraseOverlapIntervals
