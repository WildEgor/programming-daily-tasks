package determine_if_two_strings_are_close

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_closeStrings_1(t *testing.T) {
	result := closeStrings("abc", "bca")
	assert.Equal(t, true, result)
}

func Test_closeStrings_2(t *testing.T) {
	result := closeStrings("a", "aa")
	assert.Equal(t, false, result)
}

func Test_closeStrings_3(t *testing.T) {
	result := closeStrings("cabbba", "abbccc")
	assert.Equal(t, true, result)
}
