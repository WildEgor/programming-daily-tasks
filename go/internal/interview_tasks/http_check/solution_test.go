package http_check_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/interview_tasks/http_check"
	"net/http"
	url2 "net/url"
	"testing"
)

var tsc = []struct {
	input  []string
	output map[string]string
}{
	{
		input: []string{"https://www.google.com", "https://ozoni.ru"},
		output: map[string]string{
			"https://www.google.com": "https://www.google.com - ok",
			"https://ozoni.ru":       "https://ozoni.ru - not ok",
		},
	},
}

func Test_check(t *testing.T) {
	for _, test := range tsc {
		reqs := make([]*http.Request, 0)
		for _, url := range test.input {
			URL, _ := url2.Parse(url)
			reqs = append(reqs, &http.Request{
				Method: http.MethodGet,
				URL:    URL,
			})
		}

		rm := http_check.NewRequestManager()
		rm.Run(reqs)

		for res := range rm.OUT {
			assert.Equal(t, test.output[res.URL], res.String())
		}
	}
}
