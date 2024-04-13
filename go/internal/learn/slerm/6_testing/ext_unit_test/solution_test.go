package ext_unit_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	eu "github.com/wildegor/daily_tasks/go/internal/learn/slerm/6_testing/ext_unit_test"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Cases
var tcs = map[string]struct {
	result string
}{
	"test": {
		result: "https://httpbin.org/post",
	},
	"test-error": {
		result: "",
	},
}

// Simple mock
type HttpClientMock struct {
}

func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	rb := io.NopCloser(bytes.NewReader([]byte(`{"url":"https://httpbin.org/post"}`)))

	return &http.Response{
		StatusCode: 200,
		Body:       rb,
	}, nil
}

// Test 1
func TestNewRequester(t *testing.T) {
	httpMock := &HttpClientMock{}
	r := eu.NewRequester("https://httpbin.org", httpMock)

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			data, err := r.PostData()
			if err != nil {
				t.Fatal(err.Error())
				return
			}

			assert.Equal(t, tc.result, data)
		})
	}
}

// Test 2
func TestNewRequesterWithMock(t *testing.T) {
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"url":"https://httpbin.org/post"}`))
			}))
			defer server.Close()

			r := eu.NewRequester(server.URL, server.Client())

			data, err := r.PostData()
			if err != nil {
				t.Fatal(err.Error())
			}

			assert.Equal(t, tc.result, data)
		})
	}
}
