package ozon_5_http_check

import (
	"reflect"
	"testing"
)

var tsc = []struct {
	input  []string
	output []string
}{
	{
		input:  []string{"https://www.google.com", "https://www.yandex.ru"},
		output: []string{"https://www.google.com - ok", "https://www.yandex.ru - ok"},
	},
}

func Test_check(t *testing.T) {
	for _, test := range tsc {
		result := check(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("expected: %v, got: %v", test.output, result)
		}
	}
}
