package find_pivot_index

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == sum-leftSum-v {
			return i
		}
		leftSum += v
	}

	return -1
}

var Solution = pivotIndex
