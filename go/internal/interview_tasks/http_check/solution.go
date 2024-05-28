package http_check

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RequestResult struct {
	URL string
	// Body []byte
	Code int
}

func (rr RequestResult) String() string {
	sm := "ok"
	if rr.Code > http.StatusOK {
		sm = "not ok"
	}
	return fmt.Sprintf("%s - %s", rr.URL, sm)
}

type RequestManager struct {
	wg  sync.WaitGroup
	OUT chan *RequestResult
}

func NewRequestManager() *RequestManager {
	return &RequestManager{
		OUT: make(chan *RequestResult),
	}
}

func (rm *RequestManager) Run(reqs []*http.Request) {
	rm.wg.Add(len(reqs))
	for i := 0; i < len(reqs); i++ {
		go rm.fetchURL(reqs[i])
	}

	go func() {
		rm.wg.Wait()
		close(rm.OUT)
	}()
}

func (rm *RequestManager) fetchURL(req *http.Request) {
	defer rm.wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	rr := &RequestResult{
		URL:  req.URL.String(),
		Code: http.StatusInternalServerError,
	}

	resp, err := client.Do(req)
	if err != nil {
		rm.OUT <- rr
		return
	}

	//result, err := io.ReadAll(resp.Body)
	//defer resp.Body.Close()
	//if err != nil {
	//	rm.OUT <- rr
	//	return
	//}
	//
	//rr.Body = result
	rr.Code = resp.StatusCode

	rm.OUT <- rr
}
