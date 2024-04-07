package homework

import (
	"io"
	"slices"
	"strings"
)

func reverseReader(reader io.Reader) (result []string, err error) {

	bs, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	result = strings.Split(string(bs[:]), "\n")

	slices.Reverse(result)

	return
}
