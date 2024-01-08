package strings_is_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubsequence_1(t *testing.T) {
	str := "abc"
	substr := "ahbgdc"
	result := isSubsequence(str, substr)

	assert.Equal(t, true, result)
}

func Test_isSubsequence_2(t *testing.T) {
	str := "axc"
	substr := "ahbgdc"
	result := isSubsequence(str, substr)

	assert.Equal(t, false, result)
}

func Test_isSubsequence_3(t *testing.T) {
	str := ""
	substr := "ahbgdc"
	result := isSubsequence(str, substr)

	assert.Equal(t, true, result)
}

func Test_isSubsequence_4(t *testing.T) {
	str := "b"
	substr := "abc"
	result := isSubsequence(str, substr)

	assert.Equal(t, true, result)
}
