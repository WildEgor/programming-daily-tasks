package letter_combinations_of_a_phone_number

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// create dict digits = letters
	digitsMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// cause combinations n!
	size := 1
	for _, d := range digits {
		size *= len(digitsMap[d])
	}

	result := make([]string, size)
	for _, char := range digits {
		curr := digitsMap[char]
		size /= len(curr)
		for i := range result {
			result[i] += string(curr[(i/size)%len(curr)])
		}
	}

	return result
}

var Solution = letterCombinations
