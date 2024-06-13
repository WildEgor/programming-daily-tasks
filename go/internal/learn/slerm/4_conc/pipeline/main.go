package main

import (
	"strings"
	"unicode"
)

// 1 шаг - убирает лишние пробелы в строках
func step1(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for v := range in {
			out <- strings.Join(strings.Fields(v), " ")
		}
	}()

	return out
}

// 2 шаг - делит строки на предложения (на основе разделителя “точка”). Саму точку не добавляйте в результат.
func step2(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for v := range in {
			for _, sentence := range strings.Split(v, ".") {
				if sentence == "" {
					continue
				}
				out <- strings.Trim(sentence, " ")
			}
		}
	}()

	return out
}

// 3 шаг - принимает “предложения”, и в каждом предложении первые символ делает заглавным.
// Обратите внимание, что step3 должен вернуть канал, в который будет записывать.
// Это значит, что внутри функции нужно запустить отдельную горутину, читающую in.
func step3(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for v := range in {
			if v == "" {
				continue
			}
			// unicode
			r := []rune(v)
			out <- string(unicode.ToUpper(r[0])) + string(r[1:])
		}
	}()

	return out
}

func generateSentences(strs []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, str := range strs {
			out <- str
		}
		close(out)
	}()
	return out
}

func pipeline(in <-chan string) <-chan string {
	return step3(step2(step1(in)))
}

func main() {
	in := generateSentences([]string{
		"Hello World.  how are  you?",
		"Hi there! I am fine.   and you?",
	})
	out := pipeline(in)

	for v := range out {
		println(v)
	}
}
