package template

import stack2 "github.com/wildegor/daily_tasks/go/internal/data_structures/stack"

/**
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i]
is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible,
keep answer[i] == 0 instead.
*/

func dailyTemperatures(temperatures []int) []int {
	stack := stack2.NewStack()

	result := make([]int, len(temperatures), len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !stack.Empty() && temperatures[i] >= temperatures[stack.Top().(int)] {
			stack.Pop()
		}

		if stack.Empty() {
			result[i] = 0
		} else {
			// stack.Top() is the index of the next warmer temperature
			result[i] = stack.Top().(int) - i
		}

		stack.Push(i)
	}

	return result
}

var Solution = dailyTemperatures
