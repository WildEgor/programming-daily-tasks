package successful_pairs_of_spells_and_positions

import (
	"slices"
	"sort"
)

/**
You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.
You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.
Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.
@link https://leetcode.com/problems/successful-pairs-of-spells-and-potions
*/

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	result := make([]int, 0)

	for i := 0; i < len(spells); i++ {
		for j := 0; j < len(potions); j++ {
			product := int64(potions[j] * spells[i])

			if product >= success {
				result = append(result, len(potions)-j)
				break
			} else if j == len(potions)-1 {
				result = append(result, 0)
			}
		}
	}

	return result
}

func successfulPairsAlt(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})

	var result []int
	for _, spell := range spells {
		val := -1
		ind := bs(0, len(potions)-1, spell, success, &val, potions)
		if ind == -1 {
			result = append(result, 0)
			continue
		}
		result = append(result, len(potions)-ind)
	}

	return result

}

func bs(s, e, m int, success int64, val *int, pos []int) int {
	if s > e {
		return *val
	}

	mid := s + ((e - s) / 2)
	product := int64(m * pos[mid])

	if product >= success {
		*val = mid
		return bs(s, mid-1, m, success, val, pos)
	}

	if product < success {
		return bs(mid+1, e, m, success, val, pos)
	}

	return *val

}

var Solution = successfulPairsAlt
