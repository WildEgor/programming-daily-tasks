package smallest_number_in_infinite_set

import (
	"slices"
)

/**
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].
Implement the SmallestInfiniteSet class:
SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
int popSmallest() Removes and returns the smallest integer contained in the infinite set.
void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
@link https://leetcode.com/problems/smallest-number-in-infinite-set
*/

type SmallestInfiniteSet struct {
	set      map[int]bool // keep map for quick search
	smallest int          // keep smallest
}

func Constructor() SmallestInfiniteSet {
	set := make(map[int]bool)
	for i := 1; i <= 1000; i++ {
		set[i] = true
	}

	return SmallestInfiniteSet{
		set:      set,
		smallest: 1,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	ts := s.smallest
	delete(s.set, s.smallest)

	temp := make([]int, 0)
	for key := range s.set {
		temp = append(temp, key)
	}

	if len(temp) != 0 {
		s.smallest = slices.Min(temp)
	}

	return ts
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if ok := s.set[num]; ok != false {
		return
	}

	s.set[num] = true

	if num < s.smallest {
		s.smallest = num
	}
}
