package ozon_4_meetings

import (
	"sort"
)

/*
input: [][]int
output: [][]int
Дан массив интервалов, в котором каждый элемент представляет собой массив из двух целых чисел.
Ваша задача - объединить все пересекающиеся интервалы в один и вернуть новый массив непересекающихся интервалов.
*/

func merge(intervals [][]int) [][]int {
	// Если интервалов нет или один, то возвращаем исходный массив
	if len(intervals) <= 1 {
		return intervals
	}

	// Сортируем интервалы по началу
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Инициализируем "последний" элемент
	result := make([][]int, 0)
	result = append(result, intervals[0])

	// Идем по отсортированному и сравниваем с последним значением
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
