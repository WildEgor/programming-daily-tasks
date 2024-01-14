package kth_largest_element_in_an_array

import (
	"math"
	"math/rand"
)

/**
Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Can you solve it without sorting?
@link https://leetcode.com/problems/kth-largest-element-in-an-array/
*/

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		pivotIndex := left + rand.Intn(right-left+1)
		newPivotIndex := partition(nums, left, right, pivotIndex)
		if newPivotIndex == len(nums)-k {
			return nums[newPivotIndex]
		} else if newPivotIndex > len(nums)-k {
			right = newPivotIndex - 1
		} else {
			left = newPivotIndex + 1
		}
	}
}

func partition(nums []int, left int, right int, pivotIndex int) int {
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storedIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storedIndex] = nums[storedIndex], nums[i]
			storedIndex++
		}
	}
	nums[right], nums[storedIndex] = nums[storedIndex], nums[right]
	return storedIndex
}

func findKthLargestAlt(nums []int, k int) int {
	minimum, maximum := math.MaxInt64, math.MinInt64
	n := len(nums)
	// find min and max values
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	// middle value
	n = maximum - minimum + 1
	counter := make([]int, n)
	for _, num := range nums {
		counter[num-minimum] += 1
	}
	for i := n - 1; i >= 0; i-- {
		k -= counter[i]
		if k <= 0 {
			return i + minimum
		}
	}
	return -1
}

var Solution = findKthLargestAlt
