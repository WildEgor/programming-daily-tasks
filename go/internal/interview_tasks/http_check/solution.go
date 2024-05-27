package http_check

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/**
@input: []string
@output: []string
@description: Дан массив URL-адресов. Ваша задача - отправить HTTP-запрос на каждый из них и вернуть массив строк, содержащий статус-коды ответов на запросы.
*/

type RequestResult struct {
	URL  string
	Code int
}

func (rr RequestResult) String() string {
	sm := "ok"
	if rr.Code > http.StatusOK {
		sm = "not ok"
	}

	return fmt.Sprintf("%s - %s", rr.URL, sm)
}

type Requester struct {
	N       int
	In      chan http.Request
	Results chan *RequestResult
	wg      sync.WaitGroup
}

func NewRequester(n int) *Requester {
	return &Requester{
		N:       n,
		In:      make(chan http.Request, n),
		Results: make(chan *RequestResult),
	}
}

func (rq *Requester) Do() {
	rq.wg.Add(rq.N)
	for range rq.N {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for req := range rq.In {
				client := http.Client{
					Timeout: 15 * time.Second,
				}

				resp, err := client.Do(&req)
				result := RequestResult{
					URL:  req.URL.String(),
					Code: http.StatusInternalServerError,
				}
				if err != nil {
					rq.Results <- &result
					continue
				}
				defer resp.Body.Close()

				rq.Results <- &RequestResult{
					URL:  req.URL.String(),
					Code: resp.StatusCode,
				}
			}
		}(&rq.wg)
	}
}

func (rq *Requester) Done() {
	rq.wg.Wait()
	close(rq.Results)
}
