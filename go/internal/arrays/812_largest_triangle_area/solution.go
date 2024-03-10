package largest_triangle_area

import "math"

/**
Given an array of points on the X-Y plane points where points[i] = [xi, yi], return the area of the largest triangle
that can be formed by any three different points. Answers within 10-5 of the actual answer will be accepted.
@link https://leetcode.com/problems/largest-triangle-area
*/

func largestTriangleArea(points [][]int) float64 {
	s := 0.0
	pl := len(points)

	for i := 0; i < pl; i++ { // x
		for j := i + 1; j < pl; j++ { // y
			for k := j + 1; k < pl; k++ { // by all points
				area := calcArea(points, i, j, k)
				if area > s {
					s = area
				}
			}
		}
	}

	return s
}

func calcArea(points [][]int, i, j, k int) float64 {
	area := 0

	area += points[i][0] * (points[j][1] - points[k][1])
	area += points[j][0] * (points[k][1] - points[i][1])
	area += points[k][0] * (points[i][1] - points[j][1])

	return math.Abs(float64(area / 2))
}

var Solution = largestTriangleArea
