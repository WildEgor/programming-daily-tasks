package number_of_recent_calls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RecentCounter_1(t *testing.T) {

	r := Constructor()
	r.Ping(1)
	r.Ping(100)
	r.Ping(3001)
	r.Ping(3002)

	assert.Equal(t, "100,3001,3002", r.String())
}
