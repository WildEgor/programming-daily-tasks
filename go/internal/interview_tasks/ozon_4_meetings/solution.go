package ozon_4_meetings

/*
 @input: [][]int
 @output: [][]int
 @description: Дан массив интервалов, в котором каждый элемент представляет собой массив из двух целых чисел. Ваша задача - объединить все пересекающиеся интервалы в один и вернуть новый массив непересекающихся интервалов.
*/

func merge(intervals [][]int) [][]int {
	// Если интервалов нет или один, то возвращаем исходный массив
	if len(intervals) <= 1 {
		return intervals
	}

	// Сортируем интервалы по началу
	quickSort(intervals, 0, len(intervals)-1)

	result := make([][]int, 0)
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]

		if last[1] >= intervals[i][0] {
			result[len(result)-1][1] = max(last[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func quickSort(intervals [][]int, low, high int) {
	if low < high {
		pivot := partition(intervals, low, high)

		quickSort(intervals, low, pivot-1)
		quickSort(intervals, pivot+1, high)
	}
}

func partition(intervals [][]int, low, high int) int {
	pivot := intervals[high][0]
	i := low

	for j := low; j < high; j++ {
		if intervals[j][0] < pivot {
			intervals[i], intervals[j] = intervals[j], intervals[i]
			i++
		}
	}

	intervals[i], intervals[high] = intervals[high], intervals[i]

	return i
}
