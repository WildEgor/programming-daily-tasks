package remove_duplicates_from_sorted_array

/**
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
@link https://leetcode.com/problems/remove-duplicates-from-sorted-array
*/

func removeDuplicates(nums []int) int {
	k := 0 // pointer to prev
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i] // "merge"
		}
	}

	return k + 1
}

var Solution = removeDuplicates
