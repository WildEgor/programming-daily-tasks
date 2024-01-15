package combination_sum_iii

/**
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.
@link https://leetcode.com/problems/combination-sum-iii/
*/

func combinationSum3(k int, n int) [][]int {
	if k == 0 || n == 0 {
		return [][]int{}
	}

	return dfs(k, n, 0, []int{})
}

/*
*
В функции dfs происходит следующее:
Если длина текущей комбинации (source) равна k и текущая сумма (sum) равна n, то текущая комбинация добавляется в список ответов.
Затем функция определяет начальное значение i для следующего числа,
которое будет добавлено в комбинацию. Если текущая комбинация не пуста, то i устанавливается равным следующему числу после последнего числа в комбинации.
В противном случае i устанавливается равным 1.
Затем функция определяет максимальное значение maxValue для следующего числа,
которое будет добавлено в комбинацию. Это значение равно разности между n и текущей суммой (sum). Если maxValue больше 9, то оно устанавливается равным 9.
Затем функция проходит по всем числам от i до maxValue включительно.
Для каждого числа она создает новую комбинацию, добавляя это число к текущей комбинации, и вызывает функцию dfs с новой комбинацией и обновленной суммой.
Все возвращенные комбинации добавляются в список ответов.
В конце функция возвращает список всех найденных комбинаций.
*/
func dfs(k, n, sum int, source []int) [][]int {
	currLen := len(source)
	if currLen == k && sum == n {
		return [][]int{source}
	}

	i := 1
	if currLen > 0 {
		i = source[currLen-1] + 1
	}

	maxValue := n - sum
	if maxValue > 9 {
		maxValue = 9
	}

	ans := make([][]int, 0)
	for i <= maxValue {
		temp := make([]int, len(source))
		copy(temp, source)
		temp = append(temp, i)
		ans = append(ans, dfs(k, n, sum+i, temp)...)
		temp = temp[:currLen]
		i++
	}

	return ans
}

var Solution = combinationSum3
