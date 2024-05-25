package http_check

import (
	"net/http"
	"sync"
	"time"
)

/**
@input: []string
@output: []string
@description: Дан массив URL-адресов. Ваша задача - отправить HTTP-запрос на каждый из них и вернуть массив строк, содержащий статус-коды ответов на запросы.
*/

type result struct {
	url  string
	code int
}

func sendRequest(url string, wg *sync.WaitGroup, rCh chan<- result) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		rCh <- result{
			url:  url,
			code: http.StatusInternalServerError,
		}
		return
	}
	defer resp.Body.Close()

	rCh <- result{
		url:  url,
		code: resp.StatusCode,
	}
}

func check(urls []string) []string {
	// Edge case 1: empty input
	if len(urls) == 0 {
		return []string{}
	}

	var wg sync.WaitGroup
	rCh := make(chan result)

	for _, url := range urls {
		wg.Add(1)
		go sendRequest(url, &wg, rCh)
	}

	go func() {
		wg.Wait()
		close(rCh)
	}()

	data := make([]string, 0)

	for r := range rCh {
		if r.code == http.StatusOK {
			data = append(data, r.url+" - ok")
		} else {
			data = append(data, r.url+" - not ok")
		}
	}

	return data
}

var Solution = check
