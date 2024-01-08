package arrays_11_container_with_most_water

/**
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints
of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.
@link https://leetcode.com/problems/container-with-most-water
*/

func maxArea(height []int) int {
	result := 0

	for i, j := 0, len(height)-1; i < j; {
		var h int
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}

		if area := h * (j - i + 1); area > result {
			result = area
		}
	}

	return result
}

var Solution = maxArea
