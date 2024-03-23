package str_len_without_spaces

import (
	"slices"
	"strings"
)

func stringLengthWithoutSpaces(str string) int {
	arr := strings.Split(str, " ")

	slices.DeleteFunc(arr, func(s string) bool {
		return len(s) == 0 || s == "\t"
	})

	return len(arr)
}
