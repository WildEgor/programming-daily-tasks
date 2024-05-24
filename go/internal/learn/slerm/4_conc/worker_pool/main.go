package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	id    int
	ctx   context.Context
	tasks <-chan string
	wg    *sync.WaitGroup
	out   chan string
}

// NewWorker creates a new worker.
func NewWorker(ctx context.Context, id int, wg *sync.WaitGroup, tasks <-chan string, out chan string) *Worker {
	return &Worker{
		id,
		ctx,
		tasks,
		wg,
		out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	go func() {
		defer w.wg.Done()

		fmt.Println("Starting worker")
		select {
		case t := <-w.tasks:
			fmt.Printf("Worker %d task=%s\n", w.id, t)

			hasher := md5.New()
			hasher.Write([]byte(t))
			hash := hex.EncodeToString(hasher.Sum(nil))

			w.out <- fmt.Sprintf("Task: %s, Hash: %s", t, hash)
		default:
			return
		}
	}()
}

// WorkerPool represents a pool of workers.
type WorkerPool struct {
	timeout time.Duration
	nw      int
	tasks   chan string
	result  chan string
	wg      sync.WaitGroup
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool(numWorkers int, timeout time.Duration) (*WorkerPool, error) {
	return &WorkerPool{
		timeout: timeout,
		nw:      numWorkers,
		result:  make(chan string),
		tasks:   make(chan string),
	}, nil
}

// Run starts the worker pool.
func (wp *WorkerPool) Run(ctx context.Context) {
	fmt.Println("Start WorkerPool...")

	wp.wg.Add(wp.nw)
	for i := 0; i < wp.nw; i++ {
		w := NewWorker(ctx, i, &wp.wg, wp.tasks, wp.result)
		w.Run()
	}

	wp.wg.Wait()
}

// AddTask adds a task to the worker pool.
func (wp *WorkerPool) AddTask(t string) {
	wp.tasks <- t
}

// GetResults print results
func (wp *WorkerPool) GetResults() {
	for {
		select {
		case res, ok := <-wp.result:
			if !ok {
				return
			}
			fmt.Printf("Result: %s\n", res)
		default:
			return
		}
	}
}

func (wp *WorkerPool) Stop() {
}

func main() {
	wp, err := NewWorkerPool(3, 15*time.Second)
	if err != nil {
		fmt.Printf("error creating worker pool: %v", err)
		return
	}

	ctx := context.Background()
	tasks := []string{"hello", "world", "!"}
	wp.Run(ctx)
	for _, task := range tasks {
		wp.AddTask(task)
	}
	wp.Stop()

	wp.GetResults()
}
