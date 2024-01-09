package maximum_number_of_vowels_in_a_substring_of_given_length

/**
 	Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
	Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
	@link https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
*/

type Set struct {
	elements map[rune]bool
}

func NewSet(elements ...rune) *Set {
	set := map[rune]bool{}
	for _, ch := range elements {
		set[ch] = true
	}

	return &Set{elements: set}
}

func (s *Set) Add(element rune) {
	s.elements[element] = true
}

func (s *Set) Remove(element rune) {
	delete(s.elements, element)
}

func (s *Set) Has(element rune) bool {
	return s.elements[element]
}

func maxVowels(s string, k int) int {
	vowels := NewSet('a', 'e', 'i', 'o', 'u')

	max := 0
	current := 0
	for i, char := range s {
		if i >= k {
			if vowels.Has(rune(s[i-k])) {
				current--
			}
		}

		if vowels.Has(char) {
			current++
		}

		if current > max {
			max = current
		}
	}

	return max
}

var Solution = maxVowels
