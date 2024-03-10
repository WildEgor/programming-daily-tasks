package ozon_5_http_check

import (
	"net/http"
)

/**
@input: []string
@output: []string
@description: Дан массив URL-адресов. Ваша задача - отправить HTTP-запрос на каждый из них и вернуть массив строк, содержащий статус-коды ответов на запросы.
*/

func check(urls []string) []string {
	result := make([]string, 0)
	for _, url := range urls {
		func() {
			r, err := http.Get(url)
			defer r.Body.Close()

			if r == nil || err != nil {
				result = append(result, url+" - not ok")
				return
			}

			if r.StatusCode != http.StatusOK {
				result = append(result, url+" - not ok")
				return
			}

			result = append(result, url+" - ok")
		}()
	}

	return result
}
