package total_cost_to_hire_k_workers

import (
	"container/heap"
)

/**
You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the ith worker.
You are also given two integers k and candidates. We want to hire exactly k workers according to the following rules:
You will run k sessions and hire exactly one worker in each session.
In each hiring session, choose the worker with the lowest cost from either the first candidates workers or the last candidates workers. Break the tie by the smallest index.
For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session, we will choose the 4th worker because they have the lowest cost [3,2,7,7,1,2].
In the second hiring session, we will choose 1st worker because they have the same lowest cost as 4th worker but they have the smallest index [3,2,7,7,2]. Please note that the indexing may be changed in the process.
If there are fewer than candidates workers remaining, choose the worker with the lowest cost among them. Break the tie by the smallest index.
A worker can only be chosen once.
Return the total cost to hire exactly k workers.
@link https://leetcode.com/problems/total-cost-to-hire-k-workers
*/

func totalCost(costs []int, k int, candidates int) int64 {
	minHeap := &MinHeap{}

	// Add first candidates to heap
	for i := 0; i < candidates; i++ {
		heap.Push(minHeap, []int{costs[i], 0})
	}

	// Add last candidates to heap
	for i := max(candidates, len(costs)-candidates); i < len(costs); i++ {
		heap.Push(minHeap, []int{costs[i], 1})
	}

	left, right := candidates, len(costs)-1-candidates
	res := 0

	for i := 0; i < k; i++ {
		candidate := heap.Pop(minHeap).([]int)
		res += candidate[0]
		idx := candidate[1]

		if left <= right {
			if idx == 0 {
				heap.Push(minHeap, []int{costs[left], 0})
				left++
			} else {
				heap.Push(minHeap, []int{costs[right], 1})
				right--
			}
		}
	}

	return int64(res)
}

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Solution = totalCost
