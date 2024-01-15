package maximum_subsequence_score

import (
	"container/heap"
	"sort"
)

/**
You are given two 0-indexed integer arrays nums1 and nums2 of equal length n and a positive integer k. You must choose a subsequence of indices from nums1 of length k.
For chosen indices i0, i1, ..., ik - 1, your score is defined as:
The sum of the selected elements from nums1 multiplied with the minimum of the selected elements from nums2.
It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).
Return the maximum possible score.
A subsequence of indices of an array is a set that can be derived from the set {0, 1, ..., n-1} by deleting some or no elements.
@link https://leetcode.com/problems/maximum-subsequence-score
*/

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() any {
	x := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return x
}

type pair struct {
	a, b int
}

// TODO: explain
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := make([]pair, len(nums1))
	for i, n := range nums1 {
		pairs[i] = pair{n, nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].b > pairs[j].b })

	sum := 0
	h := &Heap{}
	for i := 0; i < k; i++ {
		sum += pairs[i].a
		heap.Push(h, pairs[i].a)
	}

	max := sum * pairs[k-1].b
	for i := k; i < len(pairs); i++ {
		sum += pairs[i].a
		heap.Push(h, pairs[i].a)
		sum -= heap.Pop(h).(int)
		res := sum * pairs[i].b
		if res > max {
			max = res
		}
	}

	return int64(max)
}

var Solution = maxScore
