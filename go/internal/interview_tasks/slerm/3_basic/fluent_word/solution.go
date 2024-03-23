package homework

import "strings"

/**
Напишите функцию для поиска наиболее часто встречающегося слова в строке.
Программа должна использовать map для хранения количества слов. Верните найденное слово.
В случае, если таких слов несколько, верните любое из них.
*/

func frequentWord(str string) string {
	if len(str) == 0 {
		return ""
	}

	wf := make(map[string]int)
	lf := make([]int, 2)

	words := strings.Split(str, " ")

	for i, word := range words {
		wf[word] += 1

		if lf[1] <= wf[word] {
			lf[0] = i
			lf[1] = wf[word]
		}
	}

	return words[lf[0]] // first word or by freq
}
