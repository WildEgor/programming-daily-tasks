package minimum_number_of_arrows_to_burst_ballons

import "sort"

/**
There are some spherical balloons taped onto a flat wall that represents the XY-plane. The balloons are represented as a 2D integer array points where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend. You do not know the exact y-coordinates of the balloons.
Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis. A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend. There is no limit to the number of arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path.
Given the array points, return the minimum number of arrows that must be shot to burst all balloons.
@link https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons
*/

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	curr := points[0][1]
	arrows := 1
	for i := 0; i < len(points); i++ {
		if curr >= points[i][0] && curr <= points[i][1] {
			continue
		} else {
			curr = points[i][1]
			arrows++
		}
	}

	return arrows
}

var Solution = findMinArrowShots
