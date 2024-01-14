package template

/**
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
You call a pre-defined API int guess(int num), which returns three possible results:
-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.
@link https://leetcode.com/problems/guess-number-higher-or-lower/
*/

func guess(num int) int {
	if num == 6 {
		return 0
	}

	if num > 6 {
		return -1
	}

	return 1
}

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		guessResult := guess(mid)
		if guessResult == 0 {
			return mid
		} else if guessResult == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

var Solution = guessNumber
