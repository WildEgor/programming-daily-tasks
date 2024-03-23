package remove_spaces

import (
	"slices"
	"strings"
)

func removeSpaces(str string) string {
	arr := strings.Split(str, " ")

	slices.DeleteFunc(arr, func(s string) bool {
		return len(s) == 0
	})

	return strings.Join(arr, "")
}
