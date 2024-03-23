package wrap_rune

import (
	"slices"
	"strings"
)

/**
Напишите функцию, переворачивающую строку рун (вернуть строку в обратном порядке следования символов).
*/

func reverse(s string) string {
	arr := strings.Split(s, "")

	slices.Reverse(arr)

	return strings.Join(arr, "")
}
