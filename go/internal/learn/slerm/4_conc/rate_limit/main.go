package main

// TODO: write a program that implements rate limit based on a semaphore.

/*
Напишите программу, которая реализует rate limit на основе семафора.
Например, есть 1000 запросов, которые нужно отправить по сети, но одновременно не должно выполняться более 10 запросов, чтобы не перегрузить сервер.
В качестве выполнения запросов так же используйте time.Sleep.
Выведите в консоль тайминги выполнения каждого запроса и убедитесь, что в каждую секунду выполняется не более 10 горутин.
P.S. Да, семафор уже есть в golang встроенный. Его реализацию можно посмотреть в исходнике в конце страницы https://pkg.go.dev/golang.org/x/sync/semaphore.
Она сложнее той, что предлагается написать вам.
*/

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func (s Semaphore) Acquire(n int) {

}

func (s Semaphore) Release(n int) {
	// ваш код
}
