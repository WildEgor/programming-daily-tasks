package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
)

// ========================================== Worker Start

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {

	go func() {
		log.Printf("Starting worker")
		defer w.wg.Done()

		for task := range w.tasks {
			log.Printf("Worker task=%s", task)

			hasher := md5.New()
			hasher.Write([]byte(task))
			hash := hex.EncodeToString(hasher.Sum(nil))

			w.out <- fmt.Sprintf("Task: %s, Hash: %s", task, hash)
		}
	}()
}

// ========================================== Worker End

// ========================================== WorkerPool Start

// WorkerPool represents a pool of workers.
type WorkerPool struct {
	nw     int
	tasks  chan string
	result chan string
	wg     sync.WaitGroup
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool(numWorkers int) (*WorkerPool, error) {
	return &WorkerPool{
		nw:     numWorkers,
		result: make(chan string),
		tasks:  make(chan string),
	}, nil
}

// Run starts the worker pool.
func (wp *WorkerPool) Run() {
	log.Print("Start WorkerPool...")
	wp.startWorkers()
}

// AddTask adds a task to the worker pool.
func (wp *WorkerPool) AddTask(t string) {
	wp.tasks <- t
}

func (wp *WorkerPool) GetResults() {

	for {
		select {
		case res := <-wp.result:
			log.Printf("Result: %s\n", res)
		default:
			break
		}
	}
}

func (wp *WorkerPool) startWorkers() {
	wp.wg.Add(wp.nw)
	for i := 0; i < wp.nw; i++ {
		w := NewWorker(wp.tasks, &wp.wg, wp.result)
		w.Run()
	}

	go func() {
		wp.wg.Wait()
		close(wp.result)
	}()
}

// ========================================== WorkerPool End

func main() {
	wp, err := NewWorkerPool(3)
	if err != nil {
		log.Fatalf("Error creating worker pool: %v", err)
	}

	wp.Run()

	wp.AddTask("hello")
	wp.AddTask("world")
	wp.AddTask("test")

	wp.GetResults()
}
