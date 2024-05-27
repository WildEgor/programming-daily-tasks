package http_check

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	url2 "net/url"
	"testing"
)

var tsc = []struct {
	input  []string
	output map[string]string
}{
	{
		input: []string{"https://www.google.com", "https://www.avito.ru", "https://keklol.com"},
		output: map[string]string{
			"https://www.google.com": "https://www.google.com - ok",
			"https://www.avito.ru":   "https://www.avito.ru - ok",
			"https://keklol.com":     "https://keklol.com - not ok",
		},
	},
}

func Test_check(t *testing.T) {
	for _, test := range tsc {

		rq := NewRequester(5)
		for _, url := range test.input {
			URL, _ := url2.Parse(url)
			rq.In <- http.Request{
				Method: http.MethodGet,
				URL:    URL,
			}
		}

		rq.Do()
		go func() {
			rq.Done()
		}()

		for res := range rq.Results {
			assert.Equal(t, test.output[res.URL], res.String())
		}
	}
}
