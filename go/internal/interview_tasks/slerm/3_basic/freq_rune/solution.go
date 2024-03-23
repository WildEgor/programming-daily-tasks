package freq_rune

import "math"

func frequentRune(str string) rune {
	freq := make(map[rune]int)

	for _, c := range str {
		freq[c] += 1
	}

	var r rune
	mn := math.MinInt32
	for k, v := range freq {
		if v > mn {
			mn = v
			r = k
		}
	}

	return r
}
