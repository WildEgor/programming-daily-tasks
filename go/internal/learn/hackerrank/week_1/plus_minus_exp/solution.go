package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Task func() (any, error)
type Result struct {
	wid    int
	result any
	err    error
}

type Worker struct {
	id int
	tq <-chan Task
	rq chan<- Result

	wg *sync.WaitGroup
}

func (w *Worker) Run() {
	go func() {
		defer w.wg.Done()

		for task := range w.tq {
			r, err := task()
			w.rq <- Result{wid: w.id, result: r, err: err}
		}
	}()
}

func (w *Worker) Stop() {

}

type WorkerPool struct {
	tq chan Task
	rc chan Result
	wc int

	wg sync.WaitGroup
}

func NewWorkerPool(wc int) *WorkerPool {
	return &WorkerPool{
		tq: make(chan Task),
		rc: make(chan Result, wc),
		wc: wc,
	}
}

func (wp *WorkerPool) Start() {
	wp.wg.Add(wp.wc)

	for i := 0; i < wp.wc; i++ {
		worker := Worker{wg: &wp.wg, id: i, tq: wp.tq, rq: wp.rc}
		worker.Run()
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.tq)
	wp.wg.Wait()
	close(wp.rc)
}

func (wp *WorkerPool) Submit(task Task) {
	wp.tq <- task
}

func (wp *WorkerPool) GetResult() chan Result {
	return wp.rc
}

type CountResult struct {
	Length     int
	Pos        int32
	PosRation  float64
	Neg        int32
	NegRation  float64
	Zero       int32
	ZeroRation float64
}

func plusMinus(arr []int32) {
	// Crop slice to chunks
	chunkSize := 5
	arrLength := len(arr)
	var chunks [][]int32
	for chunkSize < len(arr) {
		arr, chunks = arr[chunkSize:], append(chunks, arr[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, arr)

	wp := NewWorkerPool(chunkSize)
	wp.Start()

	// Submit all tasks
	for _, chunk := range chunks {
		chunk := chunk // Capture range variable
		wp.Submit(func() (any, error) {
			r := &CountResult{}
			for _, value := range chunk {
				if value < 0 {
					r.Neg += 1
				} else if value > 0 {
					r.Pos += 1
				} else {
					r.Zero += 1
				}
			}

			r.NegRation = float64(r.Neg) / float64(arrLength)
			r.PosRation = float64(r.Pos) / float64(arrLength)
			r.ZeroRation = float64(r.Zero) / float64(arrLength)

			return r, nil
		})
	}

	// Signal that no more tasks will be submitted
	wp.Stop()

	c := &CountResult{}

	// Collect results
	for result := range wp.GetResult() {
		switch r := result.result.(type) {
		case *CountResult:
			c.PosRation += r.PosRation
			c.NegRation += r.NegRation
			c.ZeroRation += r.ZeroRation
		}
	}
	fmt.Printf("%f\n", c.PosRation)
	fmt.Printf("%f\n", c.NegRation)
	fmt.Printf("%f\n", c.ZeroRation)
}

func main() {
	data := readData()

	plusMinus(data)
}

func readData() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var arr []int32
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	return arr
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
