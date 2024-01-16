package koko_eating_bananas

/**
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.
Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
Return the minimum integer k such that she can eat all the bananas within h hours.
@link https://leetcode.com/problems/koko-eating-bananas/
*/

func minEatingSpeed(piles []int, h int) int {
	indxLK := 0
	lK := piles[0]
	lenP := len(piles)
	for i := 1; i < lenP; i++ {
		if piles[i] > lK {
			indxLK = i
			lK = piles[i]
		}
	}

	if h == lenP {
		return lK
	}

	piles[0], piles[indxLK] = piles[indxLK], piles[0]

	left, right := 1, lK
	for left < right {
		mid := (left + right) / 2

		hoursLeft := h
		for i := 0; hoursLeft >= 0 && i < lenP; i++ {
			hoursLeft -= (piles[i] + mid - 1) / mid
		}

		if hoursLeft < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

var Solution = minEatingSpeed
