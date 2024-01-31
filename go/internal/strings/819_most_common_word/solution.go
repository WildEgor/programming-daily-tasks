package most_common_word

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	wordCounter := make(map[string]int)

	f := func(r rune) bool { return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z') }

	for _, word := range strings.FieldsFunc(paragraph, f) {
		wordCounter[strings.ToLower(word)]++
	}

	for _, b := range banned {
		delete(wordCounter, b)
	}

	w := ""
	mv := 0
	for word, v := range wordCounter {
		if v > mv {
			mv = v
			w = word
		}
	}

	return w
}

var Solution = mostCommonWord
