package dota2_senate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_predictPartyVictory_1(t *testing.T) {
	result := predictPartyVictory("RD")
	assert.Equal(t, "Radiant", result)
}

func Test_predictPartyVictory_2(t *testing.T) {
	result := predictPartyVictory("RDD")
	assert.Equal(t, "Dire", result)
}
