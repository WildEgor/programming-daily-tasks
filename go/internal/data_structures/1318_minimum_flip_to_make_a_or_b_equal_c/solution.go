package template

/**
Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.
@link https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
*/

func minFlips(a int, b int, c int) int {
	var bA, bB, bC = toBinary(a), toBinary(b), toBinary(c)
	counter := 0

	for i := 0; i < len(bA); i++ {
		if bC[i] == 0 {
			if bA[i] == 1 {
				counter++
				bA[i] = 0
			}
			if bB[i] == 1 {
				counter++
				bB[i] = 0
			}
		} else {
			if bA[i] == 0 && bB[i] == 0 {
				counter++
				bA[i] = 1
			}
		}
	}

	return counter
}

func toBinary(n int) []int {
	res := make([]int, 32)
	for i := 0; i < 32; i++ {
		res[i] = n & 1 // n % 2
		n >>= 1        // n /= 2
	}
	return res
}

var Solution = minFlips
