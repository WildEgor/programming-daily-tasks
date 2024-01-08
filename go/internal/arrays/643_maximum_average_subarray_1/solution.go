package arrays_643_maximum_average_subarray_1

import "math"

/**
You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.
@link https://leetcode.com/problems/maximum-average-subarray-i/
*/

func findMaxAverage(nums []int, k int) float64 {
	sum := sum(nums[:k])

	mV := sum
	pointer := k

	for pointer < len(nums) {
		sum = sum - nums[pointer-k] + nums[pointer]
		if sum > mV {
			mV = sum
		}
		pointer++
	}

	return math.Round(float64(mV)/float64(k)*100000) / 100000
}

func sum(nums []int) int {
	sum := 0
	if len(nums) == 0 {
		return sum
	}

	for _, e := range nums {
		sum += e
	}

	return sum
}

var Solution = findMaxAverage
