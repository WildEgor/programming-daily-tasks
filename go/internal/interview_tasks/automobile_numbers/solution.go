package main

import (
	"fmt"
	"unicode"
)

/**
В Берляндии автомобильные номера состоят из цифр и прописных букв латинского алфавита. Они бывают двух видов:
либо автомобильный номер имеет вид
буква−цифра−цифра−буква−буква (примеры корректных номеров первого вида: R48FA R48FA, O00OOO00OO,A99OKA99OK);
либо автомобильный номер имеет вид буква−цифра−буква−буквабуква−цифра−буква−буква (примеры корректных номеров второго вида:T7RRT7RR,A9PQA9PQ,O0OOO0OO).
Таким образом, каждый автомобильный номер является строкой либо первого, либо второго вида.
Вам задана строка из цифр и прописных букв латинского алфавита. Можно ли разделить её пробелами на последовательность корректных автомобильных номеров?
Иными словами, проверьте, что заданная строка может быть образована как последовательность корректных автомобильных номеров, которые записаны подряд без пробелов.
В случае положительного ответа выведите любое такое разбиение.
*/

func isValidCarNumber(number string) int {
	l := len(number)

	// Проверка номера первого вида
	if l >= 5 && unicode.IsLetter(rune(number[0])) && unicode.IsDigit(rune(number[1])) && unicode.IsDigit(rune(number[2])) &&
		unicode.IsLetter(rune(number[3])) && unicode.IsLetter(rune(number[4])) {
		return 5
	}

	// Проверка номера второго вида
	if l >= 4 && unicode.IsLetter(rune(number[0])) && unicode.IsDigit(rune(number[1])) && unicode.IsLetter(rune(number[2])) &&
		unicode.IsLetter(rune(number[3])) {
		return 4
	}

	return -1
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)

		tmp := ""
		result := ""
		cmpc := 0
		ls := len(s)
		for j := 0; j < ls; j++ {
			tmp += string(s[j])
			isl := j == ls-1
			r := isValidCarNumber(tmp)

			if j > 0 && r != -1 {
				if isl {
					result += tmp
				} else {
					result += tmp + " "
				}
				cmpc += r
				tmp = ""
			}

			if isl {
				if cmpc == ls {
					fmt.Print(result)
				} else {
					fmt.Print("-")
				}
			}
		}

		fmt.Println()
	}
}
