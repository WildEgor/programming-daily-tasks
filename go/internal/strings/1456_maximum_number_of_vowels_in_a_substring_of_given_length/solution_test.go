package maximum_number_of_vowels_in_a_substring_of_given_length

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxVowels_1(t *testing.T) {
	result := Solution("abciiidef", 3)
	assert.Equal(t, 3, result)
}

func Test_maxVowels_2(t *testing.T) {
	result := Solution("aeiou", 2)
	assert.Equal(t, 2, result)
}

func Test_maxVowels_3(t *testing.T) {
	result := Solution("leetcode", 3)
	assert.Equal(t, 2, result)
}
