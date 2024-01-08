package arrays_solutions_move_zeroes

/**
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.
*/

func moveZeroes(nums []int) {
	buffer := make([]int, 0)
	zeroCount := 0

	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}

		buffer = append(buffer, num)
	}

	if zeroCount != 0 {
		zeroArr := make([]int, zeroCount)
		buffer = append(buffer, zeroArr...)
	}

	nums = buffer
}

func moveZeroesAlt(nums []int) {
	zeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}

var Solution = moveZeroesAlt
