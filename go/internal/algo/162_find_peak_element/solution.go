package find_peak_element

/**
A peak element is an element that is strictly greater than its neighbors.
Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.
You must write an algorithm that runs in O(log n) time.
@link https://leetcode.com/problems/find-peak-element/
*/

func findPeakElement(nums []int) int {
	nlen := len(nums)
	if nlen == 1 {
		return 0
	}

	l, r := 0, nlen-1
	// infinite loop!!!
	for {
		mindx := l + (r-l)/2 // find middle index

		// if first
		if mindx == 0 {
			if nums[mindx] > nums[mindx+1] {
				return mindx
			}

			l = mindx + 1
			continue
		}

		// If last
		if mindx == nlen-1 {
			if nums[mindx] > nums[mindx-1] {
				return mindx
			}

			r = mindx - 1
			continue
		}

		// Check with neighbors
		if nums[mindx] > nums[mindx-1] && nums[mindx] > nums[mindx+1] {
			return mindx
		}

		if nums[mindx] < nums[mindx-1] {
			r = mindx - 1
		} else {
			l = mindx + 1
		}
	}

	return -1
}

var Solution = findPeakElement
