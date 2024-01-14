package template

import (
	"testing"
)

var tsc = []struct {
	n    int
	pick int
	ans  int
}{
	{
		n:    10,
		pick: 6,
		ans:  6,
	},
}

func Test_example_1(t *testing.T) {
	for _, ts := range tsc {
		got := Solution(ts.pick)
		if got != ts.ans {
			t.Errorf("input: %v, expected: %v, but got: %v", ts.pick, ts.ans, got)
		}
	}
}
