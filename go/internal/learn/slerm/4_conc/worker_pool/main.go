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
	ctx   context.Context
	tasks <-chan string
	wg    *sync.WaitGroup
	out   chan string
}

// NewWorker creates a new worker.
func NewWorker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan string, out chan string) *Worker {
	return &Worker{
		ctx,
		tasks,
		wg,
		out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	go func() {
		fmt.Println("Starting worker")
		defer w.wg.Done()

		select {
		case <-w.ctx.Done():
			return
			// HINT: Так правильно?
		default:
			// HINT: Так правильно?
			for task := range w.tasks {
				fmt.Printf("Worker task=%s\n", task)

				hasher := md5.New()
				hasher.Write([]byte(task))
				hash := hex.EncodeToString(hasher.Sum(nil))

				// Some heavy load
				time.Sleep(5 * time.Second)

				w.out <- fmt.Sprintf("Task: %s, Hash: %s", task, hash)
			}
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
func (wp *WorkerPool) Run() {
	fmt.Println("Start WorkerPool...")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, wp.timeout)

	wp.wg.Add(wp.nw)
	for i := 0; i < wp.nw; i++ {
		w := NewWorker(ctx, &wp.wg, wp.tasks, wp.result)
		w.Run()
	}

	go func() {
		wp.wg.Wait()
		close(wp.result)
		cancel() // HINT: Насколько правильно вызывать тут ???
	}()
}

// AddTask adds a task to the worker pool.
func (wp *WorkerPool) AddTask(t string) {
	wp.tasks <- t
}

// GetResults print results
func (wp *WorkerPool) GetResults() {
	for {
		select {
		case res := <-wp.result:
			fmt.Printf("Result: %s\n", res)
		default:
			// HINT: Тут программа зависает будто бы
			break
		}
	}
}

func main() {
	wp, err := NewWorkerPool(3, 100*time.Millisecond)
	if err != nil {
		fmt.Errorf("Error creating worker pool: %v", err)
	}

	wp.Run()

	wp.AddTask("hello")
	wp.AddTask("world")
	wp.AddTask("test")

	wp.GetResults()
}
