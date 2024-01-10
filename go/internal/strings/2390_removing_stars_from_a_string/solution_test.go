package removing_stars_from_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeStars_1(t *testing.T) {
	result := removeStars("leet**cod*e")
	assert.Equal(t, "lecoe", result)
}

func Test_removeStars_2(t *testing.T) {
	result := removeStars("erase*****")
	assert.Equal(t, "", result)
}

func Test_removeStars_3(t *testing.T) {
	result := removeStars("*eet**cod*e")
	assert.Equal(t, "ecoe", result)
}
