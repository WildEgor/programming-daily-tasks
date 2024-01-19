package majority_element

/**
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
@link https://leetcode.com/problems/majority-element
*/

func majorityElement(nums []int) int {
	unq := map[int]int{}

	cmp := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		unq[nums[i]]++
	}

	for k, v := range unq {
		if v > cmp {
			return k
		}
	}

	return -1
}

func majorityElementAlt(nums []int) int {
	count := 1
	k := nums[0]

	for i := 1; i < len(nums); i++ {
		if k == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			k = nums[i]
			count = 1
		}
	}

	return k
}

var Solution = majorityElementAlt
