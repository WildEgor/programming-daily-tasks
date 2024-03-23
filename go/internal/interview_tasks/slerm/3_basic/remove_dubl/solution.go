package remove_dubl

/**
Напишите функцию для удаления дублей из слайса.
*/

func removeDuplicates(slice []int) []int {
	unique := make(map[int]bool)
	result := make([]int, 0)

	for _, value := range slice {
		if ok := unique[value]; !ok {
			result = append(result, value)
			unique[value] = true
		}
	}

	return result
}
